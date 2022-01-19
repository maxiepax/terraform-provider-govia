package govia

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,
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
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	data := d.Get("id").(int)

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

	group_in := Group{}

	idstring := strconv.Itoa(data)

	_, err := c.get(fmt.Sprintf("groups/%s", idstring), &group_in)
	if err != nil {
		return diag.FromErr(err)
	}

	reformated := map[string]interface{}{
		"id":          group_in.ID,
		"pool_id":     group_in.PoolID,
		"name":        group_in.Name,
		"dns":         group_in.DNS,
		"ntp":         group_in.Ntp,
		"image_id":    group_in.ImageID,
		"ks":          group_in.Ks,
		"syslog":      group_in.Syslog,
		"vlan":        group_in.Vlan,
		"callbackurl": group_in.Callbackurl,
		"bootdisk":    group_in.BootDisk,
		"options":     group_in.Options,
	}

	for k, v := range reformated {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(int64(group_in.ID), 10))

	return diags
}
