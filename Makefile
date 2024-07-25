include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
IMAGE_NAME := "routriever"
IMAGE_VERSION := "1.0.0"

# get the lastest version of prometheus operator here:
# https://github.com/prometheus-operator/prometheus-operator
PROMETHEUS_VERSION := "0.75.2"

.DEFAULT_GOAL := start

.PHONY: update
update: ## pulls git repo and installs all dependencies
	@git -C ${ROOT_DIR} pull

.PHONY: test
test: ## test service
	@go test ${ROOT_DIR}...

.PHONY: start
start: ## build and starts the service via docker
	@docker compose -f ${ROOT_DIR}compose.yaml up --build

.PHONY: generate-helm-docs
generate-helm-docs: ## re-generates helm docs using docker
	@docker run --rm --volume "$(ROOT_DIR)/charts:/helm-docs" jnorwood/helm-docs:latest

.PHONY: start-cluster
start-cluster: # starts k3d cluster and registry
	@k3d cluster create --config ${ROOT_DIR}k3d/clusterconfig.yaml 
# @helm install go-mail-service --set service.port=$(API_PORT) \

.PHONY: k3d-start
k3d-start: start-cluster k3d-push ## starts k3d, registry, and deploys promtheus CRDs and installs local helm chart, with local image
	@kubectl create -f https://raw.githubusercontent.com/coreos/prometheus-operator/v${PROMETHEUS_VERSION}/bundle.yaml
	@helm install ${IMAGE_NAME} ${ROOT_DIR}charts/${IMAGE_NAME} --set image.repository=localhost:5000/${IMAGE_NAME} --set image.tag=${IMAGE_VERSION}

.PHONY: k3d-stop
k3d-stop: ## stop K3d
	@k3d cluster delete --config ${ROOT_DIR}k3d/clusterconfig.yaml

.PHONY: k3d-restart
k3d-restart: k3d-stop k3d-start ## restart k3d

.PHONY: k3d-push
k3d-push: # build and push docker image to local registry
	@docker build -f ${ROOT_DIR}Dockerfile . -t ${IMAGE_NAME}
	@docker tag ${IMAGE_NAME} localhost:5000/${IMAGE_NAME}:${IMAGE_VERSION}
	@docker push localhost:5000/${IMAGE_NAME}:${IMAGE_VERSION}