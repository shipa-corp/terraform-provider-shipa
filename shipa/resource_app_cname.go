package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func resourceAppCname() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCnameCreate,
		ReadContext:   resourceAppCnameRead,
		UpdateContext: resourceAppCnameUpdate,
		DeleteContext: resourceAppCnameDelete,
		Schema: map[string]*schema.Schema{
			"app": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cname": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"encrypt": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		//Importer: &schema.ResourceImporter{
		//	StateContext: schema.ImportStatePassthroughContext,
		//},
	}
}

func resourceAppCnameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Get("app").(string)
	cname := d.Get("cname").(string)
	encrypt := d.Get("encrypt").(bool)

	req := &client.AppCname{
		Cname: cname,
		Scheme: getScheme(encrypt),
	}

	c := m.(*client.Client)
	err := c.CreateAppCname(app, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app)

	//resourceAppEnvRead(ctx, d, m)

	return diags
}

func resourceAppCnameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//name := d.Id()
	//
	//c := m.(*client.Client)
	//appEnvs, err := c.GetAppEnvs(name)
	//if err != nil {
	//	return diag.FromErr(err)
	//}
	//
	//data := &client.CreateAppEnv{
	//	Envs: filterPlatformEnvs(appEnvs),
	//}
	//
	//if err = d.Set("app_env", helper.StructToTerraform(data)); err != nil {
	//	return diag.FromErr(err)
	//}

	return diags
}

func resourceAppCnameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	//app := d.Id()
	// get new state
	app := d.Get("app").(string)
	cname := d.Get("cname").(string)
	encrypt := d.Get("encrypt").(bool)

	req := &client.AppCname{
		Cname: cname,
		Scheme: getScheme(encrypt),
	}

	// update cname
	c := m.(*client.Client)
	err := c.UpdateAppCname(app, req)
	if err != nil {
		return diag.FromErr(err)
	}

	//return resourceAppEnvRead(ctx, d, m)
	return diags
}

func resourceAppCnameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Get("app").(string)
	cname := d.Get("cname").(string)
	encrypt := d.Get("encrypt").(bool)

	req := &client.AppCname{
		Cname: cname,
		Scheme: getScheme(encrypt),
	}

	c := m.(*client.Client)
	err := c.DeleteAppCname(app, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func getScheme(encrypt bool) string {
	if encrypt {
		return "https"
	}
	return "http"
}