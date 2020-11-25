package main

import (
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var (
	apiKey string
)

func getRequiredEnv(s string, required bool) string {
	v := os.Getenv(s)
	if len(v) == 0 && required {
		log.Fatalf("Missing required env: %s\n", s)
	}
	return v
}

func parseEnvs() {
	apiKey = getRequiredEnv("API_KEY", true)
}

type Resolver struct{}

func main() {
	parseEnvs()

	schema := graphql.MustParseSchema(rootSchema, &Resolver{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
