package shipa

import (
	"context"
	"strconv"
	"time"

	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	apps, err := c.ListApps()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("apps", helper.StructToTerraform(&apps)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceApps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppsRead,
		Schema: map[string]*schema.Schema{
			"apps": appsSchema,
		},
	}
}

var (
	addressSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"scheme": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"host": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"opaque": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"user": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"path": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"raw_path": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"force_query": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"raw_query": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"fragment": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"raw_fragment": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}

	unitsSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"app_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"process_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"ip": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"status": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"version": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"org": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"host_addr": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"host_port": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address": addressSchema,
			},
		},
	}

	planSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"memory": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"swap": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"cpushare": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"default": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"public": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"org": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"teams": {
					Type:     schema.TypeList,
					Computed: true,
					Elem:     schema.TypeString,
				},
			},
		},
	}

	lockSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"locked": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"reason": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"owner": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"acquire_date": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}

	entrypointsSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"cname": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"scheme": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}

	routersSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"opts": {
					Type:     schema.TypeMap,
					Computed: true,
				},
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"default": {
					Type:     schema.TypeBool,
					Computed: true,
				},
			},
		},
	}

	appsSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: appSchema,
		},
	}

	appSchema = map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment": {
			Type:     schema.TypeList,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"framework": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"namespace": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"teamowner": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"plan": planSchema,

		"units": unitsSchema,

		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"org": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"entrypoints": entrypointsSchema,

		"routers": routersSchema,

		"lock": lockSchema,

		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     schema.TypeString,
		},
		"platform": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"error": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
)
