package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/djung460/mywiki/handlers"
	"github.com/djung460/mywiki/handlers/api"
	"github.com/djung460/mywiki/models"
	"github.com/djung460/mywiki/util"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := util.Config{}

	if err := envconfig.Process("mywiki", &config); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	env, err := config.Env()
	if err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	dev := env == util.EnvDev
	log.Println(dev)
	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)

	var db models.DB
	if dev {
		d, err := models.Init()
		if err != nil {
			log.Fatalf("error connecting to mongo [%s]", err)
			os.Exit(1)
		}
		db = d
	}

	// Gorilla mux
	r := mux.NewRouter()
	r.Handle("/", handlers.Index(renderer)).Methods("GET")
	r.Handle("/articles", handlers.Articles(renderer)).Methods("GET")

	// APIs
	r.Handle("/api/articles", api.Articles(db)).Methods("GET")
	//r.Handle("/api/article", api.CreateArticle(db)).Methods("PUT")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, http.StatusNotFound, "not_found", map[string]string{
			"url": r.URL.String(),
		}, "layout")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", config.Port)
	n.Run(hostStr)
}
