package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/db"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/logs"
	"github.com/involvestecnologia/statuspage/middleware"
	"github.com/involvestecnologia/statuspage/prometheus"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {

	log.Out = os.Stdout

	mgouri, exist := os.LookupEnv("MONGO_URI")
	if !exist {
		log.Panic("MongoDB URI not informed")
	}

	env, exist := os.LookupEnv("ENV_MODE")
	if !exist {
		log.Panic("MongoDB URI not informed")
	}
	if env == "production" {
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.ErrorLevel)
	}

	session := db.InitMongo(mgouri)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/v1")

	// Initialize repositories
	componentRepository := component.NewMongoRepository(session)
	incidentRepository := incident.NewMongoRepository(session)
	clientRepository := client.NewMongoRepository(session)
	componentLogRepository := logs.NewLogRepository(log)

	// Initialize services
	componentService := component.NewService(componentRepository, componentLogRepository)
	incidentService := incident.NewService(incidentRepository, componentService)
	clientService := client.NewService(clientRepository, componentService)

	// Initialize routers
	component.Router(componentService, v1)
	incident.Router(incidentService, v1)
	client.Router(clientService, v1)
	prometheus.Router(incidentService, componentService, v1)

	router.Run(":8080")
}
