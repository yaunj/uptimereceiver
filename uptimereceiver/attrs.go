package uptimereceiver

import (
	"os"
	"runtime"

	"go.uber.org/zap"
)

type attributes struct {
	hostname string
	os       string
	arch     string
}

type attributeReader struct {
	logger            *zap.Logger
	defaultAttributes *attributes
}

func newAttributeReader(logger *zap.Logger) *attributeReader {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Warn("could not determine hostname", zap.Error(err))
		hostname = "unknown"
	}

	return &attributeReader{
		logger: logger,
		defaultAttributes: &attributes{
			hostname: hostname,
			os:       runtime.GOOS,
			arch:     runtime.GOARCH,
		},
	}
}

// getAttributes serves as a layer of indirection in case we need to look up some attributes more often
func (a *attributeReader) getAttributes() *attributes {
	return a.defaultAttributes
}
