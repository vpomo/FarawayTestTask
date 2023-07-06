package checkchain

import (
	"github.com/gorilla/mux"
)

func NewHttpServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/collections", CollectionsHandler).Methods("GET")
	r.HandleFunc("/tokenminted/{collectionAddress}", TokenMintedHandler).Methods("GET")
	return r
}
