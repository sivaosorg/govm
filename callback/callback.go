package callback

import (
	"reflect"
	"runtime"
	"time"

	"github.com/sivaosorg/govm/logger"
)

type CallbackFunc func()
type CallbackFallbackFunc func() interface{}
type CallbackFallbacksFunc func() []interface{}

func MeasureTime(callback CallbackFunc) {
	_func := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	start := time.Now()
	callback()
	end := time.Since(start)
	logger.Successf("Func:: %s elapsed time:: %s", _func, end)
}

func MeasureTimeFallback(callback CallbackFallbackFunc) interface{} {
	_func := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	start := time.Now()
	var result interface{} = nil

	defer func() {
		if err := recover(); err != nil {
			logger.Warnf("Panic occurred during func:: %s execution:: %v", _func, err)
		}
	}()
	defer func() {
		end := time.Since(start)
		logger.Successf("Func:: %s elapsed time:: %s", _func, end)
	}()
	result = callback()
	return result
}

func MeasureTimeFallbacks(callback CallbackFallbacksFunc) []interface{} {
	_func := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	start := time.Now()
	var result []interface{} = nil

	defer func() {
		if err := recover(); err != nil {
			logger.Warnf("Panic occurred during func:: %s execution:: %v", _func, err)
		}
	}()
	defer func() {
		end := time.Since(start)
		logger.Successf("Func:: %s elapsed time:: %s", _func, end)
	}()
	result = callback()
	return result
}
