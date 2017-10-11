package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/webservice_go/lib/mock_lib"
	"github.com/golang/mock/gomock"
)

func TestCreateRandomDigits(t *testing.T) {
	digit := CreateRandomDigits(100, 200)
	if digit > 200 {
		assert.Fail(t, "Invalid random number generated")
	}
}

func TestGenerateArticleId(t *testing.T) {
	articleId, err := GenerateArticleId("Test error description")
	assert.NoError(t, err)
	assert.NotEmpty(t, articleId)
}
func TestSaveArticles(t *testing.T) {
	mockReader := mock_lib.NewMockReader(gomock.NewController(t))
	recorder := mockReader.EXPECT()
	recorder.Read()
}