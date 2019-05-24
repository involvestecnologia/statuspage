package logs

import (
	"github.com/involvestecnologia/statuspage/models"
	"github.com/sirupsen/logrus"
)

type logRepository struct {
	logger *logrus.Logger
}

//NewLogRepository creates a new logger instance
func NewLogRepository(logger *logrus.Logger) Log {
	return &logRepository{
		logger: logger,
	}
}

//Info logs something to stdout as INFO
func (l *logRepository) Info(args models.LogFields, logMessage string) {
	l.logger.WithFields(logrus.Fields(args)).Info(logMessage)
}

//Warn logs somehting to stdout as Warn
func (l *logRepository) Warn(args models.LogFields, logMessage string) {
	l.logger.WithFields(logrus.Fields(args)).Warn(logMessage)
}

//Error logs something to stdout as Error
func (l *logRepository) Error(args models.LogFields, logMessage string) {
	l.logger.WithFields(logrus.Fields(args)).Error(logMessage)
}
