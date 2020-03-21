GO := $(shell which go)
GO112_VERSION := 1.12.16
GO112 := go${GO112_VERSION}

GOPHERJS := $(shell which gopherjs 2> /dev/null)

## start-web       : Starts GopherJS serving.
start-web:
	GOPHERJS_GOROOT="${HOME}/sdk/${GO112}" ${GOPHERJS} serve . -v

## prepare         : Installs Go 1.12 and GopherJS. Go must be already installed.
prepare: prepare-go prepare-gopherjs

## prepare-go      : Prepares specific Go tooling for GopherJS. Go must be already installed.
prepare-go:
	@echo "* Getting Go 1.12.16 (required by gopherjs)"
	@${GO} get golang.org/dl/${GO112}
	@${GO112} download

## prepare-gopherjs: Prepares GopherJS (e.g. compiles it).
prepare-gopherjs:
	@echo "* Getting gopherjs..."
	@${GO112} get -u -v github.com/gopherjs/gopherjs

help : Makefile
	@echo "helloworld Makefile available subcommands:\n"
	@sed -n 's/^## //p' $<

.DEFAULT_GOAL := help