package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/marcustreacy/go-terraform-provider/client"
	"github.com/marcustreacy/go-terraform-provider/helper"
)

var (
	schemaAppEnv = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"envs": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:     schema.TypeString,
								Required: true,
							},
							"value": {
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
				},
				// Optional
				"norestart": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"private": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
)

func resourceAppEnv() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppEnvCreate,
		ReadContext:   resourceAppEnvRead,
		UpdateContext: resourceAppEnvUpdate,
		DeleteContext: resourceAppEnvDelete,
		Schema: map[string]*schema.Schema{
			"app": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_env": schemaAppEnv,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAppEnvCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Get("app").(string)
	appEnvs := d.Get("app_env").([]interface{})[0].(map[string]interface{})
	req := &client.CreateAppEnv{}
	helper.TerraformToStruct(appEnvs, req)

	c := m.(*client.Client)
	req.Envs = filterPlatformEnvs(req.Envs)
	err := c.CreateAppEnvs(app, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app)

	resourceAppEnvRead(ctx, d, m)

	return diags
}

func resourceAppEnvRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	appEnvs, err := c.GetAppEnvs(name)
	if err != nil {
		return diag.FromErr(err)
	}

	data := &client.CreateAppEnv{
		Envs: filterPlatformEnvs(appEnvs),
	}

	if err = d.Set("app_env", helper.StructToTerraform(data)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func filterPlatformEnvs(input []*client.AppEnv) []*client.AppEnv {
	platformEnvs := map[string]bool{
		"SHIPA_APP_TOKEN": true,
		"SHIPA_APPNAME":   true,
		"SHIPA_SERVICES":  true,
		"SHIPA_APPDIR":    true,
	}

	out := make([]*client.AppEnv, 0)

	for _, env := range input {
		if platformEnvs[env.Name] {
			continue
		}
		out = append(out, env)
	}

	return out
}

func resourceAppEnvUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("app_env") {
		return resourceAppEnvRead(ctx, d, m)
	}

	app := d.Id()
	// get new state
	appEnv := d.Get("app_env").([]interface{})[0].(map[string]interface{})
	req := &client.CreateAppEnv{}
	helper.TerraformToStruct(appEnv, req)
	req.Envs = filterPlatformEnvs(req.Envs)
	newState := map[string]bool{}
	for _, item := range req.Envs {
		newState[item.Name] = true
	}

	// get current state
	c := m.(*client.Client)
	actualAppEnvs, err := c.GetAppEnvs(app)
	if err != nil {
		return diag.FromErr(err)
	}

	// get diff (to delete)
	envsToDelete := []*client.AppEnv{}
	for _, item := range filterPlatformEnvs(actualAppEnvs) {
		if !newState[item.Name] {
			envsToDelete = append(envsToDelete, item)
		}
	}

	// delete envs
	err = c.DeleteAppEnvs(app, &client.CreateAppEnv{
		Envs:      envsToDelete,
		NoRestart: req.NoRestart,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	// update envs
	err = c.CreateAppEnvs(app, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAppEnvRead(ctx, d, m)
}

func resourceAppEnvDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	appEnvs := d.Get("app_env").([]interface{})[0].(map[string]interface{})
	req := &client.CreateAppEnv{}
	helper.TerraformToStruct(appEnvs, req)
	req.Envs = filterPlatformEnvs(req.Envs)

	c := m.(*client.Client)
	err := c.DeleteAppEnvs(name, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
