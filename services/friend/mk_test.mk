.PHONY: test
test:
	go test -race -p 2 -count 2 ./...

.PHONY: test-cov
ci-test-with-coverage:
	go test -race -coverprofile=coverage.tmp -covermode atomic --coverpkg=./internal/... ./...
	@grep -v "mock" coverage.tmp > coverage.out
	go tool cover -func=coverage.out

