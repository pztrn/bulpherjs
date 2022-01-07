GOPHERJS := $(shell which gopherjs 2> /dev/null)

## start-web       : Starts GopherJS serving.
start-web:
	${GOPHERJS} serve . -v

help : Makefile
	@echo "helloworld Makefile available subcommands:\n"
	@sed -n 's/^## //p' $<

.DEFAULT_GOAL := help
