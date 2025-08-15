package main

import (
	"fmt"
	"go-service/gql"
	"go-service/service"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/denisenkom/go-mssqldb"
)

var svc *service.Service

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go HTTP Service!!!")
}

func main() {
	var err error
	svc, err = service.NewService()
	if err != nil {
		fmt.Println("Error creating service:", err)
		return
	}
	http.HandleFunc("/", helloHandler)
	// GraphQL handler
	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &service.Resolver{svc}}))
	http.Handle("/graphql", srv)
	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)

}
