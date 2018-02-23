package internal

import (
	"fmt"
	"github.com/vendasta/gosdks/util"
)

func newInvalidFieldError(format string, a ...interface{}) error {
	return util.Error(util.InvalidArgument, fmt.Sprintf(format, a...))
}
