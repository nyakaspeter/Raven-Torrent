//go:build arm

package httpserver

import "github.com/gorilla/mux"

func enableSwagger(router *mux.Router) {
	// empty function so swagger doesn't get built into the arm version
}
