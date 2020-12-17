package tufin

import (
  "context"
  "fmt"
  "os"
  "regexp"

  "github.com/hashicorp/go-uuid"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/jgrancell/go-tufinclient/tufinclient"
)

func resourceGroupMember() *schema.Resource {
  return &schema.Resource{
    CreateContext: resourceGroupMemberCreate,
    ReadContext:   resourceGroupMemberRead,
    //UpdateContext: resourceGroupMemberUpdate,
    DeleteContext: resourceGroupMemberDelete,
    Schema: map[string]*schema.Schema{
      "group_name": &schema.Schema{
        Type:     schema.TypeString,
        ForceNew: true,
        Required: true,
        ValidateFunc: func(val interface{}, key string) (warns[]string, errs []error) {
          v := val.(string)
          if match, _ := regexp.MatchString("[A-Z_]*", v); ! match {
            errs = append(errs, fmt.Errorf("%q includes invalid characters. May contain [uppercase letters, underscores].", key))
          }
          return
        },
      },
      "ip_address": &schema.Schema{
        Type:     schema.TypeString,
        ForceNew: true,
        Required: true,
        ValidateFunc: func(val interface{}, key string) (warns[]string, errs []error) {
          v := val.(string)
          if match, _ := regexp.MatchString("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+", v); ! match {
            errs = append(errs, fmt.Errorf("%q includes invalid characters. May contain [uppercase letters, underscores].", key))
          }
          return
        },
      },
    },
    SchemaVersion: 1,
  }
}

func resourceGroupMemberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  group_name := d.Get("group_name").(string)
  ip_address := d.Get("ip_address").(string)

  client := m.(*tufinclient.TufinClient)

  debugLogOutput("create", "beginning creation reconcilliation")

  added, err := client.AddIPToGroup(ip_address, group_name)
  if err != nil {
    diag.FromErr(err)
  }

  debugLogOutput("create", "completed creation call")

  if added == false {
    diag.FromErr(fmt.Errorf("Group %s does not exist or IP %s is not a viable member.", group_name, ip_address))
  } else {
    debugLogOutput("group membership creation", "added IP address to group membership")
  }

  newUuid, _ := uuid.GenerateUUID()
  d.SetId(newUuid)

  return diags
}

func resourceGroupMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  return diags
}

func resourceGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  new_group := d.Get("group_name").(string)
  new_ip := d.Get("ip_address").(string)
  var old_group string
  var old_ip string

  client := m.(*tufinclient.TufinClient)

  if d.HasChange("ip_address") {
    // Get change returns old, new as interfaces so we need to specifically cast to string
    old_ip_interface, _ := d.GetChange("ip_address")
    old_ip = old_ip_interface.(string)
  } else {
    old_ip = new_ip
  }

  if d.HasChange("group_name") {
    // Get change returns old, new as interfaces so we need to specifically cast to string
    group_name_interface, _ := d.GetChange("group_name")
    old_group = group_name_interface.(string)
  } else {
    old_group = new_group
  }

  removed, err := client.RemoveIPFromGroup(old_ip, old_group)
  if err != nil {
    diag.FromErr(err)
  }

  if removed == false {
    diag.FromErr(fmt.Errorf("Failed to remove %s IP from %s Group ", old_ip, old_group))
  } else {
    debugLogOutput("group membership update deletion", "removed IP address from group membership")
  }

  added, err := client.AddIPToGroup(new_ip, new_group)
  if err != nil {
    diag.FromErr(err)
  }

  if added == false {
    diag.FromErr(fmt.Errorf("Group %s does not exist or IP %s is not a viable member.", new_group, new_ip))
  } else {
    debugLogOutput("group membership update creation", "added IP address to group membership")
  }

  return diags
}

func resourceGroupMemberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  group_name := d.Get("group_name").(string)
  ip_address := d.Get("ip_address").(string)

  client := m.(*tufinclient.TufinClient)

  removed, err := client.RemoveIPFromGroup(ip_address, group_name)
  if err != nil {
    diag.FromErr(err)
  }

  if removed == false {
    diag.FromErr(fmt.Errorf("Failed to remove %s IP from %s Group ", ip_address, group_name))
  } else {
    debugLogOutput("group membership deletion", "removed IP address from group membership")
  }

  d.SetId("")

  return diags
}

func debugLogOutput(id string, output string) {
  //Debug log for development
  f, _ := os.OpenFile("./terraform-provider-tufin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  defer f.Close()
  _, err := f.WriteString(id+": "+output+"\n")
  if err != nil {
    panic(err)
  }
  f.Sync()
}
