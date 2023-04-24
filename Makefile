#Define variables for the migration
MIGRATE_CMD=migrate
MIGRATE_PATH=internal/adapter/repository/psql/migration
MIGRATE_DATABASE=postgresql://admin:secret@localhost:5432/todo?sslmode=disable
MIGRATE_VERBOSE=-verbose
MIGRATE_UP=up

#target
migrate:
	$(MIGRATE_CMD) -path $(MIGRATE_PATH) -database $(MIGRATE_DATABASE) $(MIGRATE_VERBOSE) $(MIGRATE_UP)

schema:
	$(MIGRATE_CMD) create -ext postgresql -dir $(MIGRATE_PATH) -seq init_schema

