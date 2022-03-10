package blog

import (
	"go.uber.org/zap"
)

type BankerLogger interface {
	// AddCallerSkip new cube logger with callstack skipping.
	AddCallerSkip(callerSkip int) BankerLogger

	// WithName adds some key-value pairs of context to a logger.
	// See Info for documentation on how key/value pairs work.
	WithName(name string) BankerLogger

	// WithValues adds a new element to the logger's name.
	// Successive calls with WithName continue to append
	// suffixes to the logger's name.  It's strongly recommended
	// that name segments contain only letters, digits, and hyphens
	// (see the package documentation for more information).
	WithValues(keysAndValues ...interface{}) BankerLogger

	Debug(format string, a ...interface{})

	Info(format string, a ...interface{})

	Warn(format string, a ...interface{})

	Error(format string, a ...interface{})

	Fatal(format string, a ...interface{})
}

type bankerLogger struct {
	l *zap.Logger
}

var logger BankerLogger

func Debug(format string, a ...interface{}) {
	ensureLogger().Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	ensureLogger().Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	ensureLogger().Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	ensureLogger().Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	ensureLogger().Fatal(format, a...)
}

func WithName(name string) BankerLogger {
	return ensureLogger().WithName(name).AddCallerSkip(-1)
}

func WithValues(keysAndValues ...interface{}) BankerLogger {
	return ensureLogger().WithValues(keysAndValues).AddCallerSkip(-1)
}

// ensureLogger new default banker logger if logger is nil
func ensureLogger() BankerLogger {
	if logger == nil {
		logger = newDefaultBankerLogger()
	}
	return logger
}
