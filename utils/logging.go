package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

func NewLogger() *Logger {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{errorLog: errorLog, infoLog: infoLog}
}

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (logger *Logger) ServerError(err error, w http.ResponseWriter) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	logger.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (logger *Logger) Info(message string) {
	logger.infoLog.Println(message)
}
