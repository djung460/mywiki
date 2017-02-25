package api

import (
	"encoding/json"
	"log"
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/djung460/mywiki/models"
)

// CreateArticle handles creating an article
func CreateArticle(db models.DB) http.Handler {
	type jsonInput struct {
		Title string `json:"title"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		log.Println(in.Title)
		article := models.Article{Title: in.Title}
		if _, err := db.Upsert(bson.NewObjectId().String(), article); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			log.Println()
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
