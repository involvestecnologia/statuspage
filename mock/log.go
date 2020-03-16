package mock

import (
	"log"

	"github.com/involvestecnologia/statuspage/logs"
	"github.com/involvestecnologia/statuspage/models"
)

type logRepositoryMock struct {
}

//NewLogRepositoryMock creates a new mock for the log repository
func NewLogRepositoryMock() logs.Log {
	return &logRepositoryMock{}
}

func (mock *logRepositoryMock) Info(obj models.LogFields, logMessage string) {
	log.Println("[INFO]", obj, logMessage)
}

func (mock *logRepositoryMock) Error(obj models.LogFields, logMessage string) {
	log.Println("[ERROR]", obj, logMessage)
}

func (mock *logRepositoryMock) Warn(obj models.LogFields, logMessage string) {
	log.Println("[WARN]", obj, logMessage)
}

func (mock *logRepositoryMock) Debug(obj models.LogFields, logMessage string) {
	log.Println("[DEBUG]", obj, logMessage)
}
