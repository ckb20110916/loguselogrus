package loguselogrus

import (
	nestedlogrusformatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/ckb20110916/lumberjacklogwriter"
	"github.com/ckb20110916/rotatelogswriter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var (
	Logger = logrus.New()
)

var (
	logFormatter logrus.Formatter
)

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	EnableConsole()
}

func SetReportCaller(reportCaller bool) {
	Logger.SetReportCaller(reportCaller)
}

func SetFormatter(formatter logrus.Formatter) {
	Logger.SetFormatter(formatter)
	logFormatter = formatter
}

func EnableTrace() {
	Logger.SetLevel(logrus.TraceLevel)
}

func EnableDebug() {
	Logger.SetLevel(logrus.DebugLevel)
}

func EnableInfo() {
	Logger.SetLevel(logrus.InfoLevel)
}

func EnableWarn() {
	Logger.SetLevel(logrus.WarnLevel)
}

func EnableError() {
	Logger.SetLevel(logrus.ErrorLevel)
}

func EnableFatal() {
	Logger.SetLevel(logrus.FatalLevel)
}

func EnableConsole() {
	Logger.SetOutput(os.Stdout)
	if logFormatter != nil {
		Logger.SetFormatter(logFormatter)
	} else {
		Logger.SetFormatter(&nestedlogrusformatter.Formatter{
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: time.RFC3339,
		})
	}
}

func enableOutfile(logWriter io.Writer) {
	Logger.SetOutput(logWriter)
	if logFormatter != nil {
		Logger.SetFormatter(logFormatter)
	} else {
		Logger.SetFormatter(&nestedlogrusformatter.Formatter{
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: time.RFC3339,
			NoFieldsColors:  true,
			NoColors:        true,
		})
	}
}

func EnableLogFile(folder, filename string, maxAge, rotateTime time.Duration) {
	logWriter := rotatelogswriter.New(folder, filename, maxAge, rotateTime)
	if logWriter == nil {
		EnableConsole()
	} else {
		enableOutfile(logWriter)
	}
}

func EnableLogFile2(folder, filename string, maxBackups, maxSize, maxAge int) {
	logWriter := lumberjacklogwriter.New(folder, filename, maxBackups, maxSize, maxAge)
	if logWriter == nil {
		EnableConsole()
	} else {
		enableOutfile(logWriter)
	}
}
