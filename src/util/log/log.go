package log

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/TwiN/go-color"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	echologrusLog "github.com/neko-neko/echo-logrus/v2/log"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"blog/config"
	middleware "blog/middleware/log"
)

var Logger *logrus.Logger

func init() {
	Logger = getLogger()
	Logger.Info("logger started")
}

func getLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(formatter())
	logger.SetLevel(logrus.InfoLevel)
	if config.C.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	baseLogPath := path.Join(config.C.LogConf.LogPath, config.C.LogConf.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logger.Fatal(err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})

	logger.AddHook(lfHook)
	return logger
}

func formatter() *nested.Formatter {
	fmtter := &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "15:04:05",
		CallerFirst:     false,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(color.InBlue(" ⇨ %v (line%v)"), filepath.Base(fullPath), line)
		},
	}
	fmtter.NoColors = false
	return fmtter
}

// SetLoggerOfEcho switches the default logger of Echo to logrus with custom format.
func SetLoggerOfEcho(e *echo.Echo) {
	echologrusLog.Logger().SetOutput(os.Stdout)
	echologrusLog.Logger().SetLevel(echoLog.INFO)
	if config.C.Debug {
		echologrusLog.Logger().SetLevel(echoLog.DEBUG)
	}
	echologrusLog.Logger().SetFormatter(formatter())
	e.Logger = echologrusLog.Logger()
	e.Use(middleware.GetLogger())
}
