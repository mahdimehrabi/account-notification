unit_test:
	go test ./... -v

fmt:
	gofumpt -l -w .

devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

check:
	golangci-lint run \
		--build-tags "${BUILD_TAG}" \
		--timeout=20m0s \
		--enable=gofmt \
		--enable=unconvert \
		--enable=unparam \
		--enable=asciicheck \
		--enable=misspell \
		--enable=revive \
		--enable=decorder \
		--enable=reassign \
		--enable=usestdlibvars \
		--enable=nilerr \
		--enable=gosec \
		--enable=exportloopref \
		--enable=whitespace \
		--enable=goimports \
		--enable=gocyclo \
		--enable=nestif \
		--enable=gochecknoinits \
		--enable=gocognit \
		--enable=funlen \
		--enable=forbidigo \
		--enable=godox \
		--enable=gocritic \
		--enable=gci \
		--enable=lll



build:
	docker build -t nexus.infra.io:8443/xfolio/basicinfo:release .

push:
	docker push nexus.infra.io:8443/xfolio/basicinfo:release