package logger

import (
	"free5gc-cli/lib/logger_util"
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var AppLog *logrus.Entry
var InitLog *logrus.Entry
var FreecliLog *logrus.Entry
var SubscriberLog *logrus.Entry
var GNBLog *logrus.Entry
var NFLog *logrus.Entry

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

	selfLogHook, err := logger_util.NewFileHook(
		"logs/freecli.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	AppLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "App"})
	InitLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "Init"})
	FreecliLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "Freecli"})
	SubscriberLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "Subscriber Module"})
	GNBLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "gNB Module"})
	NFLog = log.WithFields(logrus.Fields{"component": "Freecli", "category": "NF Module"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	log.SetReportCaller(bool)
}
