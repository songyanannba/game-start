package helper

import (
	"elim5/global"
	"errors"
	"go.uber.org/zap"
)

func ErrMerge(errs ...error) error {
	str := ""
	for _, err := range errs {
		if err == nil {
			continue
		}
		str += err.Error() + ". "
	}
	if str == "" {
		return nil
	}
	return errors.New(str)
}

func RecordError(err error, params ...any) {
	if err != nil {
		global.GVA_LOG.WithOptions(zap.AddCallerSkip(1)).Error(err.Error(), zap.Any("params", params))
	}
}

func FatalError(err error, params ...any) {
	if err != nil {
		global.GVA_LOG.WithOptions(zap.AddCallerSkip(1)).Fatal(err.Error(), zap.Any("params", params))
	}
}
