package logger

import (
	"os"

	"github.com/sirupsen/logrus"

	"go-rest-project/config"
)

const callerField = "caller"

var log *logrus.Logger

// Init initialize logger
func Init() (*logrus.Logger, error) {
	file, err := os.OpenFile(config.Config.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log = logrus.New()

	lvl, err := logrus.ParseLevel(config.Config.LogLevel)
	if err != nil {
		return nil, err
	}

	log.Level = lvl
	log.Out = file
	return log, nil
}

func LogEntry(caller string) *logrus.Entry {
	return log.WithFields(map[string]interface{}{callerField: caller})
}
