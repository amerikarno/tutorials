package repository_test

import (
	"fmt"
	"testing"
	"wnexam/repository"
)

func TestGetJSONUrl(t *testing.T) {
	t.Parallel()
	t.Run("get JSON error", func(t *testing.T){
		url := "https://static.wongnai.com/devinterview/covid-cases.json"

		repo := repository.NewDataRepository(url)
		gju, err := repo.GetJSONUrl()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(gju))
	})
}