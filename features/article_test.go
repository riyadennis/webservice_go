package features

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
)

func iHaveAnArticleWithDetails(article *gherkin.DataTable) error {
	for _, ar := range article.Rows {
		for _, data := range ar.Cells {
			if data.Value == "" {
				return fmt.Errorf("not a valid table")
			} else {
				fmt.Println(data.Value)
			}
		}

	}
	return godog.ErrPending
}

func iQueryElastic() error {
	return godog.ErrPending
}

func iGetStatus(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have an article with details$`, iHaveAnArticleWithDetails)
	s.Step(`^I query elastic$`, iQueryElastic)
	s.Step(`^I get status "([^"]*)"$`, iGetStatus)
}
