package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func Init() error {
	var err error
	LogFile, err = os.OpenFile("SandLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier de log : %v", err)
	}

	Logger = log.New(LogFile, "", log.Ldate|log.Ltime)
	return nil
}

func LogAPICall(endpoint, method, userEmail string, statusCode int) {
	logMessage := fmt.Sprintf("[%s] %s %s - Utilisateur: %s - Statut: %d",
		time.Now().Format("2006-01-02 15:04:05"),
		method,
		endpoint,
		userEmail,
		statusCode)
	Logger.Println(logMessage)
}

func Close() {
	if LogFile != nil {
		LogFile.Close()
	}
}
