package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/funkynunu/go_rest_api/pkg/mocks"
	"github.com/funkynunu/go_rest_api/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request from body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var updateArticle models.Article
	json.Unmarshal(body, &updateArticle)

	for index, article := range mocks.Articles {
		if article.Id == id {
			article.Title = updateArticle.Title
			article.Desc = updateArticle.Desc
			article.Content = updateArticle.Content

			mocks.Articles[index] = article

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
