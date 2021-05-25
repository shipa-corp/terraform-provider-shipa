package shipa

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/marcustreacy/go-terraform-provider/client"
	"github.com/marcustreacy/go-terraform-provider/helper"
	"log"
)

var (
	schemaClusterEndpoint = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"addresses": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"ca_cert": {
					Type:     schema.TypeString,
					Required: true,
				},
				"token": {
					Type:     schema.TypeString,
					Optional: true,
					Sensitive: true,
					ConflictsWith: []string{"cluster.0.endpoint.0.password", "cluster.0.endpoint.0.username"},
				},
				"client_key": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"client_cert": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"username": {
					Type:      schema.TypeString,
					Optional:  true,
					RequiredWith: []string{"cluster.0.endpoint.0.password"},
				},
				"password": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
					RequiredWith: []string{"cluster.0.endpoint.0.username"},
					ConflictsWith: []string{"cluster.0.endpoint.0.token"},
				},
			},
		},
	}

	schemaClusterResources = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"frameworks": schemaClusterFrameworks,
				"ingress_controllers": schemaIngressControllersCreate,
			},
		},
	}

	schemaClusterFrameworks = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaIngressControllersCreate = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ingress_ip": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"service_type": {
					Type:     schema.TypeString,
					Optional: true,
					ValidateDiagFunc: func(val interface{}, key cty.Path) (diag.Diagnostics) {
						var diags diag.Diagnostics

						v := val.(string)

						log.Println(">>>> GOT VALUE: ", v)

						allowedValues := []string{"traefik", "istio"}
						for _, allowed := range allowedValues {
							if allowed == v {
								return diags
							}
						}

						return diag.Errorf("invalid value: %s, choose one of: %+v", v, allowedValues)
					},
				},
				"type": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"http_port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"https_port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"protected_port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"debug": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"acme_email": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"acme_server": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}

)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Schema: map[string]*schema.Schema{
			"cluster": {
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
						"endpoint": schemaClusterEndpoint,

						// Optional
						"resources": schemaClusterResources,
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cluster := d.Get("cluster").([]interface{})[0].(map[string]interface{})
	req := &client.TerraformCreateClusterRequest{}

	helper.TerraformToStruct(cluster, req)

	c := m.(*client.Client)
	err := c.CreateCluster(client.NewCreateClusterRequest(req))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Name)

	resourceClusterRead(ctx, d, m)

	return diags
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	cluster, err := c.GetCluster(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("cluster", helper.StructToTerraform(client.NewTerraformCreateClusterRequest(cluster))); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("cluster") {
		return resourceClusterRead(ctx, d, m)
	}

	cluster := d.Get("cluster").([]interface{})[0].(map[string]interface{})

	req := &client.TerraformCreateClusterRequest{}
	helper.TerraformToStruct(cluster, req)

	c := m.(*client.Client)
	err := c.UpdateCluster(client.NewCreateClusterRequest(req))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceClusterRead(ctx, d, m)
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteCluster(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
