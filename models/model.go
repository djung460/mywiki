package models

import "encoding/json"

// Model mapping between json and go values
type Model interface {
	json.Marshaler
}
