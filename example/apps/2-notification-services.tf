# Notification app - Create App

resource "shipa_app" "notification" {
  app {
    name = "notification-service"
    teamowner = "shipa-team"
    framework = "cinema-services"
  }
  depends_on = [shipa_app.cinemacatalog]
}

# Notification app - Deploy app

resource "shipa_app_deploy" "notificationdeploy" {
  app = "movies-service"
  deploy {
    image = "gcr.io/cosimages-206514/cinema-catalog-service@sha256:6613440a460e9f1e6e75ec91d8686c1aa11844b3e7c5413e241c807ce9829498"
  }
  depends_on = [shipa_app.notification]
}
