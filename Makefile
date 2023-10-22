createmigrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3307)/gokapster?multiStatements=true" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3307)/gokapster?multiStatements=true" -verbose down

run:
	go run cmd/main.go

test :
	go test -v -cover ./...

.PHONY: createmigrate migrateup migratedown run