package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/djung460/mywiki/models"
)

func Articles(db models.DB) http.Handler {
	type (
		ret struct {
			Articles []string `json:"articles"`
		}
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articles, err := db.GetAllArticlesByUser("test")
		log.Print(articles)
		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}

		var jsonarticles []string
		// convert the articles slice into strings
		for _, d := range articles {
			log.Println(d.Title)
			jsonarticle, err := d.MarshalJSON()
			if err != nil {
				log.Print("ERROR PANIC")
				log.Print(err)
				jsonErr(w, http.StatusInternalServerError, err)
				return
			}
			jsonarticles = append(jsonarticles, string(jsonarticle))
		}
		if err := json.NewEncoder(w).Encode(ret{Articles: jsonarticles}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
