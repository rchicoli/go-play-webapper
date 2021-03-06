package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIDKey = key(42)

// Printf adds the request from the event context.
func Printf(ctx context.Context, msg interface{}) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Printf("error=unknown msg=%s", msg)
		return
	}
	log.Printf("id=%d msg=%s", id, msg)
}

// Decorate adds a unique request id to the context of the request.
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))

		log.Printf("id=%d path=%s", id, r.URL.EscapedPath())
	}
}
