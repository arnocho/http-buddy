package httpbuddy

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/arnocho/http-buddy/httpbuddy/gomime"
)

const (
	defaultMaxIdleConnections    = 5
	defaultResponseHeaderTimeout = 5 * time.Second
	defaultConnectionTimeout     = 1 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)

	case gomime.ContentTypeXml:
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}

func (c *httpClient) checkForString(url string, itemToCheck string, headers ...http.Header) (bool, error) {
	response, err := c.do(http.MethodGet, url, getHeaders(headers...), nil)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(response.body), itemToCheck), nil
}

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*Response, error) {
	allHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(allHeaders.Get(gomime.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create the request")
	}

	request.Header = allHeaders

	//To keep the interface clean, i put this all the time
	cookie := &http.Cookie{Name: "bm_sz", Value: "keep_it_clean"}
	request.AddCookie(cookie)

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       responseBody,
	}

	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost: c.getMaxIdleConnections(),
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}
