package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				Sensitive: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req := &client.User{
		Email: d.Get("email").(string),
		Password: d.Get("password").(string),
	}

	c := m.(*client.Client)
	err := c.CreateUser(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Email)

	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Id()

	c := m.(*client.Client)
	user, err := c.GetUser(email)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("email", user.Email)

	return diags
}

// TODO: uncomment when needed to update user
//func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//	if !d.HasChange("cluster") {
//		return resourceClusterRead(ctx, d, m)
//	}
//
//	cluster := d.Get("cluster").([]interface{})[0].(map[string]interface{})
//
//	req := &client.CreateClusterRequest{}
//	helper.TerraformToStruct(cluster, req)
//
//	c := m.(*client.Client)
//	err := c.UpdateCluster(req)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	//d.Set("last_updated", time.Now().Format(time.RFC850))
//
//	return resourceClusterRead(ctx, d, m)
//}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Id()
	c := m.(*client.Client)
	err := c.DeleteUser(email)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
