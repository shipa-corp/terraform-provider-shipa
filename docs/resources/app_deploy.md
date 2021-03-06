---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "shipa_app_deploy Resource - terraform-provider-shipa"
subcategory: ""
description: |-
  
---

# shipa_app_deploy (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app` (String)
- `deploy` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--deploy))

### Optional

- `id` (String) The ID of this resource.

<a id="nestedblock--deploy"></a>
### Nested Schema for `deploy`

Required:

- `image` (String)

Optional:

- `description` (String)
- `detach` (Boolean)
- `env` (List of String)
- `framework` (String)
- `message` (String)
- `origin` (String)
- `port` (Number)
- `private_image` (Boolean)
- `protocol` (String)
- `registry_secret` (String)
- `registry_user` (String)
- `shipa_yaml` (String)
- `step_interval` (Number)
- `step_weight` (Number)
- `steps` (Number)
- `tags` (List of String)
- `team` (String)

Read-Only:

- `plan` (List of Object) (see [below for nested schema](#nestedatt--deploy--plan))
- `router` (List of Object) (see [below for nested schema](#nestedatt--deploy--router))

<a id="nestedatt--deploy--plan"></a>
### Nested Schema for `deploy.plan`

Read-Only:

- `cpushare` (Number)
- `default` (Boolean)
- `memory` (Number)
- `name` (String)
- `org` (String)
- `public` (Boolean)
- `swap` (Number)
- `teams` (List of String)


<a id="nestedatt--deploy--router"></a>
### Nested Schema for `deploy.router`

Read-Only:

- `address` (String)
- `default` (Boolean)
- `name` (String)
- `opts` (Map of String)
- `type` (String)


