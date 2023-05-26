ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=postgres password=12345 dbname=hw-5 host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/internal/db/migrations

.PHONY: migration-create migration-up migration-down db-version
#Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND

#Creates new migration file with the current timestamp
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

# Migrate the DB to the most recent version available
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

#Roll back the version by 1
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

db-version:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" version
