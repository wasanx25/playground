package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	travis "github.com/shuheiktgw/go-travis"
)

func handler(request events.APIGatewayProxyRequest) {
	client := travis.NewClient(travis.ApiOrgUrl, os.Getenv("TRAVIS_TOKEN"))
	branch := os.Getenv("BRANCH")
	if branch == "" {
		branch = "master"
	}
	req := &travis.RequestBody{
		Branch: branch,
		Config: map[string]map[string]string{
			"env": map[string]string{
				"CYPRESS_baseUrl":      os.Getenv("URL"),
				"CYPRESS_DEV_TEST_URL": os.Getenv("URL"),
			},
		},
	}
	_, res, err := client.Requests.CreateByRepoSlug(context.Background(), "wasanx25/playground", req)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func main() {
	lambda.Start(handler)
}
