package handlers

import "net/http"

// Articles functio handles rendering articles
func Articles(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "articles", struct{}{}, "layout")
	})
}
