package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store is an interface
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server is the http server function
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}
