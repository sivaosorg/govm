package logger

import (
	"github.com/fatih/color"

	"github.com/sirupsen/logrus"
)

// Logger opens or creates the log file on first Write.  If the file exists and
// is less than MaxSize megabytes, lumberjack will open and append to that file.
// If the file exists and its size is >= MaxSize megabytes, the file is renamed
// by putting the current time in a timestamp in the name immediately before the
// file's extension (or the end of the filename if there's no extension). A new
// log file is then created using original filename.
//
// Whenever a write would cause the current log file exceed MaxSize megabytes,
// the current file is closed, renamed, and a new log file created with the
// original name. Thus, the filename you give Logger is always the "current" log
// file.
//
// Backups use the log file name given to Logger, in the form
// `name-timestamp.ext` where name is the filename without the extension,
// timestamp is the time at which the log was rotated formatted with the
// time.Time format of `2006-01-02T15-04-05.000` and the extension is the
// original extension.  For example, if your Logger.Filename is
// `/var/log/foo/server.log`, a backup created at 6:30pm on Nov 11 2016 would
// use the filename `/var/log/foo/server-2016-11-04T18-30-00.000.log`
//
// # Cleaning Up Old Log Files
//
// Whenever a new log file gets created, old log files may be deleted.  The most
// recent files according to the encoded timestamp will be retained, up to a
// number equal to MaxBackups (or all of them if MaxBackups is 0).  Any files
// with an encoded timestamp older than MaxAge days are deleted, regardless of
// MaxBackups.  Note that the time encoded in the timestamp is the rotation
// time, which may differ from the last time that file was written to.
//
// If MaxBackups and MaxAge are both 0, no old log files will be deleted.
type Logger struct {
	instance *logrus.Logger `json:"-" yaml:"-"`
	// Enable to allow using logger
	IsEnabled bool `json:"enabled" yaml:"enabled"`
	// Allow to save log into file
	PermitSnapshot bool `json:"permit_snapshot" yaml:"permit_snapshot"`
	// PermitLocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	PermitLocalTime bool `json:"permit_local_time" yaml:"permit_local_time"`
	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
	// Allow show caller detail, just like this: @caller=logger.go:487
	PermitCaller bool `json:"permit_caller" yaml:"permit_caller"`
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <process_name>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"max_size" yaml:"max_size"`
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"max_age" yaml:"max_age"`
	// MaxBackup is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackup int    `json:"max_backup" yaml:"max_backup"`
	Formatter string `json:"formatter" yaml:"formatter"`
}

type TextFormatterHook struct {
	success *color.Color
	info    *color.Color
	warn    *color.Color
	debug   *color.Color
	err     *color.Color
}

type loggerOptionConfig struct {
	MaxRetries int `json:"max_retries" yaml:"max-retries"`
}

type MultiTenantLoggerConfig struct {
	Key             string             `json:"key" binding:"required" yaml:"key"`
	IsUsableDefault bool               `json:"usable_default" yaml:"usable_default"`
	Config          Logger             `json:"config" yaml:"config"`
	Option          loggerOptionConfig `json:"option" binding:"required" yaml:"option"`
}

type ClusterMultiTenantLoggerConfig struct {
	Clusters []MultiTenantLoggerConfig `json:"clusters,omitempty" yaml:"clusters"`
}
