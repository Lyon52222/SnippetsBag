package log

import (
	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.AppConfig) *logrus.Entry {
	var log *logrus.Logger
	log = logrus.New()
	return log.WithFields(logrus.Fields{
		"verson": config.Version,
	})
}
