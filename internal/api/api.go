package api

import "github.com/VitorHFernandes/ama-go/internal/store/pgstore"

type apiHandler struct {
	q *pgstore.Queries
}
