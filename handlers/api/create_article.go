package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/djung460/mywiki/models"
)

// CreateArticle handles creating an article
func CreateArticle(db models.DB) http.Handler {
	type jsonInput struct {
		Title string `json:"title"`
		//Username     string `json:"username"`
		//Category     string `json:"category"`
		Content string `json:"content"`
		//DateCreated  string `json:"date_created"`
		//DateModified string `json:"date_modified"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		log.Print(in)
		log.Println("Title: " + in.Title)
		log.Println("Content: " + in.Content)
		article := models.Article{
			Title:        in.Title,
			Username:     "test",
			Category:     "test cat",
			Content:      in.Content,
			DateCreated:  time.Now().Format(time.RFC3339),
			DateModified: time.Now().Format(time.RFC3339),
		}
		if _, err := db.UpsertArticle(article); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			log.Print(article)
			log.Println("JSON ERROR")
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
