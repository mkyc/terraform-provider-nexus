---
page_title: "Resource nexus_security_saml"
subcategory: "Security"
description: |-
  ~> PRO Feature
  Use this resource to create a Nexus Security SAML configuration.
---
# Resource nexus_security_saml
~> PRO Feature

Use this resource to create a Nexus Security SAML configuration.
## Example Usage
```terraform
resource "nexus_security_saml" "example" {
  idp_metadata                  = "<EntityDescriptor ...>...</EntityDescriptor>"
  entity_id                     = "http://nexus.example/service/rest/v1/security/saml/metadata"
  validate_response_signature   = true
  validate_assertion_signature  = true
  username_attribute            = "username"
  first_name_attribute          = "firstName"
  last_name_attribute           = "lastName
  email_attribute               = "email
  groups_attribute              = "groups"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `idp_metadata` (String) SAML Identity Provider Metadata XML
- `username_attribute` (String) IdP field mappings for username

### Optional

- `email_attribute` (String) IdP field mappings for user's email address
- `entity_id` (String) Entity ID URI
- `first_name_attribute` (String) IdP field mappings for user's given name
- `groups_attribute` (String) IdP field mappings for user's groups
- `last_name_attribute` (String) IdP field mappings for user's family name
- `validate_assertion_signature` (Boolean) By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the assertions.
- `validate_response_signature` (Boolean) By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the response.

### Read-Only

- `id` (String) Used to identify resource at nexus
## Import
Import is supported using the following syntax:
```shell
# import of saml configuration
terraform import nexus_security_saml.example samle
```
