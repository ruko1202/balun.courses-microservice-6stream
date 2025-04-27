include mk_vars.mk

GOLANGCI_VERSION = v1.64.5

.PHONY: lint
lint: install-lint
	$(info $(M) running linter...)
	@$(GOBIN)/golangci-lint run ./...

.PHONY: install-lint
install-lint:
ifeq ($(wildcard ${GOBIN}/golangci-lint),)
	$(info $(M) install golangci-lint $(GOLANGCI_VERSION)...)
	@GOBIN=$(GOBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
else
ifneq ($(GOLANGCI_VERSION),$(shell ${GOBIN}/golangci-lint --version | awk '{print $$4}'))
	$(info $(M) update golangci-lint $(GOLANGCI_VERSION)...)
	@GOBIN=$(GOBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION)
endif
endif
	$(GOBIN)/golangci-lint --version


.PHONY: fmt
fmt: install-lint ## Format code
	$(info $(M) fmt project...)
	make update_readme
	@$(GOBIN)/golangci-lint run --disable-all -E goimports --fix ./...
