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
	"github.com/cryanbrow/eve-graphql-go/graph/generated"
)

var (
	//go:embed voyager
	res   embed.FS
	pages = map[string]string{
		"/voyager": "voyager/index.html",
	}
)

const defaultPort = "8080"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
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
		tpl, err := template.ParseFS(res, page)
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
		if err := tpl.Execute(w, data); err != nil {
			return
		}
	})
	http.FileServer(http.FS(res))

	log.Infoln("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
