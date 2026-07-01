package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

func TelemetryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.Tracer("transactmesh-wallet")

		ctx, span := tracer.Start(c.Request.Context(), c.Request.URL.Path)
		defer span.End()

		traceID := span.SpanContext().TraceID().String()
		c.Writer.Header().Set("x-trace-id", traceID)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
