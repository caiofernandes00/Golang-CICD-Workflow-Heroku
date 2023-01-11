######################## Application commands ########################
go_build:
	go build -o main ./src/cmd

go_run: go_build
	./main

######################## Docker compose commands ########################
docker_up:
	docker compose up --build --force-recreate

######################## PHONY ########################
.PHONY: docker_up go_build go_run