# Movie services app - Create App

resource "shipa_app" "moviesapp" {
  app {
    name = "movies-service"
    teamowner = "shipa-team"
    framework = "cinema-services"
  }
}

# Movie services app - Set env variables

resource "shipa_app_env" "moviesenv" {
  app = "movies-service"
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
      value = "movies"
    }
    norestart = true
    private = false
  }
  depends_on = [shipa_app.moviesapp]
}


# Movies services app - Deploy app

resource "shipa_app_deploy" "appdeploy" {
  app = "movies-service"
  deploy {
    image = "gcr.io/cosimages-206514/movies-service@sha256:da99b1f332c0f07dfee7c71fc4d6d09cf6a26299594b6d1ae1d82d57968b3c57"
  }
  depends_on = [shipa_app_env.moviesenv]
}
