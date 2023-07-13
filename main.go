package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fredele20/social-media/config"
	"github.com/fredele20/social-media/core"
	"github.com/fredele20/social-media/database/mongod"
	"github.com/fredele20/social-media/handlers"
	"github.com/fredele20/social-media/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	secrets := config.GetSecrets()
	logger := logrus.New()

	db, _ := mongod.ConnectDB(secrets.DatabaseURL, secrets.DatabaseName)
	address := fmt.Sprintf("127.0.0.1:%s", secrets.HttpPort)

	fileLogger := "logs.log"
	logFile, err := os.OpenFile(fileLogger, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("error opening log file: %s", err)
		return
	}

	defer logFile.Close()
	logrus.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	fmt.Println("log file created")

	router := gin.New()
	router.Use(gin.Logger())

	core := core.NewCoreService(db, logger)
	routes := routes.NewRoutesService(core)
	handler := handlers.NewHandler(routes)

	handlers.UserHandler(router, *handler)

	router.Run(address)

}