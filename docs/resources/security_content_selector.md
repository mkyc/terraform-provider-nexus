---
page_title: "Resource nexus_security_content_selector"
subcategory: "Security"
description: |-
  Use this resource to create a Nexus Content Selector.
---
# Resource nexus_security_content_selector
Use this resource to create a Nexus Content Selector.
## Example Usage
```terraform
resource "nexus_security_content_selector" "example" {
	name        = "example"
	description = "example content selector"
	expression  = "format == \"raw\""
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `expression` (String) The content selector expression
- `name` (String) Content selector name

### Optional

- `description` (String) A description of the content selector

### Read-Only

- `id` (String) Used to identify resource at nexus
## Import
Import is supported using the following syntax:
```shell
# import using the name of the content selector
terraform import nexus_security_content_selector.example example
```
