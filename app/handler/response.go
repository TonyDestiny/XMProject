package handler

import (
	"github.com/sirupsen/logrus"
)

func logError(message string) {
	logrus.Error(message)
}
