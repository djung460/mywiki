package handlers

import "net/http"

// Index handles rendering index
func Index(ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ren.Render(w, http.StatusOK, "index", struct{}{}, "layout")
	})
}
