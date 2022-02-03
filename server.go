package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cryanbrow/eve-graphql-go/graph"
	"github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

var (
	//go:embed voyager
	res   embed.FS
	pages = map[string]string{
		"/voyager": "voyager/index.html",
	}
)

func main() {
	setupDependencies()
	port := os.Getenv("PORT")
	if port == "" {
		port = configuration.AppConfig.Server.Port
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.HandleFunc("/voyager", func(w http.ResponseWriter, r *http.Request) {
		page, ok := pages[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		template, err := template.ParseFS(res, page)
		if err != nil {
			log.Printf("page %s not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"userAgent": r.UserAgent(),
		}
		if err := template.Execute(w, data); err != nil {
			return
		}
	})
	http.FileServer(http.FS(res))

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func setupDependencies() {
	configuration.LoadConfiguration()
	helpers.SetupRestHelper()
	caching.ConfigureRedisClient()
	universe.SetupUniverseRest()
}
