package lib

import (
	"io"
	"net/http"
	"encoding/json"
	"strings"
	"github.com/webservice_go/entities"
	"time"
	"math/rand"
	"strconv"
)

type Reader interface {
	Read()
}
type ArticleResponse struct {
	Status   string `json:status`
	Source   string `json:source`
	SortBy   string `json:sortBy`
	Articles []entities.Article `json:articles`
}
type ArticleReader struct {
	Url  string
	Body io.Reader
	Key  string
}

func (r ArticleReader) Read() (error) {
	req, err := http.NewRequest("GET", r.Url, r.Body)
	if err != nil {
		return err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Api-Key", r.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return err
	}
	aResponse := new(ArticleResponse)
	json.NewDecoder(resp.Body).Decode(&aResponse)
	SaveArticles(aResponse.Articles)

	defer resp.Body.Close()
	return nil
}

func SaveArticles(articles []entities.Article) {
	for _, ar := range articles {
		ar.Id = generateArticleId(ar.Title)
		ar.Save()
	}
}
func generateArticleId(desc string) (string) {
	description := strings.Split(desc, " ")
	des := description[0]
	i := 1
	for len(des) < 10 {
		des += "_" + description[i]
		i++
	}
	randomNum := CreateRandomDigits(10000, 99999)
	return strings.ToLower(des) + "_" + strconv.Itoa(randomNum)
}

func CreateRandomDigits(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
