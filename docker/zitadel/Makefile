zitadel-compose-up:
	docker-compose -f zitadel-docker-compose.yaml up -d

postgres:
	@echo "Running postgres container on port 5433..."
	@docker run --name rh-postgres -e POSTGRES_PASSWORD=postgres -p 5433:5432 -d postgres
	@echo "Done."

cache-database:
	@echo "Running cache database container on port 6379..."
	@docker run -d -p 6379:6379 redis
	@echo "Done."
