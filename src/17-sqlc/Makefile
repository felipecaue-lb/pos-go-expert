createMigration:
	migrate create -ext=sql -dir=sql/migrations --seq init

migrateUp:
	migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose up

migrateDown:
	migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose down

.PHONY: migrateUp migrateDown createMigration