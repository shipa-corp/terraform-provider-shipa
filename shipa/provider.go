package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/marcustreacy/go-terraform-provider/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHIPA_HOST", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SHIPA_TOKEN", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"shipa_framework":        resourceFramework(),
			"shipa_cluster":          resourceCluster(),
			"shipa_user":             resourceUser(),
			"shipa_plan":             resourcePlan(),
			"shipa_team":             resourceTeam(),
			"shipa_app":              resourceApp(),
			"shipa_app_env":          resourceAppEnv(),
			"shipa_app_cname":        resourceAppCname(),
			"shipa_app_deploy":       resourceAppDeploy(),
			"shipa_role":             resourceRole(),
			"shipa_permission":       resourcePermission(),
			"shipa_role_association": resourceRoleAssociation(),
			"shipa_network_policy":   resourceNetworkPolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"shipa_app":        dataSourceApp(),
			"shipa_apps":       dataSourceApps(),
			"shipa_framework":  dataSourceFramework(),
			"shipa_frameworks": dataSourceFrameworks(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	host := d.Get("host").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := client.NewClient(host, token)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Shipa client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	return c, diags
}
