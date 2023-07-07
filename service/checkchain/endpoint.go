package checkchain

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service OK!!!"))
}

func CollectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, jsonError := json.Marshal(GetCollectionsCreated())
	if jsonError != nil {
		log.Fatal("Unable to encode JSON")
	}
	w.Write(jsonResponse)
}

func TokenMintedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	tokenMinted := GetTokenMinted(vars["collectionAddress"])
	if len(tokenMinted) == 0 {
		w.Write([]byte("[]"))
		return
	}

	jsonResponse, jsonError := json.Marshal(tokenMinted)
	if jsonError != nil {
		log.Fatal("Unable to encode JSON")
	}
	w.Write(jsonResponse)
}
