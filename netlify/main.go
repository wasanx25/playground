package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := `{"request":{"branch":"master"}}`

	req, err := http.NewRequest(
		"POST",
		"https://api.travis-ci.org/repo/wasanx25%2Fplayground/requests",
		bytes.NewBuffer([]byte(body)),
	)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Autorization", "token "+os.Getenv("TRAVIS_TOKEN"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	defer res.Body.Close()
	fmt.Println("Travis Calling Status: " + res.Status + "\nTravis Calling Message: " + string(resBody))

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success Travis Build!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
