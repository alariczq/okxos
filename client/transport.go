package client

import "context"

type Transport interface {
	Get(ctx context.Context, path string, params map[string]string, result any) error
	Post(ctx context.Context, path string, body any, result any) error
}
