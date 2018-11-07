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
	req := &travis.RequestBody{
		Branch: "master",
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
