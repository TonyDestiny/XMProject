package handler

import (
	"github.com/sirupsen/logrus"
)

func newErrorResponse(message string) {
	logrus.Error(message)
}
