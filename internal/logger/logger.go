package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New()

	// Salida a archivo
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir archivo de logs:", err)
	}
	Logger.Out = file

	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetLevel(log.DebugLevel)
}

// Funciones helper
func LogInfo(msg string, fields log.Fields) {
	Logger.WithFields(fields).Info(msg)
}

func LogError(msg string, fields log.Fields) {
	Logger.WithFields(fields).Error(msg)
}
