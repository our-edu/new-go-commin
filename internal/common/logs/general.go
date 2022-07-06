package logs

import (
	"github.com/sirupsen/logrus"
)

func LogError(msg string, err error) {

	logrus.WithError(err).Error(msg)

}
func LogInfo(msg interface{}) {

	logrus.Info(msg)

}
