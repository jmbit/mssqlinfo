package httpserver

import "net/http"

func registerRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/v0/server", serverHandler)
	mux.HandleFunc("/v0/dbls", dblsHandler)

	return mux
}
