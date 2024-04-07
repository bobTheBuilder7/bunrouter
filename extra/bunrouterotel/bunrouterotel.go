package bunrouterotel

import (
	"net"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/bobTheBuilder7/bunrouter"
)

type config struct {
	clientIP bool
}

type Option func(c *config)

func WithClientIP() Option {
	return func(c *config) {
		c.clientIP = true
	}
}

func NewMiddleware(opts ...Option) bunrouter.MiddlewareFunc {
	c := &config{
		clientIP: false,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c.Middleware
}

func (c *config) Middleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		span := trace.SpanFromContext(req.Context())
		if !span.IsRecording() {
			return next(w, req)
		}

		params := req.Params()
		span.SetName(req.Method + " " + params.Route())

		paramSlice := params.Slice()
		attrs := make([]attribute.KeyValue, 0, 3+len(paramSlice))

		attrs = append(attrs, semconv.HTTPRouteKey.String(req.Route()))

		if req.URL != nil {
			attrs = append(attrs, semconv.HTTPTargetKey.String(req.URL.RequestURI()))
		} else {
			// This should never occur if the request was generated by the net/http
			// package. Fail gracefully, if it does though.
			attrs = append(attrs, semconv.HTTPTargetKey.String(req.RequestURI))
		}

		if c.clientIP {
			attrs = append(attrs, semconv.HTTPClientIPKey.String(remoteAddr(req.Request)))
		}

		for _, param := range paramSlice {
			attrs = append(attrs, attribute.String("http.route.param."+param.Key, param.Value))
		}

		span.SetAttributes(attrs...)

		if err := next(w, req); err != nil {
			span.SetStatus(codes.Error, err.Error())
			return err
		}

		return nil
	}
}

func remoteAddr(req *http.Request) string {
	forwarded := req.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	return host
}
