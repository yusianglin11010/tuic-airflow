dev:
	@echo "run docker postgres"
	@docker run  -d \
		-p 5432:5432 \
		--name postgres \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_SSLMODE=disable \
		postgres