package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func LoggingToFile() {
	file, err := os.OpenFile("dataLog.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)
}
