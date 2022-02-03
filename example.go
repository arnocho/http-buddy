package main

import (
	"fmt"
	"time"

	"github.com/arnocho/http-buddy/httpbuddy"
)

var (
	client = getHttpClient()
)

func getHttpClient() httpbuddy.Client {
	return httpbuddy.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseHeaderTimeout(3 * time.Second).
		Build()
}

type Endpoint struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func Get() {
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status Code:", response.StatusCode())
	fmt.Println("Status:", response.Status())
	fmt.Println("Response Body:", response.String())

	var endpoint Endpoint

	if err := response.UnmarshalJson(&endpoint); err != nil {
		panic(err)
	}

	fmt.Println("Repository url:", endpoint.RepositoryURL)
}

func main() {
	Get()
}
