package checkchain

import (
	"github.com/gorilla/mux"
)

func NewHttpServer() *mux.Router {
	r := mux.NewRouter()
	//r.Use(middleware.IsAuthenticatedMiddleware)
	//r.Use(LoggingMiddleware)
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/collections/", CollectionsHandler)
	//r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	return r
}
