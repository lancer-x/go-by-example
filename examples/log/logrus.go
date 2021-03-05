package main

import (
	log "github.com/sirupsen/logrus"
)
func main2() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.WithFields(log.Fields{
		"animal": "walurs",
	}).Info("A walrus appears")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other": "it will be logged too",
	})
	contextLogger.Info("aaaaaaa")
}
