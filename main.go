package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &RecipeSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/gravityql", h)
	http.ListenAndServe(":8080", nil)
}
