package shipa

import (
	"context"
	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("id").(string)

	app, err := c.GetApp(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("app", helper.StructToTerraform(app)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(name)

	return diags
}

func dataSourceApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"app": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: appSchema,
				},
			},
		},
	}
}
