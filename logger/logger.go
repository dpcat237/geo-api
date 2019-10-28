package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Logger wraps logrus.Logger
type Logger struct {
	*logrus.Logger
}

// New creates a new preconfigured logrus.Logger
func New() Logger {
	logger := logrus.New()
	logger.Formatter = &prefixed.TextFormatter{
		FullTimestamp:    true,
		ForceColors:      true,
		DisableSorting:   false,
		QuoteEmptyFields: true,
		SpacePadding:     32,
	}

	return Logger{
		Logger: logger,
	}
}
