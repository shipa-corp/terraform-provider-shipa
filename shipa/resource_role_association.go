package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func resourceRoleAssociation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAssociationCreate,
		ReadContext:   resourceAssociationRead,
		DeleteContext: resourceAssociationDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAssociationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	role := d.Get("name").(string)
	email := d.Get("email").(string)

	c := m.(*client.Client)
	err := c.AssociateRoleToUser(role, email)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(role)

	resourceAssociationRead(ctx, d, m)

	return diags
}

func resourceAssociationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceAssociationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*client.Client)
	role := d.Id()
	old, _ := d.GetChange("email")
	email := old.(string)

	err := c.DisassociateRoleFromUser(role, email)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
