# Payment app - Create App

resource "shipa_app" "payment" {
  app {
    name = "payment-services"
    teamowner = "shipa-team"
    framework = "cinema-payment"
  }
  depends_on = [shipa_app.notification]
}

# Payment app - Set env variables

resource "shipa_app_env" "paymentenv" {
  app = "payment-services"
  app_env {
    envs {
      name = "DB_SERVER"
      value = "xx.xx.xx.xx:27017"
    }
    envs {
      name = "DB_USER"
      value = "shipau"
    }
    envs {
      name = "DB_PASS"
      value = "shipapass"
    }
    envs {
      name = "DB"
      value = "booking"
    }
    envs {
      name = "STRIPE_SECRET"
      value = "sk_live_xxxxxxxxxxx"
    }
    envs {
      name = "STRIPE_PUBLIC"
      value = "pk_live_xxxxxxxxxxx"
    }
    norestart = true
    private = false
  }
  depends_on = [shipa_app.payment]
}


# Payment app - Deploy app

resource "shipa_app_deploy" "paymentdeploy" {
  app = "payment-services"
  deploy {
    image = "gcr.io/cosimages-206514/cinema-catalog-service@sha256:6613440a460e9f1e6e75ec91d8686c1aa11844b3e7c5413e241c807ce9829498"
  }
  depends_on = [shipa_app_env.paymentenv]
}
