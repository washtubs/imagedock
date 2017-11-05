package imagedock

import (
	"os"

	logging "github.com/op/go-logging"
)

var log *logging.Logger = GetLogger()

func GetLogger() *logging.Logger {
	newLogger := logging.MustGetLogger("imagedock")
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	leveledLogBackend := logging.AddModuleLevel(logBackend)
	leveledLogBackend.SetLevel(logging.DEBUG, "")
	logging.SetBackend(leveledLogBackend)

	logging.SetFormatter(logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	))
	return newLogger
}
