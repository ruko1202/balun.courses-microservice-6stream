include mk_vars.mk

GORELEASER_VERSION = v2.8.2

.PHONY: bin-deps
bin-deps:
	$(info Installing bin dependencies...)
	GOBIN=$(GOBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6  && \
	GOBIN=$(GOBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1  && \
	GOBIN=$(GOBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.3  && \
	GOBIN=$(GOBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.26.3  && \
	GOBIN=$(GOBIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.6.0  && \
	GOBIN=$(GOBIN) go install github.com/bufbuild/buf/cmd/buf@v1.53.0

.PHONY: build
build: install-goreleaser
	# see .build args in .goreleaser.yaml
	# --single-target      Builds only for current GOOS and GOARCH, regardless of what's set in the configuration file
	$(GOBIN)/goreleaser build --snapshot --single-target --clean --output=$(GOBIN)/${SERVICE_NAME} ${args}


install-goreleaser:
ifeq ($(wildcard ${GOBIN}/goreleaser),)
	$(info $(M) install goreleaser $(GORELEASER_VERSION)...)
	@GOBIN=$(GOBIN) go install github.com/goreleaser/goreleaser/v2@$(GORELEASER_VERSION)
else
ifneq ($(shell echo ${GORELEASER_VERSION} | sed 's/^v//'),$(shell ${GOBIN}/goreleaser --version | grep GitVersion | awk '{print $$2}'| sed 's/^v//'))
	$(info $(M) update goreleaser $(GORELEASER_VERSION)...)
	@GOBIN=$(GOBIN) go install github.com/goreleaser/goreleaser/v2@$(GORELEASER_VERSION)
endif
endif
	$(GOBIN)/goreleaser --version

generate:
	$(GOBIN)/buf generate

.PHONY: deps
deps:
	go mod tidy
