package myfile

import (
	"context"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider for custom file handling
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"encoding": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "uft8",
				ValidateDiagFunc: validateEncoding,
				Description:      "File Encoding",
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap:         getProviderResources(),
	}
}

func validateEncoding(
	v interface{}, path cty.Path,
) diag.Diagnostics {
	encoding := v.(string)
	if encoding != "utf8" {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Only supported option is 'utf8'",
			},
		}
	}
	return diag.Diagnostics{}
}

func providerConfigure(
	ctx context.Context, d *schema.ResourceData,
) (interface{}, diag.Diagnostics) {
	encoding := d.Get("encoding").(string)
	client := FileClient{
		Encoding: encoding,
	}
	return client, diag.Diagnostics{}
}

func getProviderResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"myfile_file": {
			Schema: map[string]*schema.Schema{
				"path": {
					Type:        schema.TypeString,
					Required:    true,
					ForceNew:    true,
					Description: "File path",
				},
				"contents": {
					Type:        schema.TypeString,
					Required:    true,
					ForceNew:    false,
					Description: "File contents",
				},
				"owner": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "File owner",
				},
			},
			CreateContext: resourceFileCreate,
			ReadContext:   resourceFileRead,
			UpdateContext: resourceFileUpdate,
			DeleteContext: resourceFileDelete,
		},
	}
}

func resourceFileCreate(
	ctx context.Context, d *schema.ResourceData, m interface{},
) diag.Diagnostics {
	path := d.Get("path").(string)
	contents := d.Get("contents").(string)
	client := m.(FileClient)

	err := client.Create(path, contents)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(path)

	owner, err := client.Owner(path)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("owner", owner)

	return diag.Diagnostics{}
}

func resourceFileRead(
	ctx context.Context, d *schema.ResourceData, m interface{},
) diag.Diagnostics {
	path := d.Get("path").(string)
	client := m.(FileClient)

	contents, err := client.Read(path)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("contents", contents); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func resourceFileUpdate(
	ctx context.Context, d *schema.ResourceData, m interface{},
) diag.Diagnostics {
	path := d.Get("path").(string)
	contents := d.Get("contents").(string)
	client := m.(FileClient)

	err := client.Update(path, contents)
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func resourceFileDelete(
	ctx context.Context, d *schema.ResourceData, m interface{},
) diag.Diagnostics {
	path := d.Get("path").(string)
	client := m.(FileClient)

	err := client.Delete(path)
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
