package models

import "encoding/json"

// Model mapping between json and go values
type Model interface {
	// Any model just has to impement this
	json.Marshaler
}
