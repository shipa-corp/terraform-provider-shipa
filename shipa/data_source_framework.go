package shipa

import (
	"context"
	"github.com/marcustreacy/go-terraform-provider/client"
	"github.com/marcustreacy/go-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFrameworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("id").(string)

	pool, err := c.GetPool(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("framework", helper.StructToTerraform(pool)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(name)

	return diags
}

func dataSourceFramework() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFrameworkRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"framework": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: frameworkSchema,
				},
			},
		},
	}
}

var (
	frameworkSchema = map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"provisioner": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"teams": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allowed": allowedSchema,
	}

	allowedSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"driver": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"plan": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"team": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
)
