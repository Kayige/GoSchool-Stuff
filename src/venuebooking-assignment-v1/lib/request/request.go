package request

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	// headerClientSecret is a Header key to set client secret
	headerClientSecret = "X-Client-Secret"
	// ContextKeyRequestID is the key holding the request ID
	contextKeyRequestID = "requestID"
)

// assignRequestID will attach a brand new request ID to a http request
func AssignRequestID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, contextKeyRequestID, reqID.String())
}

// GetRequestID will get reqID from a http request and return it as a string
func GetRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if ret, ok := reqID.(string); ok {
		return ret
	}
	return ""
}

// GetRequestClientSecret returns the value of the client secret sent in the header of the request
func GetRequestClientSecret(r *http.Request) string {
	return r.Header.Get(headerClientSecret)
}
