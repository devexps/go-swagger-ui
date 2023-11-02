package swaggerUI

import (
	"github.com/devexps/go-swagger-ui/internal/swagger"
	"net/http"
)

// Handler handles swagger UI request.
type Handler = swagger.Handler

// New creates HTTP handler for Swagger UI.
func New(title, swaggerJSONPath string, basePath string) http.Handler {
	return newHandler(title, swaggerJSONPath, basePath)
}

// NewWithOption creates configurable handler constructor.
func NewWithOption(handlerOpts ...HandlerOption) http.Handler {
	opts := swagger.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}
	return newHandlerWithConfig(opts)
}

// NewHandler creates HTTP handler for Swagger UI.
func newHandler(title, swaggerJSONPath string, basePath string) *Handler {
	return newHandlerWithConfig(&swagger.Config{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	})
}

// newHandlerWithConfig creates HTTP handler for Swagger UI.
func newHandlerWithConfig(config *swagger.Config) *Handler {
	return swagger.NewHandlerWithConfig(config, assetsBase, faviconBase, staticServer)
}
