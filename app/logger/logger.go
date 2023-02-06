package logger

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"gitlab.com/chaihanij/evat/app/env"
)

func Init() {

	if env.Debug {
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	} else {
		os.MkdirAll(env.LogPath, os.ModePerm)
		logFilePath := filepath.Join(env.LogPath, "evat-emc-app.log")
		lumberjackLogrotate := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    10, // Max megabytes before log is rotated
			MaxBackups: 5,  // Max number of old log files to keep
			MaxAge:     30, // Max number of days to retain log files
			Compress:   true,
		}
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
		logMultiWriter := io.MultiWriter(lumberjackLogrotate)
		log.SetOutput(logMultiWriter)
	}

	log.WithFields(log.Fields{
		"RuntimeVersion": runtime.Version(),
		"NumberOfCPUs":   runtime.NumCPU(),
		"Arch":           runtime.GOARCH,
	}).Info("Application Initializing")
}
