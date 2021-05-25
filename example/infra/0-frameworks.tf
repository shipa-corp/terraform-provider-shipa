# Creates cinema-services framework
resource "shipa_framework" "servicesfw" {
  framework {
    name = "cinema-services"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = false
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "5"
        }
        access {
          append = ["shipa-team"]
        }
        plan {
          name = "shipa-plan"
        }
        network_policy {
          ingress {
             policy_mode = "allow-all"
          }
          egress {
             policy_mode = "allow-all"
          }
          disable_app_policies = "false"
        }
      }
    }
  }
}

# Creates cinema-payment framework
resource "shipa_framework" "paymentfw" {
  framework {
    name = "cinema-payment"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = false
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "5"
        }
        access {
          append = ["shipa-team"]
        }
        plan {
          name = "shipa-plan"
        }
        network_policy {
          ingress {
             policy_mode = "allow-all"
          }
          egress {
             policy_mode = "allow-all"
          }
          disable_app_policies = "false"
        }
      }
    }
  }
}

# Creates cinema-ui framework
resource "shipa_framework" "uifw" {
  framework {
    name = "cinema-ui"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = false
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "5"
        }
        access {
          append = ["shipa-team"]
        }
        plan {
          name = "shipa-plan"
        }
        network_policy {
          ingress {
             policy_mode = "allow-all"
          }
          egress {
             policy_mode = "allow-all"
          }
          disable_app_policies = "false"
        }
      }
    }
  }
}

# Creates cinema-gateway framework
resource "shipa_framework" "gatewayfw" {
  framework {
    name = "cinema-gateway"
    provisioner = "kubernetes"
    resources {
      general {
        setup {
          public = false
          default = false
        }
        security {
          disable_scan = true
          scan_platform_layers = false
        }
        router = "traefik"
        app_quota {
          limit = "5"
        }
        access {
          append = ["shipa-team"]
        }
        plan {
          name = "shipa-plan"
        }
        network_policy {
          ingress {
             policy_mode = "allow-all"
          }
          egress {
             policy_mode = "allow-all"
          }
          disable_app_policies = "false"
        }
      }
    }
  }
}

resource "null_resource" "resource-to-wait-on" {
  provisioner "local-exec" {
    command = "sleep 30"
  }
}
