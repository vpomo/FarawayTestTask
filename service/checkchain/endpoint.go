package checkchain

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CollectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//w.Write()
	log.Info("CollectionsHandler")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	fmt.Println("ArticleHandler")
	fmt.Println("category", vars["category"])
	fmt.Println("id", vars["id"])
}
