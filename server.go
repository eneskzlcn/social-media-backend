package main

import (
	"fmt"
	"github.com/eneskzlcn/hackernews/config"
	"github.com/eneskzlcn/hackernews/mongo"
)

const defaultPort = "8080"

func main() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = defaultPort
	//}
	//
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)
	//
	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	mongoDB := mongo.NewDB(config.DB{
		Host: "localhost",
		Port: 27017,
		ConnectionTimeoutInSeconds: 10,
	})
	if err := mongoDB.Connect() ; err != nil {
		fmt.Println(err.Error())
	}

}
