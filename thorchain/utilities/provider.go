package utilities

import (
	"context"
	"net/http"
)

type Provider interface {
	Do(ctx context.Context, request *http.Request)
}
