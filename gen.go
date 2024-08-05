package gen

//go:generate go run ./cmd/tools/turndotenv/main.go
//go:generate ./internal/store/pgstore/sqlc.exe generate -f ./internal/store/pgstore/sqlc.yaml
