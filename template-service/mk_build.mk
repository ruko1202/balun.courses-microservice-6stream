include mk_vars.mk

GORELEASER_VERSION = v2.8.2


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
