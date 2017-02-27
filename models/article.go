package models

import "encoding/json"

type User struct {
	Username string
	Password string
	Email    string
}

// Article defines the attributes of the article
type Article struct {
	Username     string
	Title        string
	Content      string
	DateCreated  string
	DateModified string
	Category     string
}

type Category struct {
	Name string
}

// MarshalJSON converts go struct to json
func (a Article) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"username":      a.Username,
		"title":         a.Title,
		"content":       a.Content,
		"date_created":  a.DateCreated,
		"date_modified": a.DateModified,
		"category":      a.Category,
	})
}
