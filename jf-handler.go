package swaggerUI

import (
	"io"
	"net/http"
	"os"
)

type openJsonFileHandler struct {
	Content []byte
}

// ServeHTTP writes out the content.
func (h *openJsonFileHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write(h.Content)
}

// LoadFile loads json content form the input filePath.
func (h *openJsonFileHandler) LoadFile(filePath string) error {
	content, err := h.loadOpenApiFile(filePath)
	if err != nil {
		return err
	}
	h.Content = content
	return nil
}

// loadOpenApiFile reads a file content.
func (h *openJsonFileHandler) loadOpenApiFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	return content, err
}
