include mk_vars.mk

.PHONY: .curl
.curl: path=
.curl:
	curl -s -w "\nStatus code: %{response_code}\n" localhost:8080/${path}

liveness:
	@make .curl path="liveness"

readiness:
	@make .curl path="readiness"

version:
	@make .curl path="version"

