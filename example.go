package main

import (
	"fmt"

	"github.com/arnocho/http-buddy/httpbuddy"
	"github.com/arnocho/http-buddy/httpbuddy/gomime"
)

var (
	client = getHttpClient()
)

func getHttpClient() httpbuddy.Client {
	return httpbuddy.NewBuilder().
		SetUserAgent(gomime.UserAgent).
		Build()
}

type Endpoint struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func Get() {
	response, err := client.Get("https://jobs.louisvuitton.com/fra-fr/careers/jobs", true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status Code:", response.StatusCode())
	fmt.Println("Status:", response.Status())
	fmt.Println("Response Body:", response.String())

	//var endpoint Endpoint

	//if err := response.UnmarshalJson(&endpoint); err != nil {
	//	panic(err)
	//}

	//fmt.Println("Repository url:", endpoint.RepositoryURL)
}

func main() {
	//req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("User-Agent", userAgent)
	//cooky := &http.Cookie{Name: "bm_sz", Value: "just_to_bypass_bot_detection"}
	//req.AddCookie(cooky)
	Get()
}
