package logs

import (
	"github.com/sirupsen/logrus"
)

func LogRepoExecution(tableName string, data interface{}, err error) {
	log := logrus.WithField("data", data)

	if err == nil {
		log.Info("table: " + tableName + " execute succeeded")
	} else {
		log.WithError(err).Error("table: " + tableName + " execute failed")
	}
}
