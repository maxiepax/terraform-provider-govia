package govia

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOVIA_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GOVIA_PASSWORD", nil),
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GOVIA_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"govia_group":   resourceGroups(),
			"govia_address": resourceAddresses(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"govia_groups": dataSourceGroups(),
			"govia_group":  dataSourceGroup(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	url := d.Get("url").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if username == "" {
		return nil, diag.FromErr(fmt.Errorf("missing go-via username"))
	}

	if password == "" {
		return nil, diag.FromErr(fmt.Errorf("missing go-via password"))
	}

	if url == "" {
		return nil, diag.FromErr(fmt.Errorf("missing go-via url e.g. (https://172.30.0.10:8443)"))
	}

	c := newClient(username, password, url)

	return c, diags
}
