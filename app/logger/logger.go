package logger

import (
	"io"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"gitlab.com/chaihanij/evat/app/env"
)

const LogFilePath = "logs/evat.log"

func Init() {
	if env.Debug {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
	} else {
		lumberjackLogrotate := &lumberjack.Logger{
			Filename:   LogFilePath,
			MaxSize:    10, // Max megabytes before log is rotated
			MaxBackups: 5,  // Max number of old log files to keep
			MaxAge:     30, // Max number of days to retain log files
			Compress:   true,
		}
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
		logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
		log.SetOutput(logMultiWriter)
	}

	log.WithFields(log.Fields{
		"RuntimeVersion": runtime.Version(),
		"NumberOfCPUs":   runtime.NumCPU(),
		"Arch":           runtime.GOARCH,
	}).Info("Application Initializing")
}
