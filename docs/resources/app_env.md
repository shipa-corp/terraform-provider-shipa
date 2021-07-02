---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "shipa_app_env Resource - terraform-provider-shipa"
subcategory: ""
description: |-
  
---

# shipa_app_env (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **app** (String)
- **app_env** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--app_env))

### Optional

- **id** (String) The ID of this resource.

<a id="nestedblock--app_env"></a>
### Nested Schema for `app_env`

Required:

- **envs** (Block List, Min: 1) (see [below for nested schema](#nestedblock--app_env--envs))

Optional:

- **norestart** (Boolean)
- **private** (Boolean)

<a id="nestedblock--app_env--envs"></a>
### Nested Schema for `app_env.envs`

Required:

- **name** (String)
- **value** (String)

