runserver: 
	go run ./cmd/* runserver

test:
	go test -v ./...

migrate-up:
	go run ./cmd/* migrate up

migrate-up-tests:
	go run ./cmd/* migrate--environment=test up

migrate-drop:
	go run ./cmd/* migrate drop

migrate-drop-tests:
	go run ./cmd/* migrate -environment=test drop

services-up: 
	docker-compose -f deployments/docker-compose.yml up -d

services-down: 
	docker-compose -f deployments/docker-compose.yml down

