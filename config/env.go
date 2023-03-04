package config

import (
	"fmt"
	"os"
)

var ENV env

type env struct {
	NotionKey    string
	NotionApiUrl string
	DatabaseId   string
}

func LoadEnv() env {
	ENV.NotionKey = os.Getenv("NOTION_API_KEY")
	ENV.DatabaseId = os.Getenv("DATABASE_ID")
	ENV.NotionApiUrl = os.Getenv("URL") + os.Getenv("DATABASE_ID") + "/query"

	fmt.Println(ENV.NotionApiUrl)

	return ENV
}
