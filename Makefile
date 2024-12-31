lint:
	revive -config ./linter/revive.toml -formatter friendly ./...

watch:
	templ generate --watch --proxy="http://localhost:8082" --cmd="make run"

tailwind-init:
	npx tailwindcss init

tailwind:
	tailwindcss -i web/public/css/input.css -o web/public/css/output.css

build:
	@echo "Building..."
	@templ generate
	@make tailwind
	@go build -o main .

run:
	@echo "Running..."
	@make tailwind
	@go run .

ent-describe:
	@echo "Describing ent..."
	@go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
	@echo "Saving the schema to /ent-describe.md"
	@go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema > ent-describe.md
	@echo "Saved."

ent-generate:
	@echo "Generating ent..."
	@cd ent && go generate .

postgres:
	@echo "Running postgres container on port 5433..."
	@docker run --name rh-postgres -e POSTGRES_PASSWORD=postgres -p 5433:5432 -d postgres
	@echo "Done."

cache-database:
	@echo "Running cache database container on port 6379..."
	@docker run -d -p 6379:6379 redis
	@echo "Done."
