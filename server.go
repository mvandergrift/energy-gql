package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mvandergrift/energy-gql/graph"
	"github.com/mvandergrift/energy-gql/graph/generated"
	"github.com/rs/cors"

	_ "github.com/joho/godotenv/autoload" // autoload .env configuration
	_ "github.com/lib/pq"                 // initialize postgres driver
)

func main() {
	cn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv(("DB_DATABASE")))
	debugMode := flag.Bool("debug", false, "Enable debug logging to console.")
	sqlDebug := flag.Bool("sql", false, "Enable SQL logging to console.")
	flag.Parse()

	db, err := gorm.Open("postgres", cn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.BlockGlobalUpdate(true)
	db.LogMode(*sqlDebug)

	// Setup gqlgen (GraphQL)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	// Setup router
	r := mux.NewRouter()
	r.Use(cors.AllowAll().Handler)
	r.Use(handlers.CompressHandler)
	r.Handle("/query", srv)
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))

	if *debugMode {
		log.Println("[iex]\tStarting iex-gql")
		log.Printf("[iex]\tconnect to http://localhost:%s/ for GraphQL playground", os.Getenv("PORT"))
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.LoggingHandler(os.Stdout, r)))
	}

	if *sqlDebug {
		log.Println("[iex]\tSQL debugging enabled")
	}

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

}
