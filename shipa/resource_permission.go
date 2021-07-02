package shipa

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shipa-corp/terraform-provider-shipa/client"
)

func resourcePermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePermissionCreate,
		ReadContext:   resourcePermissionRead,
		UpdateContext: resourcePermissionUpdate,
		DeleteContext: resourcePermissionDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == new {
						return true
					}

					oldVal, newVal := d.GetChange("permission")

					oldPermissions := convertToList(oldVal)
					newPermissions := convertToList(newVal)

					if len(oldPermissions) != len(newPermissions) {
						return false
					}

					lookup := map[string]bool{}
					for _, val := range oldPermissions {
						lookup[val] = true
					}

					for _, val := range newPermissions {
						if !lookup[val] {
							return false
						}
					}

					return true
				},
			},
		},
	}
}

func parsePermissions(d *schema.ResourceData) []string {
	return convertToList(d.Get("permission"))
}

func convertToList(value interface{}) []string {
	rawList := value.([]interface{})
	arr := make([]string, 0, len(rawList))
	for _, val := range rawList {
		arr = append(arr, val.(string))
	}
	return arr
}

func resourcePermissionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req := &client.Permission{
		Role:        d.Get("name").(string),
		Permissions: parsePermissions(d),
	}

	c := m.(*client.Client)
	err := c.CreatePermission(req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(req.Role)

	resourcePermissionRead(ctx, d, m)

	return diags
}

func resourcePermissionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Id()

	c := m.(*client.Client)
	permission, err := c.GetPermission(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", permission.Role)
	d.Set("permission", permission.Permissions)

	return diags
}

func resourcePermissionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*client.Client)
	name := d.Id()

	oldVal, newVal := d.GetChange("permission")

	oldList := convertToList(oldVal)
	newList := convertToList(newVal)

	newMap := map[string]bool{}
	for _, val := range newList {
		newMap[val] = true
	}

	// to delete
	for _, val := range oldList {
		if !newMap[val] {
			err := c.DeletePermission(name, val)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	// create new permissions
	err := c.CreatePermission(&client.Permission{
		Role:        name,
		Permissions: newList,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	resourcePermissionRead(ctx, d, m)

	return diags
}

func resourcePermissionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*client.Client)
	name := d.Id()
	old, _ := d.GetChange("permission")
	permissions := convertToList(old)

	for _, val := range permissions {
		err := c.DeletePermission(name, val)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
