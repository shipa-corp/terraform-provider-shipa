terraform {
  required_providers {
    shipa = {
      version = "0.0.1"
      source = "shipa.io/terraform/shipa"
    }
  }
}

provider "shipa" {}

# Returns specific app
data "shipa_app" "app" {
  id = shipa_app.tf.app[0].name
}

output "app" {
  value = data.shipa_app.app
}

# Returns specific framework
data "shipa_framework" "tf" {
  id = shipa_framework.tf.framework[0].name
}

output "my_val" {
  value = data.shipa_framework.tf
}

# Create framework
resource "shipa_framework" "tf" {
  framework {
    name = "f1"
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
        access {
          append = [shipa_team.tf.team[0].name]
        }
        plan {
          name = shipa_plan.tf.plan[0].name
        }
      }
    }
  }
  depends_on = [shipa_plan.tf, shipa_team.tf]
}

output "my_tf_framework" {
  value = shipa_framework.tf
}

# Create cluster
resource "shipa_cluster" "tf" {
  cluster {
    name = "cl-1"
    endpoint {
      addresses = ["https://E873062014D66818F6B1E5E19AD27DF0.gr7.us-west-1.eks.amazonaws.com"]
      ca_cert = file("path/to/cluster/ca-cert")
      token = "cluster-token-value" //FIXME: file("path/to/cluster/ca-cert") has a bug
    }
    resources {
      frameworks {
        name = [
          shipa_framework.tf.framework[0].name,
          ]
      }
    }
  }
  depends_on = [shipa_framework.tf]
}

# Create user
resource "shipa_user" "tf" {
  email = "user@example.com"
  password = "terraform"
}

# Create plan
resource "shipa_plan" "tf" {
  plan {
    name = "terraform-plan"
    teams = [shipa_team.tf.team[0].name]
    cpushare = 4
    memory = 0
    swap = 0
  }
  depends_on = [shipa_team.tf]
}

# Create team
resource "shipa_team" "tf" {
  team {
    name = "dev"
    tags = ["dev"]
  }
}

# Create app
resource "shipa_app" "tf" {
  app {
    name = "app1"
    teamowner = shipa_team.tf.team[0].name
    framework = shipa_framework.tf.framework[0].name
    description = "test description update"
    tags = ["dev"]
  }
  depends_on = [shipa_cluster.tf]
}

# Set app envs
resource "shipa_app_env" "tf" {
  app = shipa_app.tf.app[0].name
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
resource "shipa_app_cname" "tf" {
  app = shipa_app.tf.app[0].name
  cname = "test.com"
  encrypt = true
}

# Deploy app
resource "shipa_app_deploy" "tf" {
  app = shipa_app.tf.app[0].name
  deploy {
    image = "docker.io/shipasoftware/bulletinboard:1.0"
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
  email = shipa_user.tf.email
}