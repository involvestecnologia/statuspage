package logs

import "github.com/involvestecnologia/statuspage/models"

//Log interface stabilishes a default contract for using a log repository
type Log interface {
	Error(args models.LogFields, logMessage string)
	Info(args models.LogFields, logMessage string)
	Warn(args models.LogFields, logMessage string)
}
