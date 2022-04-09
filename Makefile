runserver: 
	go run ./cmd/* runserver

migrate_up:
	go run ./cmd/* migrate up

migrate_drop:
	go run ./cmd/* migrate drop

services-up: 
	docker-compose -f deployments/docker-compose.yml up

services-down: 
	docker-compose -f deployments/docker-compose.yml down

