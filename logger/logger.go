package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/sivaosorg/govm/utils"
)

func NewLogger() *Logger {
	l := &Logger{}
	l.SetEnabled(true)
	l.SetInstance(l.NewInstance())
	return l
}

func NewLoggerWith(filename string, maxSize, maxBackups, maxAge int) *Logger {
	l := &Logger{}
	l.SetEnabled(true)
	l.SetAllowSnapshot(true)
	l.SetFilename(filename)
	l.SetMaxAge(maxAge)
	l.SetMaxBackups(maxBackups)
	l.SetMaxSize(maxSize)
	l.SetInstance(l.NewInstance())
	return l
}

func (l *Logger) NewInstance() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(l.Config(), os.Stdout))
	return logger
}

func (l *Logger) Config() *lumberjack.Logger {
	LoggerValidator(l)
	j := &lumberjack.Logger{}
	if l.AllowSnapshot {
		j.Filename = l.Filename
		j.MaxSize = l.MaxSize
		j.MaxBackups = l.MaxBackups
		j.MaxAge = l.MaxAge
		j.Compress = l.Compress
		j.LocalTime = l.AllowLocalTime
	}
	return j
}

func (l *Logger) ApplyConfig() *Logger {
	if l.instance == nil {
		l.SetInstance(l.NewInstance())
	} else {
		l.instance.SetOutput(io.MultiWriter(l.Config(), os.Stdout))
	}
	return l
}

func (l *Logger) SetInstance(value *logrus.Logger) *Logger {
	l.instance = value
	return l
}

func (l *Logger) SetEnabled(value bool) *Logger {
	l.IsEnabled = value
	return l
}

func (l *Logger) SetAllowSnapshot(value bool) *Logger {
	l.AllowSnapshot = value
	return l
}

func (l *Logger) SetAllowLocalTime(value bool) *Logger {
	l.AllowLocalTime = value
	return l
}

func (l *Logger) SetCompress(value bool) *Logger {
	l.Compress = value
	return l
}

func (l *Logger) SetFilename(value string) *Logger {
	if l.AllowSnapshot {
		if utils.IsEmpty(value) {
			log.Panic("Filename is required")
		}
	}
	if utils.IsNotEmpty(value) {
		l.Filename = utils.TrimSpaces(value)
	}
	return l
}

func (l *Logger) SetMaxSize(value int) *Logger {
	if l.AllowSnapshot {
		if value < 0 {
			log.Panic("Invalid max-size")
		}
	}
	l.MaxSize = value
	return l
}

func (l *Logger) SetMaxAge(value int) *Logger {
	if l.AllowSnapshot {
		if value <= 0 {
			log.Panic("Invalid max-age")
		}
	}
	l.MaxAge = value
	return l
}

func (l *Logger) SetMaxBackups(value int) *Logger {
	if l.AllowSnapshot {
		if value <= 0 {
			log.Panic("Invalid max-backups")
		}
	}
	l.MaxBackups = value
	return l
}

func (l *Logger) Json() string {
	return utils.ToJson(l)
}

func LoggerValidator(l *Logger) {
	if !l.IsEnabled {
		return
	}
	l.
		SetMaxSize(l.MaxSize).
		SetMaxAge(l.MaxAge).
		SetMaxBackups(l.MaxBackups).
		SetFilename(l.Filename)
}

func (l *Logger) Info(message string, params ...interface{}) {
	var fields logrus.Fields
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+1)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	} else {
		fields = make(logrus.Fields, 1)
		fields[LoggerMessageField] = message
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	}
	l.instance.WithFields(fields).Info()
}

func (l *Logger) Error(message string, err error, params ...interface{}) {
	var fields logrus.Fields
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+2)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		fields[LoggerErrorField] = err.Error()
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	} else {
		fields = make(logrus.Fields, 2)
		fields[LoggerMessageField] = message
		fields[LoggerErrorField] = err.Error()
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	}
	l.instance.WithFields(fields).Error()
}

func (l *Logger) Warn(message string, params ...interface{}) {
	var fields logrus.Fields
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+1)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	} else {
		fields = make(logrus.Fields, 1)
		fields[LoggerMessageField] = message
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	}
	l.instance.WithFields(fields).Warn()
}

func (l *Logger) Success(message string, params ...interface{}) {
	var fields logrus.Fields
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+2)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		fields[LoggerSuccessField] = true
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	} else {
		fields = make(logrus.Fields, 2)
		fields[LoggerMessageField] = message
		fields[LoggerSuccessField] = true
		for i := 0; i < len(params); i += 2 {
			if i+1 >= len(params) {
				break
			}
			key, ok := params[i].(string)
			if !ok {
				continue
			}
			fields[key] = params[i+1]
		}
	}
	l.instance.WithFields(fields).Info()
}
