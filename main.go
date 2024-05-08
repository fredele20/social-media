package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fredele20/social-media/source/routes"
	"github.com/fredele20/social-media/source/setup"
	"github.com/gin-gonic/gin"
)

func init() {
	setup.LoadEnv()
	setup.ConnectDB()
}

func incrementAppVersion(currentVersion string) (string, error) {
	incrementVersion := "0.0.1"
	splitedCurrentVersion := strings.Split(currentVersion, ".")
	splitedIncrementVersion := strings.Split(incrementVersion, ".")

	if len(splitedCurrentVersion) != 3 || len(splitedIncrementVersion) != 3 {
		return "", errors.New("values must be more than 3")
	}

	var currentVersionInts, incrementVersionInts []int

	for i := range splitedCurrentVersion {
		currentVersionNum, err := strconv.Atoi(splitedCurrentVersion[i])
		if err != nil {
			return "", err
		}
		currentVersionInts = append(currentVersionInts, currentVersionNum)

		incrementVersionNum, err := strconv.Atoi(splitedIncrementVersion[i])
		if err != nil {
			return "", err
		}

		incrementVersionInts = append(incrementVersionInts, incrementVersionNum)
	}

	for i := range currentVersionInts {
		currentVersionInts[i] += incrementVersionInts[i]
	}

	newVersion := fmt.Sprintf("%d.%d.%d", currentVersionInts[0], currentVersionInts[1], currentVersionInts[2])

	return newVersion, nil
}

func main() {
	version, _ := incrementAppVersion("4.2.10")
	fmt.Println("Version: ", version)
	// secrets := config.GetSecrets()
	// logger := logrus.New()

	// db, _ := mongod.ConnectDB(secrets.DatabaseURL, secrets.DatabaseName)
	// address := fmt.Sprintf("127.0.0.1:%s", secrets.HttpPort)

	// fileLogger := "logs.log"
	// logFile, err := os.OpenFile(fileLogger, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Printf("error opening log file: %s", err)
	// 	return
	// }

	// defer logFile.Close()
	// logrus.SetOutput(logFile)
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	// fmt.Println("log file created")

	// session := sessions.NewSessionManager(logger, secrets.JwtSecretKey, db)

	// core := core.NewCoreService(db, logger, session)
	// routes := routes.NewRoutesService(core)
	// handler := handlers.NewHandler(routes)
	// handlers.UserHandler(router, *handler)

	address := fmt.Sprintf("127.0.0.1:%s", setup.Conf.HttpPort)

	router := gin.New()
	router.Use(gin.Logger())

	routes.RouteHandlers(router)

	err := router.Run(address)

	if err != nil {
		panic(err)
	}

}
