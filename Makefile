runserver: 
	go run ./cmd/* runserver

test:
	go test -v ./...

test-sequentially:
	go test -v ./... -p 1

migrate-up:
	go run ./cmd/* migrate up

migrate-up-tests:
	go run ./cmd/* migrate --environment=test up

migrate-drop:
	go run ./cmd/* migrate drop

migrate-drop-tests:
	go run ./cmd/* migrate --environment=test drop

services-up: 
	docker-compose -f deployments/docker-compose.yml up -d

services-down: 
	docker-compose -f deployments/docker-compose.yml down

test-services-up:
	docker-compose -f deployments/docker-compose.yml up -d tests

test-coverage:
	go test -v ./... -coverprofile coverage.out -coverpkg ./... ./... &&  go tool cover -func coverage.out > coverage.txt && rm coverage.out
