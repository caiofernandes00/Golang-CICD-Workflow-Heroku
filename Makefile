######################## Application ########################
app_build:
	go build -o main ./app/cmd/main.go

app_run: app_build
	./main

app_test:
	go test -v ./app/...

app_lint:
	golangci-lint run ./app/...

######################## Mocks ########################
CIRCUIT_BREAKER_PATH = ./app/infrastructure/resilience/observable/circuitbreaker/observable.go
CIRCUIT_BREAKER_PATH_MOCKGEN = ./app/infrastructure/resilience/observable/circuitbreaker/mock

mockgen:
	@for file in $(CIRCUIT_BREAKER_PATH); do \
		mockgen -source=$$file -destination=$(CIRCUIT_BREAKER_PATH_MOCKGEN)/`basename $$file` ; \
	done

######################## Docker ########################
docker_up:
	docker compose up --build --force-recreate

######################## Docks ########################
gen_doc:
	echo "Starting swagger generating"
	swag init -dir ".\app\cmd,.\app\infrastructure\api"

######################## PHONY ########################
.PHONY: docker_up go_build go_run