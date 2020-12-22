package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/arrivals/*", spaHandler)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "ui/public"))
	FileServer(r, "/", filesDir)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(rootSchema, &Resolver{}, opts...)
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", r))
}

// Catch-all function for spa routes
func spaHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content Type", "text/html")
	http.ServeFile(rw, req, "ui/public/index.html")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
