package lib

import (
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/webservice_go/models"
)

type ArticleResponse struct {
	Status   string `json:status`
	Source   string `json:source`
	SortBy   string `json:sortBy`
	Articles []models.Article `json:articles`
}

func ReadArticles(url string, body io.Reader) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Api-Key", "abd04ef3bdb24e8089494cbd128b0928")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}
	aResponse := new(ArticleResponse)
	json.NewDecoder(resp.Body).Decode(&aResponse)
	fmt.Println(aResponse.Articles)
	for _, ar := range aResponse.Articles{
		ar.Id = string(ar.Description)
		ar.Save()
	}
	defer resp.Body.Close()
}
