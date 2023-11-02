package swaggerUI

import (
	"fmt"
	"github.com/devexps/go-swagger-ui/internal/swagger"
	"net/http"
	"path"
	"strings"
)

var _openJsonFileHandler = &openJsonFileHandler{}

type httpServerInterface interface {
	HandlePrefix(prefix string, h http.Handler)
	Handle(path string, h http.Handler)
	HandleFunc(path string, h http.HandlerFunc)
}

// RegisterSwaggerUIServer creates handler and register router
func RegisterSwaggerUIServer[T httpServerInterface](srv T, title, swaggerJSONPath string, basePath string) {
	swaggerHandler := newHandler(title, swaggerJSONPath, basePath)
	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

// RegisterSwaggerUIServerWithOption creates handler and register router base on the handler options.
func RegisterSwaggerUIServerWithOption[T httpServerInterface](srv T, handlerOpts ...HandlerOption) {
	opts := swagger.NewConfig()

	for _, o := range handlerOpts {
		o(opts)
	}
	if opts.LocalOpenApiFile != "" {
		registerOpenApiLocalFileRouter(srv, opts)
	} else if len(opts.OpenApiData) != 0 {
		registerOpenApiMemoryDataRouter(srv, opts)
	}
	swaggerHandler := newHandlerWithConfig(opts)

	srv.HandlePrefix(swaggerHandler.BasePath, swaggerHandler)
}

// registerOpenApiLocalFileRouter .
func registerOpenApiLocalFileRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	err := _openJsonFileHandler.LoadFile(cfg.LocalOpenApiFile)
	if err == nil {
		pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi" + path.Ext(cfg.LocalOpenApiFile)
		cfg.SwaggerJSON = pattern
		srv.Handle(pattern, _openJsonFileHandler)
	} else {
		fmt.Println("load openapi file failed: ", err)
	}
}

// registerOpenApiMemoryDataRouter .
func registerOpenApiMemoryDataRouter[T httpServerInterface](srv T, cfg *swagger.Config) {
	_openJsonFileHandler.Content = cfg.OpenApiData
	pattern := strings.TrimRight(cfg.BasePath, "/") + "/openapi." + cfg.OpenApiDataType
	cfg.SwaggerJSON = pattern
	srv.Handle(pattern, _openJsonFileHandler)
	cfg.OpenApiData = nil
}
