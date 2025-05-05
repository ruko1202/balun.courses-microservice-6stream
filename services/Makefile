compose-rebuild: compose-stop
	docker-compose up -d --build
	sleep 1
	docker-compose ps -a

compose-up: compose-stop
	docker-compose up -d
	sleep 1
	@make compose-ps

compose-ps:
	docker-compose ps -a

compose-stop:
	docker-compose stop -t0
	make docker-prune

docker-prune:
	$(info docker prune)
	docker container prune -f
	docker image prune -f

.curl: url=
.curl:
	@echo "Запрос к $$url"; \
	curl -w "\n%{response_code}\n" "$$url"

.check_service: service=
.check_service:
	@make .curl url="localhost:8001/$$service/version"
	@make .curl url="localhost:8001/$$service/readiness"
	@make .curl url="localhost:8001/$$service/liveness"

check_service:
	go test -v ./test/...



GOBIN ?= $(PWD)/bin
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
	@$(GOBIN)/golangci-lint run --disable-all -E goimports --fix ./...
