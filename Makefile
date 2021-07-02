HOSTNAME=shipa.io
NAMESPACE=terraform
NAME=shipa
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

TEST?=$(go list ./... | grep -v helpers)

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

testacc:
	TF_ACC=1 SHIPA_HOST=localhost SHIPA_TOKEN=123 go test $(TEST) -v -parallel 20 $(TESTARGS) -timeout 300m -cover -