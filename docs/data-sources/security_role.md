---
page_title: "Data Source nexus_security_role"
subcategory: "Security"
description: |-
  Use this to get a specified secrity role.
---
# Data Source nexus_security_role
Use this to get a specified secrity role.
## Example Usage
```terraform
data "nexus_security_role" "nx_admin" {
  roleid      = "nx-admin"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `roleid` (String) The id of the role.

### Read-Only

- `description` (String) The description of this role.
- `id` (String) Used to identify data source at nexus
- `name` (String) The name of the role.
- `privileges` (Set of String) The privileges of this role.
- `roles` (Set of String) The roles of this role.
