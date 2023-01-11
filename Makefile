######################## Docker compose commands ########################
docker_up:
	docker compose up --build --force-recreate

######################## PHONY ########################
.PHONY: docker_up