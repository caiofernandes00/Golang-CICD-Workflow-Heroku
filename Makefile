######################## Application commands ########################
go_build:
	go build -o server ./app/cmd/server.go

go_run: go_build
	./server

######################## Docker compose commands ########################
docker_up:
	docker compose up --build --force-recreate

######################## PHONY ########################
.PHONY: docker_up go_build go_run