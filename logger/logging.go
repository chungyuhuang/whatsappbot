package logger

import (
	"github.com/op/go-logging"
	"io"
	"log"
	"os"
)

var format = logging.MustStringFormatter(
	`%{color:bold}[%{level:.8s}]%{color:reset} %{color}%{time:2006-01-02T15:04:05Z07:00} %{shortfile} %{shortfunc}%{color:reset} ▶ %{id:03x} %{message}`,
)

/* sets up logging directing it to the given log file */
func SetupLogging(serviceLogs string, LOGGER *logging.Logger) (*os.File, error) {
	LOGGER.Infof("Log File: %s", serviceLogs)
	f, err := os.OpenFile(serviceLogs, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v: %v", err, serviceLogs)
		return nil, err
	}
	logWriter := io.MultiWriter(f, os.Stdout)

	backend1 := logging.NewLogBackend(logWriter, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend2Formatter)
	return f, nil
}