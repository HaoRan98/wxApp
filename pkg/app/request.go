package app

import (
	"github.com/astaxie/beego/validation"

	"wxApp/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(err.Key, err.Message)
	}
	return
}
