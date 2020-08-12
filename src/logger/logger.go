package logger

import (
	"os"

	"github.com/sirupsen/logrus"

	"go-rest-project/config"
)

const callerField = "caller"

var log *logrus.Logger

// Load initialize logger
func Load() (*logrus.Logger, error) {
	return loadLogger(config.Config.LogFilePath, config.Config.LogLevel)
}

// LoadForTest initialize logger for tests
func LoadForTest(path, level string) error {
	_, err := loadLogger(path, level)
	return err
}

// LogEntry logger with caller field
func LogEntry(caller string) *logrus.Entry {
	return log.WithFields(map[string]interface{}{callerField: caller})
}

func loadLogger(path, level string) (*logrus.Logger, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log = logrus.New()

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	log.Level = lvl
	log.Out = file
	return log, nil
}
