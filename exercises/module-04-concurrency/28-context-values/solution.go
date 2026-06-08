package ctxvalues

import "context"

type contextKey string

const RequestIDKey contextKey = "request_id"

// WithRequestID adds a request ID to the context
// FIX ME: Use context.WithValue
func WithRequestID(ctx context.Context, id string) context.Context {
	return ctx // FIX
}

// GetRequestID extracts the request ID from context
// FIX ME: Use ctx.Value
func GetRequestID(ctx context.Context) string {
	return "" // FIX
}
