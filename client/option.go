package client

import (
	"net/http"
)

type Options struct {
	endpoint string
	headers  http.Header
	client   *http.Client
}

type Option interface {
	apply(o *Options)
}

type optionFunc func(o *Options)

func (f optionFunc) apply(o *Options) {
	f(o)
}

func WithClient(client *http.Client) Option {
	return optionFunc(func(o *Options) {
		o.client = client
	})
}

func WithEndpoint(endpoint string) Option {
	return optionFunc(func(o *Options) {
		o.endpoint = endpoint
	})
}

func WithHeaders(headers http.Header) Option {
	return optionFunc(func(o *Options) {
		o.headers = headers
	})
}

func WithHeader(key, value string) Option {
	return optionFunc(func(o *Options) {
		if o.headers == nil {
			o.headers = make(http.Header)
		}
		o.headers.Set(key, value)
	})
}

func WithProjectID(projectID string) Option {
	return WithHeader("OK-ACCESS-PROJECT", projectID)
}

func newOptions(opts ...Option) Options {
	o := Options{
		endpoint: "https://www.okx.com",
		headers:  make(http.Header),
		client:   http.DefaultClient,
	}
	for _, opt := range opts {
		opt.apply(&o)
	}
	return o
}
