package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	//"github.com/kr/pretty"
	"github.com/rusakov/doto/gql"
	//"github.com/rusakov/doto/gql/db"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", handler.Playground("DOTO way", "/query"))
	//http.Handle("/query", handler.GraphQL(starwars.NewExecutableSchema(starwars.NewResolver()),
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.NewResolver()),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetResolverContext(ctx)
			fmt.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
			return res, err
		}),
	))
	//if err := db.Run(); err != nil {
	//	pretty.Print(err)
	//os.Exit(1)
	//}
	//log.Fatal(db.Run())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
