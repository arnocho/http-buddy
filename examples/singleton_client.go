package examples

import (
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
