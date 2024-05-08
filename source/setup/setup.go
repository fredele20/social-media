package setup

import (
	"github.com/fredele20/social-media/source/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Conf *config.Config

var Logget = logrus.New()

func LoadEnv() {
	var err error
	Conf, err = config.Load()
	if err != nil {
		panic(err)
	}
}

func setGinLogLevel() {
	logLevel := "dev"
	switch env := Conf.Env; env {
	case "dev":
		logLevel = gin.DebugMode
	case "stage":
		logLevel = gin.TestMode
	case "prod":
		logLevel = gin.ReleaseMode
	}
	gin.SetMode(logLevel)
}

// func init() {
// 	ConnectDB()
// }
