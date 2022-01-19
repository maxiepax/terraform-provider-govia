package govia

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"pool_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"dns": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ntp": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"image_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ks": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"syslog": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"callbackurl": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bootdisk": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type:     schema.TypeBool,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	type Group struct {
		ID          int             `json:"id"`
		PoolID      int             `json:"pool_id"`
		Name        string          `json:"name"`
		DNS         string          `json:"dns"`
		Ntp         string          `json:"ntp"`
		ImageID     int             `json:"image_id"`
		Ks          string          `json:"ks"`
		Syslog      string          `json:"syslog"`
		Vlan        string          `json:"vlan"`
		Callbackurl string          `json:"callbackurl"`
		BootDisk    string          `json:"bootdisk" gorm:"type:varchar(255)"`
		Options     map[string]bool `json:"options,omitempty"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`
	}

	groups_in := make([]Group, 0)

	_, err := c.get("groups", &groups_in)
	if err != nil {
		return diag.FromErr(err)
	}

	reformated := make([]map[string]interface{}, 0)

	for _, v := range groups_in {
		reformated = append(reformated, map[string]interface{}{
			"id":          v.ID,
			"pool_id":     v.PoolID,
			"name":        v.Name,
			"dns":         v.DNS,
			"ntp":         v.Ntp,
			"image_id":    v.ImageID,
			"ks":          v.Ks,
			"syslog":      v.Syslog,
			"vlan":        v.Vlan,
			"callbackurl": v.Callbackurl,
			"bootdisk":    v.BootDisk,
			"options":     v.Options,
		})
	}

	if err := d.Set("groups", reformated); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
