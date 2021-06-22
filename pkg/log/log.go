package log

import (
	"io/ioutil"

	"github.com/Lyon52222/snippetsbag/pkg/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.AppConfig) *logrus.Entry {
	var log *logrus.Logger
	log = newProductionLogger()
	return log.WithFields(logrus.Fields{
		"version": config.Version,
	})
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	log.SetLevel(logrus.ErrorLevel)
	return log
}
