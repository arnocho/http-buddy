package httpbuddy

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, bypassBotFilter bool, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error)
	Delete(url string, bypassBotFilter bool, headers ...http.Header) (*Response, error)

	CheckForString(url string, itemToCheck string, bypassBotFilter bool, headers ...http.Header) (bool, error)
}

func (c *httpClient) Get(url string, bypassBotFilter bool, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, bypassBotFilter, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, bypassBotFilter, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, bypassBotFilter, getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, body interface{}, bypassBotFilter bool, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, bypassBotFilter, getHeaders(headers...), body)
}

func (c *httpClient) Delete(url string, bypassBotFilter bool, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, bypassBotFilter, getHeaders(headers...), nil)
}

func (c *httpClient) CheckForString(url string, itemToCheck string, bypassBotFilter bool, headers ...http.Header) (bool, error) {
	return c.checkForString(url, itemToCheck, bypassBotFilter, headers...)
}
