---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cognito_user_pool_resource_server Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Cognito::UserPoolResourceServer
---

# awscc_cognito_user_pool_resource_server (Resource)

Resource Type definition for AWS::Cognito::UserPoolResourceServer



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String)
- `name` (String)
- `user_pool_id` (String)

### Optional

- `scopes` (Attributes List) (see [below for nested schema](#nestedatt--scopes))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--scopes"></a>
### Nested Schema for `scopes`

Required:

- `scope_description` (String)
- `scope_name` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cognito_user_pool_resource_server.example <resource ID>
```