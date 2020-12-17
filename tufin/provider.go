package tufin

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jgrancell/go-tufinclient/tufinclient"
)

type ManagementConfig struct {
	Server string
	Domain string
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"securetrack_host": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("TUFIN_SECURETRACK_HOST", nil),
			},
			"securechange_host": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("TUFIN_SECURECHANGE_HOST", nil),
			},
			"user": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("TUFIN_USER", nil),
			},
			"password": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("TUFIN_PASSWORD", nil),
			},
			"allow_insecure": &schema.Schema{
				Type: schema.TypeBool,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("TUFIN_ALLOW_INSECURE", nil),
			},
		},
		ResourcesMap:   map[string]*schema.Resource{
			"tufin_group_member": resourceGroupMember(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	debugLogOutput("configuration", "creating client connection")
	client := tufinclient.NewTufinClient(
		d.Get("securechange_host").(string),
		d.Get("securetrack_host").(string),
		d.Get("user").(string),
		d.Get("password").(string),
		d.Get("allow_insecure").(bool),
		false,
	)
	debugLogOutput("configuration", "client connection created")

	return client, diags
}
