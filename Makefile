ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

WORKDIR='/workspaces/unibee-api'
dotUser='$(shell basename "$$HOME")'
VERSION ?= dev
BINARY := unibee-api-$(VERSION)
container='unibee_api'

bash: ## SSH into the container with bash
	@docker exec -it $(container) bash

zsh: ## SSH into the container with zsh
	@docker exec -it $(container) zsh

down-volumes: ## down devcontainer and remove orphans and volumes
	@docker compose -f .devcontainer/docker/docker-compose.yml down -v --remove-orphans

expose: ## share the container to the outside world with expose
	@docker exec -w $(WORKDIR) -it $(container) /bin/bash $(WORKDIR)/.devcontainer/docker/install-php-cli.sh
	@docker exec -w $(WORKDIR) -it $(container) .devcontainer/dotdev/bin/expose share "http://localhost:$(APP_PORT)" \
                                                                         --server="https://dotshare.dev" \
                                                                         --server-host="dotshare.dev" \
                                                                         --server-port="443" \
                                                                         --auth="admin:dotworld-test-admin" \
                                                                         --subdomain=$(dotUser)-$(container)

down: ## down devcontainer and remove orphans
	@docker compose -f .devcontainer/docker/docker-compose.yml down --remove-orphans

serve: ## start the server
	@docker exec -w $(WORKDIR) -it $(container) /root/go/bin/air -c $(WORKDIR)/.air.toml

run: ## up devcontainer
	@devcontainer up --workspace-folder .