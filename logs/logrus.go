package logs

import (
	"github.com/involvestecnologia/statuspage/component"
	"github.com/sirupsen/logrus"
)

type logRepository struct {
	logger *logrus.Logger
}

//NewLogRepository creates a new logger instance
func NewLogRepository(logger *logrus.Logger) component.Log {
	return &logRepository{
		logger: logger,
	}
}

//Info logs something to stdout as INFO
func (l *logRepository) info(arg interface{}, logMessage string) {
	l.logger.WithFields(logrus.Fields{
		"obj": arg,
	}).Info(logMessage)
}

//Warn logs somehting to stdout as Warn
func (l *logRepository) warn(arg interface{}, logMessage string) {
	l.logger.WithFields(logrus.Fields{
		"obj": arg,
	}).Warn(logMessage)
}

//Error logs something to stdout as Error
func (l *logRepository) error(arg interface{}, logMessage string) {
	l.logger.WithFields(logrus.Fields{
		"obj": arg,
	}).Error(logMessage)
}
