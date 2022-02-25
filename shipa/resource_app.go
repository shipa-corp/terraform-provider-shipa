package shipa

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"
)

var (
	schemaApp = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"name": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"teamowner": {
					Type:     schema.TypeString,
					Required: true,
				},
				"framework": {
					Type:     schema.TypeString,
					Required: true,
				},

				// Optional
				"platform": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"tags": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},

				// Computed
				"org": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"plan": planSchema,

				"units": unitsSchema,

				"ip": {
					Type:     schema.TypeString,
					Computed: true,
				},

				"entrypoints": entrypointsSchema,

				"routers": routersSchema,

				"lock": lockSchema,

				"status": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"error": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"app": schemaApp,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Get("app").([]interface{})[0].(map[string]interface{})
	req := &client.App{}

	helper.TerraformToStruct(app, req)

	c := m.(*client.Client)
	err := c.CreateApp(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Name)

	resourceAppRead(ctx, d, m)

	return diags
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	app, err := c.GetApp(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("app", helper.StructToTerraform(app)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("app") {
		return resourceAppRead(ctx, d, m)
	}

	app := d.Get("app").([]interface{})[0].(map[string]interface{})
	log.Printf(" ### RAW app data: %+v\n", app)

	req := &client.App{}
	helper.TerraformToStruct(app, req)

	c := m.(*client.Client)
	err := c.UpdateApp(req.Name, client.NewUpdateAppRequest(req))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAppRead(ctx, d, m)
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteApp(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
