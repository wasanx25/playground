package main

import (
	"bytes"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.travis-ci.org",
		Path:   "repo/wasanx25%2Fplayground/requests",
	}

	body := `{"request":{"branch":"master"}}`

	req, err := http.NewRequest(
		"POST",
		u.String(),
		bytes.NewBuffer([]byte(body)),
	)

	if err != nil {
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
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	defer res.Body.Close()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success Travis Build!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
