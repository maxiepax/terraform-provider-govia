package govia

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maxiepax/terraform-provider-govia/govia/models"
)

func resourceAddresses() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAddressesCreate,
		ReadContext:   resourceAddressesRead,
		UpdateContext: resourceAddressesUpdate,
		DeleteContext: resourceAddressesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reimage": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"pool_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"progresstext": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAddressesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	address := models.Address{}

	address.IP = d.Get("ip").(string)
	address.Mac = d.Get("mac").(string)
	address.Hostname = d.Get("hostname").(string)
	address.Domain = d.Get("domain").(string)
	address.PoolID = d.Get("pool_id").(int)
	address.GroupID = d.Get("group_id").(int)
	if reimage, ok := d.GetOk("reimage"); ok {
		address.Reimage = reimage.(bool)
	}

	ret := models.Address{}

	err := c.post("addresses", address, &ret)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	d.SetId(strconv.Itoa(ret.ID))

	resourceAddressesRead(ctx, d, m)

	return diags
}

func resourceAddressesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	address_in := models.Address{}

	r, err := c.get("addresses/"+d.Id(), &address_in)
	if err != nil {
		return diag.FromErr(err)
	}

	refid := strconv.Itoa(address_in.ID)

	reformated := map[string]interface{}{
		"id":           refid,
		"ip":           address_in.IP,
		"mac":          address_in.Mac,
		"hostname":     address_in.Hostname,
		"domain":       address_in.Domain,
		"reimage":      address_in.Reimage,
		"pool_id":      address_in.PoolID,
		"group_id":     address_in.GroupID,
		"progress":     address_in.Progress,
		"progresstext": address_in.Progresstext,
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

func resourceAddressesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*Client)

	//changed is set to false, will change to true if anything has changed
	changed := false

	//check every value if something has changed
	if d.HasChange("ip") {
		changed = true
	}
	if d.HasChange("mac") {
		changed = true
	}
	if d.HasChange("hostname") {
		changed = true
	}
	if d.HasChange("domain") {
		changed = true
	}
	if d.HasChange("reimage") {
		changed = true
	}
	if d.HasChange("pool_id") {
		changed = true
	}
	if d.HasChange("group_id") {
		changed = true
	}

	//if any variable was changed, patch
	if changed {

		address := models.Address{}

		address.ID, _ = strconv.Atoi(d.Id())
		if ip, ok := d.GetOk("ip"); ok {
			address.IP = ip.(string)
		}
		if mac, ok := d.GetOk("mac"); ok {
			address.Mac = mac.(string)
		}

		if hostname, ok := d.GetOk("hostname"); ok {
			address.Hostname = hostname.(string)
		}
		if domain, ok := d.GetOk("domain"); ok {
			address.Domain = domain.(string)
		}
		if reimage, ok := d.GetOk("reimage"); ok {
			address.Reimage = reimage.(bool)
		}
		if pool_id, ok := d.GetOk("pool_id"); ok {
			address.PoolID = pool_id.(int)
		}
		if group_id, ok := d.GetOk("group_id"); ok {
			address.GroupID = group_id.(int)
		}

		err := c.patch("addresses/"+d.Id(), address)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceAddressesRead(ctx, d, m)
}

func resourceAddressesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	err := c.delete("addresses/" + d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
