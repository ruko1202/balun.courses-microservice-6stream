GOBIN ?= $(PWD)/bin
GORELEASER_VERSION = v2.8.2


SERVICE_NAME:=
ifndef ${SERVICE_NAME}
SERVICE_NAME:=$(shell cat ./.goreleaser.yaml | grep project_name | awk '{print $$2}')
endif
