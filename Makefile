sqlc: 
	sqlc generate
run:
	go run main.go
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fastprint?sslmode=disable" -verbose down