include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := start

.PHONY: update
update: ## pulls git repo and installs all dependencies
	@git -C ${ROOT_DIR} pull

.PHONY: test
test: ## test service
	@go test ${ROOT_DIR}...

.PHONY: start-docker
start-docker: ## build and starts the service via docker
	@docker compose -f ${ROOT_DIR}compose.yaml up --build

.PHONY: start
start: ## starts the service
	@go run ${ROOT_DIR}main.go