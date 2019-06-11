package main

import (
	"github.com/charlesakalugwu/go-http-metrics/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := getLogger()
	log.Info("starting the server")
	s := server.NewServer(log, "", 8080)
	s.Run()
}

func getLogger() *logrus.Entry {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetReportCaller(logrus.GetLevel() == logrus.DebugLevel)
	return logrus.NewEntry(logrus.StandardLogger())
}
