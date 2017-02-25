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
		keys, err := db.GetAllKeys()
		log.Print(keys)
		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		// Note: this is a potentially large scale operation.
		// several improvements could be made:
		// - paginate the results, to provide an upper bound on amount of work in a single request
		// - send only the keys down to the browser, and have the browser do a GET on only the keys it needs
		articles := []string{}
		for _, key := range keys {
			article := new(models.Article)
			log.Print("Hello?")
			db.Get(key, article)
			articles = append(articles, article.Title)
		}
		if err := json.NewEncoder(w).Encode(ret{Articles: articles}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
