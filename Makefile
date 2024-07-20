include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := test

.PHONY: update
update: ## pulls git repo and installs all dependencies
	@git pull

.PHONY: test
test: ## start via docker
	go test ./... -v