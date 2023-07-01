package logger

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger/contract"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/logger/drivers/zap"
	"os"
)

type Logger struct {
	Driver contract.LoggerInterface
}

func NewLogger(logConf rcconfig.Log) (logger *Logger, err error) {

	var driverLogger contract.LoggerInterface
	if logConf.Driver == "" || logConf.Driver == "zap" {
		env := logConf.Env
		if env == "" {
			env = "develop"
		}

		infoLog := logConf.InfoLog
		if infoLog == "" {
			infoLog = "./robot.log"
		}

		errLog := logConf.ErrorLog
		if errLog == "" {
			errLog = "./robot.log"
		}

		driverLogger, err = zap.NewLogger(&object.HashMap{
			"env":        env,
			"outputPath": infoLog,
			"errorPath":  errLog,
		})
	}

	logger = &Logger{
		Driver: driverLogger,
	}

	return logger, err

}

func (log *Logger) Debug(msg string, v ...interface{}) {
	log.Driver.Debug(msg, v...)
}
func (log *Logger) Info(msg string, v ...interface{}) {
	log.Driver.Info(msg, v...)
}
func (log *Logger) Warn(msg string, v ...interface{}) {
	log.Driver.Warn(msg, v...)
}
func (log *Logger) Error(msg string, v ...interface{}) {
	log.Driver.Error(msg, v...)
}
func (log *Logger) Panic(msg string, v ...interface{}) {
	log.Driver.Panic(msg, v...)
}
func (log *Logger) Fatal(msg string, v ...interface{}) {
	log.Driver.Fatal(msg, v...)
}

func (log *Logger) DebugF(format string, args ...interface{}) {
	log.Driver.DebugF(format, args)
}
func (log *Logger) InfoF(format string, args ...interface{}) {
	log.Driver.InfoF(format, args)
}
func (log *Logger) WarnF(format string, args ...interface{}) {
	log.Driver.WarnF(format, args)
}
func (log *Logger) ErrorF(format string, args ...interface{}) {
	log.Driver.ErrorF(format, args)
}
func (log *Logger) PanicF(format string, args ...interface{}) {
	log.Driver.PanicF(format, args)
}
func (log *Logger) FatalF(format string, args ...interface{}) {
	log.Driver.FatalF(format, args)
}

func InitLogPath(path string, files ...string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	} else if os.IsPermission(err) {
		return err
	}

	for _, fileName := range files {
		if _, err = os.Stat(fileName); os.IsNotExist(err) {
			_, err = os.Create(fileName)
			if err != nil {
				return err
			}
		}
	}

	return err

}
