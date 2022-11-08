package handlers;

import (
	"encoding/json"
	"net/http"

	"github.com/funkynunu/go_rest_api/v2/pkg/mocks"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Articles)
}