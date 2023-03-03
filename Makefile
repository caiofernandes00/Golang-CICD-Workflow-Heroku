######################## Application ########################
app_build:
	go build -o main ./app/cmd/main.go

app_run: app_build
	./main

app_test:
	go test -v ./app/...

######################## Docker ########################
docker_up:
	docker compose up --build --force-recreate

######################## Docks ########################
gen_doc:
	echo "Starting swagger generating"
	swag init -dir ".\app\cmd,.\app\infrastructure\api"

######################## PHONY ########################
.PHONY: docker_up go_build go_run