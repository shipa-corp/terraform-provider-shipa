package shipa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"
)

var (
	schemaPlan = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"name": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"memory": {
					Type:     schema.TypeInt,
					Required: true,
					ForceNew: true,
				},
				"swap": {
					Type:     schema.TypeInt,
					Required: true,
					ForceNew: true,
				},
				"cpushare": {
					Type:     schema.TypeInt,
					Required: true,
					ForceNew: true,
				},
				"teams": {
					Type:     schema.TypeList,
					Required: true,
					ForceNew: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},

				// Optional
				"public": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"org": {
					Type:     schema.TypeString,
					Optional: true,
					DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						return true
					},
				},
				"default": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
)

func resourcePlan() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePlanCreate,
		ReadContext:   resourcePlanRead,
		DeleteContext: resourcePlanDelete,
		Schema: map[string]*schema.Schema{
			"plan": schemaPlan,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePlanCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	plan := d.Get("plan").([]interface{})[0].(map[string]interface{})
	req := &client.Plan{}

	helper.TerraformToStruct(plan, req)

	c := m.(*client.Client)
	err := c.CreatePlan(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Name)

	resourcePlanRead(ctx, d, m)

	return diags
}

func resourcePlanRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	plan, err := c.GetPlan(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("plan", helper.StructToTerraform(plan)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePlanDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	c := m.(*client.Client)
	err := c.DeletePlan(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
