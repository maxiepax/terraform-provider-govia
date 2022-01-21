package govia

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maxiepax/terraform-provider-govia/govia/models"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
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
				Required: true,
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
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	group := models.Group{}

	group.PoolID = d.Get("pool_id").(int)
	group.Name = d.Get("name").(string)
	group.Password = d.Get("password").(string)
	if dns, ok := d.GetOk("dns"); ok {
		group.DNS = dns.(string)
	}
	if ntp, ok := d.GetOk("ntp"); ok {
		group.Ntp = ntp.(string)
	}
	group.ImageID = d.Get("image_id").(int)
	if syslog, ok := d.GetOk("syslog"); ok {
		group.Syslog = syslog.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		group.Vlan = vlan.(string)
	}
	if callback, ok := d.GetOk("callbackurl"); ok {
		group.CallbackURL = callback.(string)
	}
	if bootdisk, ok := d.GetOk("bootdisk"); ok {
		group.BootDisk = bootdisk.(string)
	}
	group.Options = make(map[string]bool)
	if options, ok := d.GetOk("options"); ok {
		opt := options.(map[string]interface{})
		for k, v := range opt {
			group.Options[k] = v.(bool)
		}
	}

	ret := models.Group{}

	err := c.post("groups", group, &ret)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	d.SetId(strconv.Itoa(ret.ID))

	resourceGroupRead(ctx, d, m)

	return diags
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	group_in := models.Group{}

	r, err := c.get("groups/"+d.Id(), &group_in)
	if err != nil {
		return diag.FromErr(err)
	}

	refid := strconv.Itoa(group_in.ID)

	reformated := map[string]interface{}{
		"id":          refid,
		"pool_id":     group_in.PoolID,
		"name":        group_in.Name,
		"dns":         group_in.DNS,
		"ntp":         group_in.Ntp,
		"image_id":    group_in.ImageID,
		"syslog":      group_in.Syslog,
		"vlan":        group_in.Vlan,
		"callbackurl": group_in.CallbackURL,
		"bootdisk":    group_in.BootDisk,
		"options":     group_in.Options,
	}

	for k, v := range reformated {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	if r.StatusCode == 404 {
		d.SetId("")
	}

	return diags
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*Client)

	//changed is set to false, will change to true if anything has changed
	changed := false

	//check every value if something has changed
	if d.HasChange("pool_id") {
		changed = true
	}
	if d.HasChange("name") {
		changed = true
	}
	if d.HasChange("dns") {
		changed = true
	}
	if d.HasChange("ntp") {
		changed = true
	}
	if d.HasChange("image_id") {
		changed = true
	}
	if d.HasChange("syslog") {
		changed = true
	}
	if d.HasChange("vlan") {
		changed = true
	}
	if d.HasChange("callbackurl") {
		changed = true
	}
	if d.HasChange("bootdisk") {
		changed = true
	}
	if d.HasChange("options") {
		changed = true
	}

	//if any variable was changed, patch
	if changed {

		group := models.Group{}

		group.ID, _ = strconv.Atoi(d.Id())
		if pool_id, ok := d.GetOk("pool_id"); ok {
			group.PoolID = pool_id.(int)
		}
		if name, ok := d.GetOk("name"); ok {
			group.Name = name.(string)
		}
		if password, ok := d.GetOk("password"); ok {
			group.Password = password.(string)
		}
		if dns, ok := d.GetOk("dns"); ok {
			group.DNS = dns.(string)
		}
		if ntp, ok := d.GetOk("ntp"); ok {
			group.Ntp = ntp.(string)
		}
		group.ImageID = d.Get("image_id").(int)
		if syslog, ok := d.GetOk("syslog"); ok {
			group.Syslog = syslog.(string)
		}
		if vlan, ok := d.GetOk("vlan"); ok {
			group.Vlan = vlan.(string)
		}
		if callback, ok := d.GetOk("callback"); ok {
			group.CallbackURL = callback.(string)
		}
		if bootdisk, ok := d.GetOk("bootdisk"); ok {
			group.BootDisk = bootdisk.(string)
		}
		group.Options = make(map[string]bool)
		if options, ok := d.GetOk("options"); ok {
			opt := options.(map[string]interface{})
			for k, v := range opt {
				group.Options[k] = v.(bool)
			}
		}

		err := c.patch("groups/"+d.Id(), group)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceGroupRead(ctx, d, m)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	err := c.delete("groups/" + d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
