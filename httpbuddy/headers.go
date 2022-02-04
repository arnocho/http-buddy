package httpbuddy

import (
	"net/http"

	"github.com/arnocho/http-buddy/httpbuddy/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	//Add common headers
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Add request headers (overwrites common)
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Set user agent
	if c.builder.userAgent != "" {
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}

	return result
}
