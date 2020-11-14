package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var MongoDBLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	MongoDBLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "MongoDB"})
}

func SetLogLevel(level logrus.Level) {
	MongoDBLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	MongoDBLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}
