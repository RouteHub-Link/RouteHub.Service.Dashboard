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