GO_BIN = go
ifneq (${GO},)
	GO_BIN = ${GO}
endif

.PHONY: help init fmt test

help: Makefile
	@printf "${BLUE}Choose a command run:${NC}\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/    /'

## make init: Install dependencies
init:
	${GO_BIN} install go.uber.org/mock/mockgen@main
	${GO_BIN} install github.com/fsgo/go_fmt/cmd/gorgeous@latest

## make fmt: Formats source code
fmt:
	gofumpt -l -w .
	goimports -local github.com/zunkk -w .

## make test: Run go unittest
test:
	${GO_BIN} test -timeout 300s ./... -count=1
