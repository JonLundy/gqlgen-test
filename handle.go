package test

import (
	"net/http"
	"github.com/vektah/gqlgen/handler"
	"test/graph"
	"test/resolver"
)

func GetQuery(w http.ResponseWriter, r *http.Request) {
	h := handler.GraphQL(graph.NewExecutableSchema(&resolver.Root{}))

	h.ServeHTTP(w, r)
}