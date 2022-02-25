package shipa

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"
)

var (
	schemaAppDeploy = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"image": {
					Type:     schema.TypeString,
					Required: true,
				},
				// Optional
				"private_image": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"registry_user": {
					Type:         schema.TypeString,
					Optional:     true,
					RequiredWith: []string{"deploy.0.private_image"},
				},
				"registry_secret": {
					Type:         schema.TypeString,
					Optional:     true,
					RequiredWith: []string{"deploy.0.private_image"},
				},
				"steps": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"step_weight": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"step_interval": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"protocol": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice([]string{"TCP", "UDP"}, false),
				},
				"detach": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"message": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"shipa_yaml": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"origin": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
)

func resourceAppDeploy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppDeployCreateUpdate,
		UpdateContext: resourceAppDeployCreateUpdate,
		ReadContext:   resourceAppDeployRead,
		DeleteContext: resourceAppDeployDelete,
		Schema: map[string]*schema.Schema{
			"app": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deploy": schemaAppDeploy,
		},
	}
}

func resourceAppDeployCreateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Get("app").(string)
	deploy := d.Get("deploy").([]interface{})[0].(map[string]interface{})
	req := &client.AppDeploy{}
	helper.TerraformToStruct(deploy, req)

	c := m.(*client.Client)

	retries := 0
	err := resource.RetryContext(ctx, time.Minute*1, func() *resource.RetryError {
		err := c.DeployApp(app, req)
		if err != nil {
			if retries == 3 {
				return resource.NonRetryableError(fmt.Errorf("failed to deploy app after 3 tries, %v", err))
			}
			time.Sleep(time.Minute)
			retries++
			return resource.RetryableError(err)
		}

		return nil
	})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app)

	resourceAppDeployRead(ctx, d, m)

	return diags
}

func resourceAppDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceAppDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
