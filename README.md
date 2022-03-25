# shipa-terraform-provider

## Building and testing

1. Build terraform provider by using:

```bash
    make install
```

1. Set env values:

```bash
    export SHIPA_HOST=http://target.shipa.cloud:8080
    export SHIPA_TOKEN=xxxxxxxxx
```

1. Run terraform

```bash
    cd example
    terraform init && terraform apply --auto-approve  
```

## Updating documentation

You need to install [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs/releases) and then run it from the root directory of the repository.

```bash
TFPLUGINDOCS_VERSION=0.7.0
OS="darwin"
ARCH="amd64"

mkdir -p /tmp/tfplugindocs
pushd /tmp/tfplugindocs
curl -sSLo tfplugindocs.zip https://github.com/hashicorp/terraform-plugin-docs/releases/download/v${TFPLUGINDOCS_VERSION}/tfplugindocs_${TFPLUGINDOCS_VERSION}_${OS}_${ARCH}.zip
unzip tfplugindocs.zip
sudo mv tfplugindocs /usr/local/bin
popd

tfplugindocs
```
