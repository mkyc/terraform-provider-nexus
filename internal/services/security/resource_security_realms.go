package security

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	"github.com/datadrivers/terraform-provider-nexus/internal/tools"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecurityRealms() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to activate Nexus Security realms.",

		Create: resourceRealmsCreate,
		Read:   resourceRealmsRead,
		Update: resourceRealmsUpdate,
		Delete: resourceRealmsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": common.ResourceID,
			"active": {
				Description: "The realm IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				Type:     schema.TypeList,
			},
		},
	}
}

func resourceRealmsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)
	realmIDs := tools.InterfaceSliceToStringSlice(d.Get("active").([]interface{}))
	if err := client.Security.Realm.Activate(realmIDs); err != nil {
		return err
	}

	return resourceRealmsRead(d, m)
}

func resourceRealmsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)
	activeRealms, err := client.Security.Realm.ListActive()
	if err != nil {
		return err
	}

	d.SetId("active")
	if err := d.Set("active", tools.StringSliceToInterfaceSlice(activeRealms)); err != nil {
		return err
	}

	return nil
}

func resourceRealmsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRealmsCreate(d, m)
}

func resourceRealmsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
