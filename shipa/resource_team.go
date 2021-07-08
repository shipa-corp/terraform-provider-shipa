package shipa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/shipa-corp/terraform-provider-shipa/client"
	"github.com/shipa-corp/terraform-provider-shipa/helper"
)

var (
	schemaTeam = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				// Optional
				"tags": {
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

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
		Schema: map[string]*schema.Schema{
			"team": schemaTeam,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	team := d.Get("team").([]interface{})[0].(map[string]interface{})
	req := &client.Team{}

	helper.TerraformToStruct(team, req)

	c := m.(*client.Client)
	err := c.CreateTeam(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Name)

	resourceTeamRead(ctx, d, m)

	return diags
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	team, err := c.GetTeam(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("team", helper.StructToTerraform(team)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTeamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("team") {
		return resourceTeamRead(ctx, d, m)
	}

	name := d.Id()
	team := d.Get("team").([]interface{})[0].(map[string]interface{})

	req := &client.Team{}
	helper.TerraformToStruct(team, req)
	var newName string

	if name != req.Name {
		newName = req.Name
	}

	updateReq := &client.UpdateTeamRequest{
		Tags: req.Tags,
	}

	if newName != "" {
		updateReq.Name = newName
	}


	c := m.(*client.Client)
	err := c.UpdateTeam(name, updateReq)
	if err != nil {
		return diag.FromErr(err)
	}
	if newName != "" {
		d.SetId(newName)
	}

	return resourceTeamRead(ctx, d, m)
}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteTeam(name)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
