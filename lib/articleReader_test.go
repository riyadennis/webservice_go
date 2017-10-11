package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
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