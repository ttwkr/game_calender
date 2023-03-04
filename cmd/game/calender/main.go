package main

import (
	"bytes"
	"context"
	"game_calender/config"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayV2HTTPResponse

func requestHttp(url string, token string, body []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", "2022-06-28")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(resBody))
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (bool, error) {
	env := config.LoadEnv()
	url := env.NotionApiUrl
	token := env.NotionKey
	body := []byte(`{}`)
	requestHttp(url, token, body)
	return true, nil
}

func main() {
	lambda.Start(Handler)
}
