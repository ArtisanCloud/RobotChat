package logger

import (
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"net/http"
	"os"
	"testing"
)

var strArtisanCloudPath = "/var/log/ArtisanCloud/PowerLibs"
var strOutputPath = strArtisanCloudPath + "/output.log"
var strErrorPath = strArtisanCloudPath + "/errors.log"

func init() {
	err := initLogPath(strArtisanCloudPath, strOutputPath, strErrorPath)
	if err != nil {
		panic(err)
	}
}

func Test_Log_Info(t *testing.T) {
	driver := "zap"
	logger, err := NewLogger(driver, rcconfig.Log{
		Env:      "test",
		InfoLog:  strOutputPath,
		ErrorLog: strErrorPath,
	})
	if err != nil {
		t.Error(err)
	}

	logger.Info("test info", "app response", &http.Response{})

}

func initLogPath(path string, files ...string) (err error) {
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
