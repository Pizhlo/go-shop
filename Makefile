migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8081/shop?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8081/shop?sslmode=disable" -verbose down

.PHONY: migrateup migratedown