// Copyright (c) 2024-NOW imzhongqi <imzhongqi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
