SHIPA_LOCAL_PROVIDER_NAMESPACE=terraform.local/local/shipa
VERSION=0.0.10
BINARY=terraform-provider-shipa_v${VERSION}
default: install

build:
	go build -o ${BINARY}

fmt:
	go fmt ./...

test:
	go test -race $$(go list ./... | grep -v /vendor/)

install: GOOS=$(shell go env GOOS)
install: GOARCH=$(shell go env GOARCH)
install: DESTINATION=$(HOME)/.terraform.d/plugins/$(SHIPA_LOCAL_PROVIDER_NAMESPACE)/$(VERSION)/$(GOOS)_$(GOARCH)
install: build
	mkdir -p ${DESTINATION}
	mv ${BINARY} ${DESTINATION}
