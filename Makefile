include ../../Makefile
DEV_DIR=../..
SERVICE=auth

.PHONY: build
build:
	$(call compose,build --no-cache && $(call down))

.PHONY: up
up:
	$(call compose,up auth && $(call down))

.PHONY: shell
shell:
	$(call run,bash)


.PHONY: test
test:
	$(call run,sh -c "ginkgo -vv ./...")

.PHONY: test_cov
test_cov:
	$(call run,sh -c "rm -f coverage.* && go test ./... -coverpkg=./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html")

.PHONY: fmt
fmt:
	$(call run,go fmt $$(go list ./...))