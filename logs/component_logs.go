package logs

import "github.com/involvestecnologia/statuspage/models"

func (l *logRepository) Info(obj models.Component, logMessage string) {
	l.info(obj, logMessage)
}

func (l *logRepository) Error(obj models.Component, logMessage string) {
	l.error(obj, logMessage)
}

func (l *logRepository) Warn(obj models.Component, logMessage string) {
	l.warn(obj, logMessage)
}
