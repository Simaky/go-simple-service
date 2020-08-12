package web

import "go-rest-project/logger"

const (
	logTestPath  = "../go-simple-service-test.log"
	logTestLevel = "INFO"
)

func init() {
	err := logger.LoadForTest(logTestPath, logTestLevel)
	if err != nil {
		panic(err)
	}
}
