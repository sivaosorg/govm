package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"github.com/sivaosorg/govm/timex"
	"github.com/sivaosorg/govm/utils"
)

var logger *Logger

func NewLogger() *Logger {
	if logger != nil {
		return logger
	}
	l := &Logger{}
	l.SetEnabled(true)
	l.SetFormatter(LoggerTextFormatter)
	l.SetInstance(l.NewInstance())
	logger = l
	return l
}

func NewLoggerWith(filename string, maxSize, maxBackups, maxAge int) *Logger {
	if logger != nil {
		return logger
	}
	l := &Logger{}
	l.SetEnabled(true)
	l.SetAllowSnapshot(true)
	l.SetFormatter(LoggerJsonFormatter)
	l.SetFilename(filename)
	l.SetMaxAge(maxAge)
	l.SetMaxBackups(maxBackups)
	l.SetMaxSize(maxSize)
	l.SetInstance(l.NewInstance())
	logger = l
	return l
}

func (l *Logger) NewInstance() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(io.MultiWriter(l.Config(), os.Stdout))
	if strings.EqualFold(LoggerJsonFormatter, l.Formatter) {
		logger.SetFormatter(l.JsonFormatter())
	}
	if strings.EqualFold(LoggerTextFormatter, l.Formatter) {
		logger.SetFormatter(l.TextFormatter())
	}
	if l.AllowSnapshot {
		logger.SetFormatter(l.JsonFormatter())
	}
	logger.AddHook(l.TextHook())
	return logger
}

func (l *Logger) TextHook() *TextFormatterHook {
	successHex := color.FgGreen
	infoHex := color.FgHiBlue
	warnHex := color.FgYellow
	errorHex := color.FgHiRed
	debugHex := color.FgCyan
	successColor := color.New(successHex)
	infoColor := color.New(infoHex)
	warnColor := color.New(warnHex)
	errorColor := color.New(errorHex)
	debugColor := color.New(debugHex)
	return &TextFormatterHook{
		success: successColor,
		info:    infoColor,
		warn:    warnColor,
		err:     errorColor,
		debug:   debugColor,
	}
}

func (l *Logger) JsonFormatter() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		DisableTimestamp:  false,
		PrettyPrint:       false,
		DisableHTMLEscape: false,
		TimestampFormat:   timex.DateTimeFormYearMonthDayHourMinuteSecond,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	}
}

func (l *Logger) TextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:    false,
		ForceColors:      true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		DisableQuote:     true,
		TimestampFormat:  timex.DateTimeFormYearMonthDayHourMinuteSecond,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	}
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
	if strings.EqualFold(LoggerJsonFormatter, l.Formatter) {
		l.instance.SetFormatter(l.JsonFormatter())
	}
	if strings.EqualFold(LoggerTextFormatter, l.Formatter) {
		l.instance.SetFormatter(l.TextFormatter())
	}
	if l.AllowSnapshot {
		l.instance.SetFormatter(l.JsonFormatter())
	}
	l.ResetLogger()
	return l
}

func (l *Logger) ResetLogger() {
	logger = nil
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

func (l *Logger) SetFormatter(value string) *Logger {
	if utils.IsNotEmpty(value) {
		l.Formatter = utils.TrimAllSpaces(value)
	}
	return l
}

func (l *Logger) SetAllowCaller(value bool) *Logger {
	l.AllowCaller = value
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

func (l *Logger) Callers() (filename, function string, line int) {
	_, file, line, _ := runtime.Caller(2)
	filename = filepath.Base(file)
	return filename, function, line
}

func (l *Logger) CallerString() string {
	filename, _, line := l.Callers()
	return fmt.Sprintf("%s:%d", filename, line)
}

func (h *TextFormatterHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.InfoLevel:
		entry.Message = h.info.Sprint(entry.Message)
	case logrus.WarnLevel:
		entry.Message = h.warn.Sprint(entry.Message)
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		entry.Message = h.err.Sprint(entry.Message)
	case logrus.DebugLevel, logrus.TraceLevel:
		entry.Message = h.debug.Sprint(entry.Message)
	}
	return nil
}

func (h *TextFormatterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *Logger) Info(message string, params ...interface{}) {
	if !l.IsEnabled {
		return
	}
	if len(params) > 0 {
		var _p []interface{}
		for _, v := range params {
			if utils.IsPrimitiveType(v) {
				_p = append(_p, v)
			} else {
				_p = append(_p, utils.ToJson(v))
			}
		}
		params = _p
	}
	var fields logrus.Fields
	filename, _, line := l.Callers()
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+1)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
	if !l.IsEnabled {
		return
	}
	if len(params) > 0 {
		var _p []interface{}
		for _, v := range params {
			if utils.IsPrimitiveType(v) {
				_p = append(_p, v)
			} else {
				_p = append(_p, utils.ToJson(v))
			}
		}
		params = _p
	}
	var fields logrus.Fields
	filename, _, line := l.Callers()
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+2)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
		if err != nil {
			fields[LoggerErrorField] = err.Error()
		}
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
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
		if err != nil {
			fields[LoggerErrorField] = err.Error()
		}
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
	if !l.IsEnabled {
		return
	}
	if len(params) > 0 {
		var _p []interface{}
		for _, v := range params {
			if utils.IsPrimitiveType(v) {
				_p = append(_p, v)
			} else {
				_p = append(_p, utils.ToJson(v))
			}
		}
		params = _p
	}
	var fields logrus.Fields
	filename, _, line := l.Callers()
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+1)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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

func (l *Logger) Debug(message string, params ...interface{}) {
	if !l.IsEnabled {
		return
	}
	if len(params) > 0 {
		var _p []interface{}
		for _, v := range params {
			if utils.IsPrimitiveType(v) {
				_p = append(_p, v)
			} else {
				_p = append(_p, utils.ToJson(v))
			}
		}
		params = _p
	}
	var fields logrus.Fields
	filename, _, line := l.Callers()
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+1)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
	l.instance.WithFields(fields).Debug()
}

func (l *Logger) Success(message string, params ...interface{}) {
	if !l.IsEnabled {
		return
	}
	if len(params) > 0 {
		var _p []interface{}
		for _, v := range params {
			if utils.IsPrimitiveType(v) {
				_p = append(_p, v)
			} else {
				_p = append(_p, utils.ToJson(v))
			}
		}
		params = _p
	}
	var fields logrus.Fields
	filename, _, line := l.Callers()
	if strings.Contains(message, "%") {
		fields = make(logrus.Fields, (len(params)/2)+2)
		fields[LoggerMessageField] = fmt.Sprintf(message, params...)
		fields[LoggerSuccessField] = true
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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
		if l.AllowCaller {
			fields[LoggerCallerField] = fmt.Sprintf("%s:%d", filename, line)
		}
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

func Infof(message string, params ...interface{}) {
	NewLogger().Info(message, params...)
}

func Errorf(message string, err error, params ...interface{}) {
	NewLogger().Error(message, err, params...)
}

func Warnf(message string, params ...interface{}) {
	NewLogger().Warn(message, params...)
}

func Successf(message string, params ...interface{}) {
	NewLogger().Success(message, params...)
}

func Debugf(message string, params ...interface{}) {
	NewLogger().Debug(message, params...)
}