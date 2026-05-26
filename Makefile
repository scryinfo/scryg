
ifeq ($(OS),Windows_NT)
	ifneq (,$(findstring MINGW,$(UNAME_S)))
	    go := ${shell which go}
	else
	    go := $(subst \,/,$(shell where.exe go))
	endif
else
	go := ${shell which go}
endif
$(info "go: ${go}")

.PHONY: clean upgrade format build samples

clean:
	rm -rf go.sum go.work.sum demo/go.sum
	${go} clean
tidy:
	${go} mod tidy
upgrade:
	${go} get -t -u ./... && ${go} mod tidy
format:
	${go} fmt ./...
build:
	${go} build ./...
rebuild: clean build
go_fix:
	${go} fix ./...
lint:
	${go} vet ./...
	govulncheck ./...
	staticcheck ./...
	nilaway ./...
	golangci-lint run --no-config --disable-all -E zerologlint ./...
lint_more:
	gosec -tags "Release" -quiet ./...
	revive ./...
	gocyclo ./...
	golangci-lint run ./...

go_tools:
	${go} install github.com/fatih/gomodifytags@latest
	${go} install golang.org/x/tools/gopls@latest
	${go} install honnef.co/go/tools/cmd/staticcheck@latest
	${go} install github.com/cweill/gotests/gotests@latest
	${go} install github.com/josharian/impl@latest
	${go} install github.com/go-delve/delve/cmd/dlv@latest
	${go} install go.uber.org/nilaway/cmd/nilaway@latest
	${go} install golang.org/x/vuln/cmd/govulncheck@latest
	${go} install github.com/securego/gosec/v2/cmd/gosec@latest
	${go} install github.com/mgechev/revive@latest
	${go} install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	${go} install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
