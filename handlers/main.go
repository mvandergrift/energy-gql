package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/jinzhu/gorm"

	"github.com/mvandergrift/energy-gql/graph"
	"github.com/mvandergrift/energy-gql/graph/generated"

	_ "github.com/lib/pq" // initialize postgres driver
)

var muxAdapter *gorillamux.GorillaMuxAdapter

func init() {
	r := mux.NewRouter()

	cn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv(("DB_DATABASE")))
	db, err := gorm.Open("postgres", cn)
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") == "TRUE" || os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.BlockGlobalUpdate(true)

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}})
	server := handler.NewDefaultServer(schema)

	// Catch-all path. Use AWS API Gateway to handle assigning specific paths to this handler @mvandergrift
	r.PathPrefix("/").Handler(server).Methods("POST")

	// Disabled playground to avoid obvious security issue. Also, reduce configuration complexity around
	// explicit pathing of /query endpoint. @mvandergrift
	// r.Handle("/energy/", playground.Handler("GraphQL playground", "/energy/query"))
	//r.PathPrefix("/").Handler(playground.Handler("GraphQL playground", "/energy/query")).Methods("GET")

	muxAdapter = gorillamux.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if os.Getenv("DEBUG") == "TRUE" || os.Getenv("DEBUG") == "true" {
		log.Printf("AWS Events: %+v", req)
	}

	rsp, err := muxAdapter.Proxy(req)

	if err != nil {
		log.Printf("[iex]\tError in handler: %v", err)
	}
	rsp.Headers = map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": "true"}
	return rsp, err
}

func main() {
	lambda.Start(Handler)
}
