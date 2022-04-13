runserver: 
	go run ./cmd/* runserver

test:
	go test -v ./...

migrate_up:
	go run ./cmd/* migrate up

migrate_up_tests:
	go run ./cmd/* migrate--environment=test up

migrate_drop:
	go run ./cmd/* migrate drop

migrate_drop_tests:
	go run ./cmd/* migrate -environment=test drop

services-up: 
	docker-compose -f deployments/docker-compose.yml up

services-down: 
	docker-compose -f deployments/docker-compose.yml down

