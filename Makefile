include .env

postgres:
    docker run --name $(POSTGRES_CONTAINER_NAME) -p $(DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres

create_db:
	docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_NAME)

drop_db:
	docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb $(DB_NAME)

migrate_up:
	migrate -path db/migrations -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down