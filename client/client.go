package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/imzhongqi/okxos/errcode"
)

type Client struct {
	key        string
	secretKey  []byte
	passphrase string

	client *http.Client

	endpoint string
	headers  http.Header
}

func NewClient(key, secretKey, passphrase string, opts ...Option) *Client {
	options := newOptions(opts...)

	c := &Client{
		key:        key,
		secretKey:  []byte(secretKey),
		passphrase: passphrase,
		client:     options.client,
		endpoint:   options.endpoint,
		headers:    options.headers,
	}
	return c
}

func (c *Client) request(method string, path string, params map[string]string, body any, result any) error {
	req, err := c.newRequest(method, path, params, body)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.decode(resp.Body, result)
}

func (c *Client) decode(body io.Reader, result any) error {
	resp := &Response{
		Data: result,
	}
	if err := json.NewDecoder(body).Decode(resp); err != nil {
		return err
	}

	if code := resp.Code; code != 0 {
		return errcode.New(code.Int(), resp.Message)
	}

	return nil
}

func (c *Client) sign(ts string, method string, path string, body *bytes.Buffer) string {
	buf := &bytes.Buffer{}
	size := len(ts) + len(method) + len(path)
	if body != nil {
		size += body.Len()
	}
	buf.Grow(size)
	buf.WriteString(ts)
	buf.WriteString(method)
	buf.WriteString(path)
	if body != nil {
		buf.Write(body.Bytes())
	}
	return sign(c.secretKey, buf)
}

func (c *Client) newRequest(method string, path string, params map[string]string, body any) (*http.Request, error) {
	bodyBuf := bytes.NewBuffer(nil)
	if body != nil {
		if err := json.NewEncoder(bodyBuf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.endpoint+path, bodyBuf)
	if err != nil {
		return nil, err
	}
	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	ts := time.Now().In(time.UTC).Format("2006-01-02T15:04:05.000Z")
	signature := c.sign(ts, req.Method, req.URL.RequestURI(), bodyBuf)
	req.Header.Set("OK-ACCESS-KEY", c.key)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.passphrase)
	req.Header.Set("OK-ACCESS-TIMESTAMP", ts)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.headers {
		req.Header.Add(k, v[0])
	}

	return req, nil
}

func (c *Client) Get(ctx context.Context, path string, params map[string]string, result any) error {
	return c.request(http.MethodGet, path, params, nil, result)
}

func (c *Client) Post(ctx context.Context, path string, body any, result any) error {
	return c.request(http.MethodPost, path, nil, body, result)
}

func sign(key []byte, reader io.Reader) string {
	hasher := hmac.New(sha256.New, key)
	io.Copy(hasher, reader)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
