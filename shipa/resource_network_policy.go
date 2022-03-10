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
	schemaConfigNetworkPolicy = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ingress": schemaNetworkPolicy,
				"egress":  schemaNetworkPolicy,
				"restart_app": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
)

func resourceNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkPolicyCreate,
		ReadContext:   resourceNetworkPolicyRead,
		UpdateContext: resourceNetworkPolicyUpdate,
		DeleteContext: resourceNetworkPolicyDelete,
		Schema: map[string]*schema.Schema{
			"network_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required for creation
						"app": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Optional
						"network_policy": schemaConfigNetworkPolicy,
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetworkPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	networkPolicy := d.Get("network_policy").([]interface{})[0].(map[string]interface{})
	app := networkPolicy["app"].(string)

	config := &client.NetworkPolicy{}

	helper.TerraformToStruct(networkPolicy["network_policy"], config)

	log.Printf("RAW NetworkPolicy: %+v\n", networkPolicy)

	c := m.(*client.Client)
	err := c.CreateOrUpdateNetworkPolicy(app, config)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app)

	resourcePoolRead(ctx, d, m)

	return diags
}

func resourceNetworkPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appName := d.Id()

	c := m.(*client.Client)
	networkPolicy, err := c.GetNetworkPolicy(appName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("network_policy", convertNetworkPolicyToRawData(appName, networkPolicy)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetworkPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("network_policy") {
		return resourceNetworkPolicyRead(ctx, d, m)
	}

	networkPolicy := d.Get("network_policy").([]interface{})[0].(map[string]interface{})
	app := networkPolicy["app"].(string)

	config := &client.NetworkPolicy{}

	helper.TerraformToStruct(networkPolicy["network_policy"], config)

	log.Printf("RAW NetworkPolicy: %+v\n", networkPolicy)

	c := m.(*client.Client)
	err := c.CreateOrUpdateNetworkPolicy(app, config)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetworkPolicyRead(ctx, d, m)
}

func resourceNetworkPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	app := d.Id()
	c := m.(*client.Client)
	err := c.DeleteNetworkPolicy(app)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertNetworkPolicyToRawData(app string, p *client.NetworkPolicy) []interface{} {
	raw := map[string]interface{}{
		"app":            app,
		"network_policy": helper.StructToTerraform(p),
	}
	return []interface{}{raw}
}
