package govia

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maxiepax/terraform-provider-govia/govia/models"
)

func resourcePool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePoolCreate,
		ReadContext:   resourcePoolRead,
		UpdateContext: resourcePoolUpdate,
		DeleteContext: resourcePoolDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"net_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"end_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"lease_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"only_serve_reimage": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourcePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	pool := models.Pool{}

	pool.Name = d.Get("name").(string)
	pool.NetAddress = d.Get("net_address").(string)
	pool.StartAddress = d.Get("start_address").(string)
	pool.EndAddress = d.Get("end_address").(string)
	pool.Netmask = d.Get("netmask").(int)
	pool.Gateway = d.Get("gateway").(string)
	pool.LeaseTime = 7000
	pool.OnlyServeReimage = true

	ret := models.Pool{}

	err := c.post("pools", pool, &ret)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	d.SetId(strconv.Itoa(ret.ID))

	resourcePoolRead(ctx, d, m)

	return diags
}

func resourcePoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	pool_in := models.Pool{}

	r, err := c.get("pools/"+d.Id(), &pool_in)
	if err != nil {
		return diag.FromErr(err)
	}

	refid := strconv.Itoa(pool_in.ID)

	reformated := map[string]interface{}{
		"id":            refid,
		"name":          pool_in.Name,
		"net_address":   pool_in.NetAddress,
		"start_address": pool_in.StartAddress,
		"end_address":   pool_in.EndAddress,
		"netmask":       pool_in.Netmask,
		"gateway":       pool_in.Gateway,
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

func resourcePoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*Client)

	//changed is set to false, will change to true if anything has changed
	changed := false

	//check every value if something has changed
	if d.HasChange("name") {
		changed = true
	}
	if d.HasChange("net_address") {
		changed = true
	}
	if d.HasChange("start_address") {
		changed = true
	}
	if d.HasChange("end_address") {
		changed = true
	}
	if d.HasChange("netmask") {
		changed = true
	}
	if d.HasChange("gateway") {
		changed = true
	}

	//if any variable was changed, patch
	if changed {

		pool := models.Pool{}

		pool.ID, _ = strconv.Atoi(d.Id())
		if name, ok := d.GetOk("name"); ok {
			pool.Name = name.(string)
		}
		if na, ok := d.GetOk("net_address"); ok {
			pool.NetAddress = na.(string)
		}
		if sa, ok := d.GetOk("start_address"); ok {
			pool.StartAddress = sa.(string)
		}
		if ea, ok := d.GetOk("end_address"); ok {
			pool.EndAddress = ea.(string)
		}
		if nm, ok := d.GetOk("netmask"); ok {
			pool.Netmask = nm.(int)
		}
		if gw, ok := d.GetOk("gateway"); ok {
			pool.Gateway = gw.(string)
		}

		err := c.patch("pools/"+d.Id(), pool)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourcePoolRead(ctx, d, m)
}

func resourcePoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	err := c.delete("pools/" + d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
