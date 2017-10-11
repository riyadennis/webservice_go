package config

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	path := fmt.Sprintf("../%s", "config.yml")
	myConfig, err := GetConfig(path)
	assert.NoError(t, err)
	assert.NotEmpty(t, myConfig.Article.Url)
}
func TestGetConfigInvalidFile(t *testing.T){
	path := fmt.Sprintf("../%s", "invalidConfig.yaml")
	_, err := GetConfig(path)
	assert.NotEmpty(t, err)
}