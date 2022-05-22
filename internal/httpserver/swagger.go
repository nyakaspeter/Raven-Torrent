//go:build !arm

package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/nyakaspeter/raven-torrent/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func enableSwagger(router *mux.Router) {
	router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
