package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/djung460/mywiki/models"
)

// CreateArticle handles creating an article
func CreateArticle(db models.DB) http.Handler {
	type jsonInput struct {
		Title        string `json:"title"`
		Username     string `json:"username"`
		Category     string `json:"category"`
		Content      string `json:"content"`
		DateCreated  string `json:"date_created"`
		DateModified string `json:"date_modified"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		log.Println(in.Title)
		article := models.Article{
			Title:        in.Title,
			Username:     in.Username,
			Category:     in.Category,
			Content:      in.Content,
			DateCreated:  in.DateCreated,
			DateModified: in.DateModified,
		}
		if _, err := db.UpsertArticle(article); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			log.Println()
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
