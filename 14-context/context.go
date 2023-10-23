package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context context.Context) (error, string)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err, data := store.Fetch(r.Context())

		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}
