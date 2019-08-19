package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/LandvibeDev/gofka/collector/config"
	"github.com/LandvibeDev/gofka/collector/kafka"
	"github.com/LandvibeDev/gofka/collector/router"
	"github.com/LandvibeDev/gofka/collector/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Use multi core
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load Configuration
	gofkaConfig := config.LoadConfiguration(e.Logger)

	// Connect Kafka
	_, err := kafka.EnsureTopic(gofkaConfig.Kafka)
	if err != nil {
		e.Logger.Fatal(err)
	}

	producer, err := kafka.GetProducer(gofkaConfig.Kafka)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer producer.Close()

	// Create Service
	logService := service.NewLogService(producer)

	// Create Router
	v1 := e.Group("/api/v1")
	h := router.NewHandler(logService)
	h.Register(v1)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(gofkaConfig.Server.Port)))
}
