app-build: Dockerfile
	docker build -t justdoit .

app-run: .env
	docker container run \
		--name justdoit \
		-p 8000:8080 \
		--network just-do-it_default \
		--env-file .env \
		--rm \
		-d \
		justdoit

app-down:
	docker container rm -f justdoit

db-up: compose.yml
	docker compose up -d

db-down: compose.yml
	docker compose down

db-migrate: .goose.env
	goose -env .goose.env up
