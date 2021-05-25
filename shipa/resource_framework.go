package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/marcustreacy/go-terraform-provider/client"
	"github.com/marcustreacy/go-terraform-provider/helper"
	"log"
)

var (
	schemaPoolConfig = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"general":    schemaPoolConfigGeneral,
				"shipa_node": schemaPoolConfigNode,
			},
		},
	}

	schemaPoolConfigNode = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"drivers": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"auto_scale": schemaPoolConfigNodeAutoScale,
			},
		},
	}

	schemaPoolConfigNodeAutoScale = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"max_container": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"max_memory": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"scale_down": {
					Type:     schema.TypeFloat,
					Optional: true,
				},
				"rebalance": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}

	schemaPoolConfigGeneral = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"setup":    schemaPoolConfigGeneralSetup,
				"plan":     schemaPoolConfigGeneralPlan,
				"security": schemaPoolConfigGeneralSecurity,
				"access":   schemaPoolConfigGeneralServiceAccess,
				"services": schemaPoolConfigGeneralServiceAccess,
				"router": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"volumes": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"app_quota":        schemaPoolConfigGeneralAppQuota,
				"container_policy": schemaPoolConfigGeneralContainerPolicy,
				"network_policy":   schemaPoolConfigGeneralNetworkPolicy,
			},
		},
	}

	schemaPoolConfigGeneralSetup = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"provisioner": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"kubernetes_namespace": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"default": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"public": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}

	schemaPoolConfigGeneralPlan = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}

	schemaPoolConfigGeneralSecurity = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"disable_scan": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"scan_platform_layers": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"ignore_components": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"ignore_cves": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaPoolConfigGeneralServiceAccess = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"append": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"blacklist": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaPoolConfigGeneralAppQuota = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"limit": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}

	schemaPoolConfigGeneralContainerPolicy = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allowed_hosts": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaPoolConfigGeneralNetworkPolicy = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ingress": schemaNetworkPolicy,
				"egress":  schemaNetworkPolicy,
				"disable_app_policies": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}

	schemaNetworkPolicy = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"policy_mode": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"custom_rules": schemaNetworkPolicyRule,
				"shipa_rules":  schemaNetworkPolicyRule,
				"shipa_rules_enabled": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaNetworkPolicyRule = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"ports": schemaNetworkPort,
				"peers": schemaNetworkPeer,
				"allowed_apps": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaNetworkPort = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"protocol": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"port": schemaNetworkPortDetails,
			},
		},
	}

	schemaNetworkPortDetails = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"intval": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"strval": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}

	schemaNetworkPeer = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"podselector":       schemaNetworkPeerSelector,
				"namespaceselector": schemaNetworkPeerSelector,
				"ipblock": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaNetworkPeerSelector = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"matchlabels": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"matchexpressions": schemaNetworkPeerSelectorExpression,
			},
		},
	}

	schemaNetworkPeerSelectorExpression = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"operator": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"values": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
)

func resourceFramework() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePoolCreate,
		ReadContext:   resourcePoolRead,
		UpdateContext: resourcePoolUpdate,
		DeleteContext: resourcePoolDelete,
		Schema: map[string]*schema.Schema{
			"framework": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required for creation
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"provisioner": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Optional
						"resources": schemaPoolConfig,
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pool := d.Get("framework").([]interface{})[0].(map[string]interface{})
	name := pool["name"].(string)
	provisioner := pool["provisioner"].(string)

	config := &client.PoolConfig{
		Name:      name,
		Resources: &client.PoolResources{},
	}

	helper.TerraformToStruct(pool["resources"], config.Resources)

	log.Printf("RAW Pool: %+v\n", pool)

	if config.Resources == nil {
		config.Resources = &client.PoolResources{}
	}

	if config.Resources.General == nil {
		config.Resources.General = &client.PoolGeneral{}
	}

	if config.Resources.General.Setup == nil {
		config.Resources.General.Setup = &client.PoolSetup{}
	}

	config.Resources.General.Setup.Provisioner = provisioner

	c := m.(*client.Client)
	err := c.CreatPoolConfig(config)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	resourcePoolRead(ctx, d, m)

	return diags
}

func resourcePoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	pool, err := c.GetPoolConfig(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("framework", convertPoolConfigToRawData(pool)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("framework") {
		return resourcePoolRead(ctx, d, m)
	}

	pool := d.Get("framework").([]interface{})[0].(map[string]interface{})
	provisioner := pool["provisioner"].(string)

	config := &client.PoolConfig{
		Name:      d.Id(),
		Resources: &client.PoolResources{},
	}

	helper.TerraformToStruct(pool["resources"], config.Resources)

	if config.Resources == nil {
		config.Resources = &client.PoolResources{}
	}

	if config.Resources.General == nil {
		config.Resources.General = &client.PoolGeneral{}
	}

	if config.Resources.General.Setup == nil {
		config.Resources.General.Setup = &client.PoolSetup{}
	}

	config.Resources.General.Setup.Provisioner = provisioner

	c := m.(*client.Client)
	err := c.UpdatePoolConfig(config)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourcePoolRead(ctx, d, m)
}

func resourcePoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	c := m.(*client.Client)
	err := c.DeletePool(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertPoolConfigToRawData(p *client.PoolConfig) []interface{} {
	var resources []interface{}
	if p.Resources != nil {
		resources = helper.StructToTerraform(p.Resources)
	}

	var provisioner string
	if p.Resources != nil && p.Resources.General != nil && p.Resources.General.Setup != nil {
		provisioner = p.Resources.General.Setup.Provisioner
	}

	raw := map[string]interface{}{
		"name":        p.Name,
		"provisioner": provisioner,
		"resources":   resources,
	}
	return []interface{}{raw}
}
