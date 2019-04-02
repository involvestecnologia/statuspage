package mock

import (
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/models"
)

type componentLogRepositoryMock struct {
}

func NewComponentLogRepositoryMock() component.Log {
	return &componentLogRepositoryMock{}
}

func (mock *componentLogRepositoryMock) Info(obj models.Component, logMessage string) {
	return
}

func (mock *componentLogRepositoryMock) Error(obj models.Component, logMessage string) {
	return
}

func (mock *componentLogRepositoryMock) Warn(obj models.Component, logMessage string) {
	return
}
