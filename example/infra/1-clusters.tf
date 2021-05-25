# Creates cluster for cinema services
resource "shipa_cluster" "servicescl" {
  cluster {
    name = "cinema-services-cl"
    endpoint {
      addresses = ["https://k8s-api.com:443"]
      ca_cert = <<-EOT
	  -----BEGIN CERTIFICATE-----
	  -----END CERTIFICATE-----
      EOT
      token = "xxxxxxx"
    }
    resources {
      frameworks {
        name = ["cinema-services", "cinema-payment", "cinema-gateway"]
      }
    }
  }
  depends_on = [shipa_framework.servicesfw]
}

# Creates cluster for cinema ui
resource "shipa_cluster" "uicl" {
  cluster {
    name = "cinema-ui-cl"
    endpoint {
      addresses = ["https://k8s-api.com:443"]
      ca_cert = <<-EOT
	  -----BEGIN CERTIFICATE-----
	  -----END CERTIFICATE-----
      EOT
      token = "xxxx"
    }
    resources {
      frameworks {
        name = ["cinema-ui"]
      }
    }
  }
  depends_on = [shipa_framework.servicesfw]
}
