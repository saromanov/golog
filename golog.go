package golog

import (
	"os"
	"github.com/sirupsen/logrus"
)

// GoLog implements wrapper over logs
type GoLog struct {
	logger    *logrus.Logger
	showLevel string
}

// New creates GoLog
func New() *GoLog {
	return &GoLog{
		logger: logrus.New(),
	}
}

// Before implements steps for running before starting of logger
func (g*GoLog) Before(f func(l *logrus.Logger)){
	if f == nil {
		g.beforeDefault()
		return
	}
	f(g.logger)
}

func (g *GoLog) beforeDefault(){
	g.logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	g.logger.SetOutput(os.Stdout)
}

// Fatalf for fatal errors
func (l *GoLog) Fatalf(format string, data ...interface{}) {
	l.logger.Fatalf(format, data...)
}

// Infof for info errors
func (l *GoLog) Infof(format string, data ...interface{}) {
	l.logger.Infof(format, data...)
}

// Errorf for errors with "Error" level
func (l *GoLog) Errorf(format string, data ...interface{}) {
	l.logger.Errorf(format, data...)
}
