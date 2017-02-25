package models

import "encoding/json"

// Article defines the attributes of the article
type (
	Article struct {
		Title string `bson:"title"`
		Body  string `bson:"body"`
	}
)

// MarshalJSON converts go struct to json
func (a Article) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"title": a.Title, "body": a.Body})
}
