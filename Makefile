SHELL         := /bin/bash
.SHELLFLAGS   := -eu -o pipefail -c
.DEFAULT_GOAL := help

.ONESHELL:
.PHONY: help update_index

## Show this help message
help:
	@echo "Makefile targets:"
	@echo ""
	@awk '/^## /{desc = substr($$0, 4)} /^[a-zA-Z_-]+:/{if (desc) {target = $$1; sub(/:.*/, "", target); printf "  make %-20s- %s\n", target, desc; desc = ""}}' $(MAKEFILE_LIST)
	@echo ""

## Updates index.md so GH Pages shows always up to date content
update_index:
