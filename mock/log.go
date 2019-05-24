package mock

import (
	"github.com/involvestecnologia/statuspage/logs"
	"github.com/involvestecnologia/statuspage/models"
)

type componentLogRepositoryMock struct {
}

//NewComponentLogRepositoryMock creates a new mock for the log repository
func NewComponentLogRepositoryMock() logs.Log {
	return &componentLogRepositoryMock{}
}

func (mock *componentLogRepositoryMock) Info(obj models.LogFields, logMessage string) {
	return
}

func (mock *componentLogRepositoryMock) Error(obj models.LogFields, logMessage string) {
	return
}

func (mock *componentLogRepositoryMock) Warn(obj models.LogFields, logMessage string) {
	return
}
