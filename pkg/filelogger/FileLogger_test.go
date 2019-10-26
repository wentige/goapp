package filelogger

import (
	"testing"
)

func Test_FileLogger(t *testing.T) {
	logger := New("test.log")

	logger.Print("logger.Print")
	logger.Emerg("logger.Emerg")
	logger.Alert("logger.Alert")
	logger.Crit("logger.Crit")
	logger.Error("logger.Error")
	logger.Warn("logger.Warn")
	logger.Notice("logger.Notice")
	logger.Info("logger.Info")
	logger.Debug("logger.Debug")

	logger.Flush()
}
