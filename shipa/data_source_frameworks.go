package shipa

import (
	"context"
	"github.com/marcustreacy/go-terraform-provider/client"
	"github.com/marcustreacy/go-terraform-provider/helper"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFrameworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	pools, err := c.ListPools()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("frameworks", helper.StructToTerraform(&pools)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceFrameworks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFrameworksRead,
		Schema: map[string]*schema.Schema{
			"frameworks": frameworksSchema,
		},
	}
}

var (
	frameworksSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: frameworkSchema,
		},
	}
)
