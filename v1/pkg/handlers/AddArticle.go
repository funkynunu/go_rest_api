package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/funkynunu/go_rest_api/v1/pkg/mocks"
	"github.com/funkynunu/go_rest_api/v1/pkg/models"

	"github.com/google/uuid"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	// Read request from body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var article models.Article
	json.Unmarshal(body, &article)

	article.Id = (uuid.New()).String()
	mocks.Articles = append(mocks.Articles, article)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
