package rest

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// OnPanic, if defined, will handle
var OnPanic AppHandler

// Namespace enables setting custom namespace
var Namespace func(ctx context.Context) (string, error) = func(ctx context.Context) (string, error) {
	return "", nil
}

// OnError is a custom error handler definition. By default it simple returns an http.Error()
// with the given code and status text.
var OnError func(ctx context.Context, code int, err error) = func(ctx context.Context, code int, err error) {
	Errorf(ctx, "error: %v", err)
	http.Error(ResponseWriter(ctx), err.Error(), code)
}

// GetErrorCode allows customizing of the http.StatusCode for any given error
var GetErrorCode func(ctx context.Context, err error) int = func(ctx context.Context, err error) (code int) {
	switch {
	case err == datastore.ErrNoSuchEntity:
		code = http.StatusNotFound
	case appengine.IsOverQuota(err):
		code = http.StatusTooManyRequests
	case appengine.IsTimeoutError(err):
		code = http.StatusGatewayTimeout
	default:
		code = http.StatusInternalServerError
	}
	return
}
