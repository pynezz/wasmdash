package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
)

func init() {
	logPath := os.Getenv("WLOGPATH")
	if logPath == "" {
		logPath = "wasmdash.log"
	}
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}

var MiddlewareLogger *log.Logger = log.New(os.Stdout, "", log.LstdFlags)

type Message struct {
	Level   Level
	Message string
}

type Logger struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) Logger {
	return Logger{
		logger: logger,
	}
}

func (l *Logger) writeErrorMsg(err string, ctx echo.Context) {
	host := os.Getenv("HOST")
	l.Printf("[ERROR] [%s] %s\n%s\n%s", host, err, ctx.Request())
}

func (l *Logger) writeDebugMsg(message string, ctx echo.Context) {
	details := map[string]string{
		"request": fmt.Sprintf("%s %s", ctx.Request().Method, ctx.Request().URL.Path),
	}
	l.Printf("[DEBUG] %s\n%s", message, details)
}

func (l *Logger) writeInfoMsg(message string, ctx echo.Context) {
	l.Printf("[INFO] %s", message)
}

func (l *Logger) writeWarningMsg(message string, ctx echo.Context) {
	l.Printf("[WARNING] %s\n%s", message)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Log(level Level, message string, ctx echo.Context) Message {
	switch level {
	case Debug:
		l.writeDebugMsg(message, ctx)
	case Info:
		l.writeInfoMsg(message, ctx)
	case Warning:
		l.writeWarningMsg(message, ctx)
	case Error:
		l.writeErrorMsg(message, ctx)
	default:
		l.logger.Printf("[UNKNOWN] %s", message)
	}

	return Message{
		Level:   level,
		Message: message,
	}
}

// Logger middleware for logging certain important stuff

func Log(c echo.Context) error {
	MiddlewareLogger.Printf("Request: %s %s", c.Request().Method, c.Request().URL.Path)
	return nil
}
