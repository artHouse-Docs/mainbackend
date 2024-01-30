package project

import (
	"context"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := r.body
}
