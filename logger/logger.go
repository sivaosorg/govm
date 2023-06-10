package logger

import (
	"log"

	"github.com/sivaosorg/govm/utils"
)

func NewLogger() *Logger {
	l := &Logger{}
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
		if utils.IsEmpty(l.Filename) {
			log.Panic("Filename is required")
		}
	}
	if utils.IsNotEmpty(value) {
		l.Filename = utils.TrimSpaces(value)
	}
	return l
}

func (l *Logger) SetMaxSize(value int) *Logger {
	if value < 0 {
		log.Panic("Invalid max-size")
	}
	l.MaxSize = value
	return l
}

func (l *Logger) SetMaxAge(value int) *Logger {
	if value <= 0 {
		log.Panic("Invalid max-age")
	}
	l.MaxAge = value
	return l
}

func (l *Logger) SetMaxBackups(value int) *Logger {
	if value <= 0 {
		log.Panic("Invalid max-backups")
	}
	l.MaxBackups = value
	return l
}

func LoggerValidator(l *Logger) {
	l.
		SetMaxSize(l.MaxSize).
		SetMaxAge(l.MaxAge).
		SetMaxBackups(l.MaxBackups).
		SetFilename(l.Filename)
}
