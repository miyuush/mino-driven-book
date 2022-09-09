GOTOOLS = \
	github.com/ramya-rao-a/go-outline@latest \
	github.com/cweill/gotests/gotests@latest \
	github.com/fatih/gomodifytags@latest \
	github.com/josharian/impl@latest \
	github.com/haya14busa/goplay/cmd/goplay@latest \
	github.com/go-delve/delve/cmd/dlv@latest \
	honnef.co/go/tools/cmd/staticcheck@latest \
	golang.org/x/tools/gopls@latest \
	golang.org/x/tools/cmd/goimports@latest \
    github.com/golangci/golangci-lint/cmd/golangci-lint@latest \

.PHONY: install-go-tools
install-go-tools:
		@for t in $(GOTOOLS); do \
		go install $$t; \
		done
