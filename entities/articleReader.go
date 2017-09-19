package entities

import (
	"io"
	"net/http"
	"encoding/json"
	"strings"
)

type ArticleResponse struct {
	Status   string `json:status`
	Source   string `json:source`
	SortBy   string `json:sortBy`
	Articles []Article `json:articles`
}

func ReadArticles(url string, body io.Reader, key string) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Api-Key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}
	aResponse := new(ArticleResponse)
	json.NewDecoder(resp.Body).Decode(&aResponse)

	for _, ar := range aResponse.Articles {
		ar.Id = generateArticleId(ar.Title)
		ar.Save()
	}
	defer resp.Body.Close()
}
func generateArticleId(desc string) (string) {
	description := strings.Split(desc, " ")
	des := description[0]
	i :=1
	for len(des) < 10 {
		des += "_"+description[i]
		i++
	}
	return strings.ToLower(des)
}
