include mk_vars.mk
include mk_build.mk
include mk_lint.mk
include mk_test.mk
include mk_docker.mk
include mk_curls.mk


# for local debug uncomment the following rows
#include .env.local
#export

.PHONY: run
run:
	go run ./cmd/service
