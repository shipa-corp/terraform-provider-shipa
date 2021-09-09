HOSTNAME=shipa.io
NAMESPACE=terraform
NAME=shipa
BINARY=terraform-provider-${NAME}
VERSION=0.0.2
#OS_ARCH=darwin_amd64
OS_ARCH=linux_amd64
default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
