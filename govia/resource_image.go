package govia

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maxiepax/terraform-provider-govia/govia/models"
)

func resourceImage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceImageCreate,
		ReadContext:   resourceImageRead,
		UpdateContext: resourceImageUpdate,
		DeleteContext: resourceImageDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hash": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceImageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	image := models.Image{}

	image.ISOImage = d.Get("iso_image").(string)
	if hash, ok := d.GetOk("hash"); ok {
		image.Hash = hash.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		image.Description = desc.(string)
	}

	ret := models.Image{}

	err := c.postFile("images", image, &ret)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	d.SetId(strconv.Itoa(ret.ID))

	resourceImageRead(ctx, d, m)

	return diags
}

func resourceImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	image_in := models.Image{}

	r, err := c.get("images/"+d.Id(), &image_in)
	if err != nil {
		return diag.FromErr(err)
	}

	refid := strconv.Itoa(image_in.ID)

	reformated := map[string]interface{}{
		"id":          refid,
		"iso_image":   image_in.ISOImage,
		"hash":        image_in.Hash,
		"description": image_in.Description,
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

func resourceImageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*Client)

	//changed is set to false, will change to true if anything has changed
	changed := false

	//check every value if something has changed
	if d.HasChange("description") {
		changed = true
	}

	//if any variable was changed, patch
	if changed {

		image := models.Image{}

		image.ID, _ = strconv.Atoi(d.Id())
		if desc, ok := d.GetOk("description"); ok {
			image.Description = desc.(string)
		}

		err := c.patch("images/"+d.Id(), image)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceImageRead(ctx, d, m)
}

func resourceImageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	err := c.delete("images/" + d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
