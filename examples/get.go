package examples

import "fmt"

type Endpoint struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func GetEndpoint() (*Endpoint, error) {
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Status Code:", response.StatusCode())
	fmt.Println("Status:", response.Status())
	fmt.Println("Response Body:", response.String())
	var endpoint Endpoint

	if err := response.UnmarshalJson(&endpoint); err != nil {
		return nil, err
	}

	return &endpoint, nil
}
