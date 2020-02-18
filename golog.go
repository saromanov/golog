package golog

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Level defines logger level
type Level uint

const (
	All Level = iota + 1
	Debug
	Info
	Warning
	Error
	Fatal
	Panic
)

// GoLog implements wrapper over logs
type GoLog struct {
	logger       *logrus.Logger
	minShowLevel Level
}

// New creates GoLog
func New() *GoLog {
	return &GoLog{
		logger: logrus.New(),
	}
}

// Before implements steps for running before starting of logger
func (g *GoLog) Before(f func(l *logrus.Logger)) {
	if f == nil {
		g.beforeDefault()
		return
	}
	f(g.logger)
}

func (g *GoLog) beforeDefault() {
	g.logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	g.logger.SetOutput(os.Stdout)
}

// Fatalf for fatal errors
func (g *GoLog) Fatalf(format string, data ...interface{}) {
	if g.minShowLevel > Fatal {
		return
	}
	g.logger.Fatalf(format, data...)
}

// Infof for info errors
func (g *GoLog) Infof(format string, data ...interface{}) {
	if g.minShowLevel > Info {
		return
	}
	g.logger.Infof(format, data...)
}

// Errorf for errors with "Error" level
func (g *GoLog) Errorf(format string, data ...interface{}) {
	if g.minShowLevel > Error {
		return
	}
	g.logger.Errorf(format, data...)
}

// Panicf for panic errors
func (g *GoLog) Panicf(format string, data ...interface{}) {
	if g.minShowLevel > Panic {
		return
	}
	g.logger.Panicf(format, data...)
}

// Warningf for warning errors
func (g *GoLog) Warningf(format string, data ...interface{}) {
	if g.minShowLevel > Warning {
		return
	}
	g.logger.Warningf(format, data...)
}
