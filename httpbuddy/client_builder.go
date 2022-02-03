package httpbuddy

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers http.Header

	maxIdleConnections    int
	connectionTimeout     time.Duration
	responseHeaderTimeout time.Duration
	disableTimeouts       bool
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetMaxIdleConnections(m int) ClientBuilder
	SetConnectionTimeout(t time.Duration) ClientBuilder
	SetResponseHeaderTimeout(t time.Duration) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	clientBuilder := &clientBuilder{}
	return clientBuilder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(m int) ClientBuilder {
	c.maxIdleConnections = m
	return c
}

func (c *clientBuilder) SetConnectionTimeout(t time.Duration) ClientBuilder {
	c.connectionTimeout = t
	return c
}

func (c *clientBuilder) SetResponseHeaderTimeout(t time.Duration) ClientBuilder {
	c.responseHeaderTimeout = t
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}
