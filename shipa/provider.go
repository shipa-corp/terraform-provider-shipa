package shipa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shipa-corp/terraform-provider-shipa/client"
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
			"admin_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SHIPA_ADMIN_EMAIL", nil),
			},
			"admin_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SHIPA_ADMIN_PASSWORD", nil),
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
	var c *client.Client
	var err error

	token := d.Get("token").(string)
	host := d.Get("host").(string)
	adminEmail := d.Get("admin_email").(string)
	adminPassword := d.Get("admin_password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err = client.NewClient(
		client.WithHost(host),
		client.WithAuth(token, adminEmail, adminPassword))

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
