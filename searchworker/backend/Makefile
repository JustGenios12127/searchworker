.PHONY: up down build logs migrate

up:
	docker compose up --build -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

migrate:
	docker compose run --rm worker-backend /app/server -migrate
