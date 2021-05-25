terraform {
  required_providers {
    shipa = {
      version = "0.0.1"
      source = "shipa.io/terraform/shipa"
    }
  }
}

provider "shipa" {}

# Returns all apps
data "shipa_apps" "all" {}

output "all_apps" {
  value = data.shipa_apps.all.apps
}

# Returns specific app
data "shipa_app" "app" {
  id = "terraform-app-3"
}

output "app" {
  value = data.shipa_app.app
}

# Returns specific framework
data "shipa_framework" "f1" {
  id = "shipa-framework"
}

output "my_val" {
  value = data.shipa_framework.f1
}

# Returns all frameworks
data "shipa_frameworks" "all" {}

output "all_frameworks" {
  value = data.shipa_frameworks.all.frameworks
}

# Create framework
resource "shipa_framework" "tf" {
  framework {
    name = "test-tf-144"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = true
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "2"
        }
      }
    }
  }
}

resource "shipa_framework" "tf2" {
  framework {
    name = "test-tf-1442"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = true
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "2"
        }
      }
    }
  }
}

output "my_tf_framework" {
  value = shipa_framework.tf
}

# Create cluster
resource "shipa_cluster" "tf" {
  cluster {
    name = "cl-1"
    endpoint {
      addresses = ["https://k8s-api.com:443"]
      ca_cert = "<<ca_cert here>>"
      token = "<<token goes here>>"
      client_key = "key"
    }
    resources {
      frameworks {
        name = [
          "${shipa_framework.tf.framework[0].name}",
          "${shipa_framework.tf2.framework[0].name}",
          ]
      }
      ingress_controllers {
        type = "traefik"
        debug = false
        acme_email = "acme@email.com"
        acme_server = "127.0.0.1"
      }
    }
  }
}

# Create user
resource "shipa_user" "tf" {
  email = "user@terraform.com"
  password = "terraform"
}

# Create plan
resource "shipa_plan" "tf" {
  plan {
    name = "terraform-plan"
    teams = ["dev"]
    cpushare = 4
    memory = 0
    swap = 0
  }
}

# Create team
resource "shipa_team" "t2" {
  team {
    name = "terraform-team-2"
    tags = ["dev", "sandbox"]
  }
}

# Create app
resource "shipa_app" "app2" {
  app {
    name = "terraform-app-3"
    teamowner = "dev"
    framework = "test-tf-142"
    description = "test description update"
    tags = ["dev", "sandbox", "updated"]
  }
}

# Set app envs
resource "shipa_app_env" "env1" {
  app = "terraform-app-1"
  app_env {
    envs {
      name = "SHIPA_ENV1"
      value = "test-1"
    }
    envs {
      name = "SHIPA_ENV2"
      value = "test-2"
    }
    norestart = true
    private = false
  }
}

# Set app cname
resource "shipa_app_cname" "cname1" {
  app = "terraform-app-2"
  cname = "test.com"
  encrypt = true
}

# Deploy app
resource "shipa_app_deploy" "deploy1" {
  app = "terraform-app-2"
  deploy {
    image = "docker.io/shipasoftware/bulletinboard:1.0"
    // optional
    private_image = true
    registry_user = "Test user"
    registry_secret = "Test Secret"
    steps = 2
    step_weight = 1
    step_interval = 1
    port = 8000
    detach = true
    message = "Message"
  }
}

# Create role
resource "shipa_role" "role1" {
  name = "RoleName"
  context = "app"
  // Optional
  description = "test"
}

# Create role permission
resource "shipa_permission" "p1" {
  name = "RoleName"
  permission = ["app.read",  "app.deploy"]
}

# Create role association
resource "shipa_role_association" "r1" {
  name = "RoleName"
  email = "terraform@terraform.com"
}
