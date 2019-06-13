package mock

import (
	"log"
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
	log.Printf("[INFO] %v+", obj)
}

func (mock *componentLogRepositoryMock) Error(obj models.LogFields, logMessage string) {
	log.Printf("[ERROR] %v+", obj)
}

func (mock *componentLogRepositoryMock) Warn(obj models.LogFields, logMessage string) {
	log.Printf("[WARN] %v+", obj)
}

func (mock *componentLogRepositoryMock) Debug(obj models.LogFields, logMessage string) {
	log.Printf("[DEBUG] %v+", obj)
}