PORT=8020 PSQL_URI=postgres://localhost/secondbrain_test?sslmode=disable go test -covermode=count -coverprofile=coverage.out
go tool cover -html=coverage.out

