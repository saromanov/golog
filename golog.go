package golog

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Level defines logger level
type Level uint

const (
	All Level = iota + 1
	Trace
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
	fields       logrus.Fields
}

// New creates GoLog
func New(c *Config) *GoLog {
	l := logrus.New()
	r := &GoLog{
		logger: l,
		fields: logrus.Fields{},
	}
	if c != nil {
		r.minShowLevel = c.MinShowLevel
		if len(c.Hooks) > 0 {
			for _, h := range c.Hooks {
				l.AddHook(h)
			}
		}
	}
	return r
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
	g.output(g.makeData().Fatalf, format, data...)
}

// Infof for info errors
func (g *GoLog) Infof(format string, data ...interface{}) {
	g.output(g.makeData().Infof, format, data...)
}

func (g *GoLog) output(f func(string, ...interface{}), format string, data ...interface{}) {
	if g.minShowLevel > Info {
		return
	}
	f(format, data...)
}

// Errorf for errors with "Error" level
func (g *GoLog) Errorf(format string, data ...interface{}) {
	if g.minShowLevel > Error {
		return
	}
	g.output(g.makeData().Errorf, format, data...)
}

// Panicf for panic errors
func (g *GoLog) Panicf(format string, data ...interface{}) {
	if g.minShowLevel > Panic {
		return
	}
	g.output(g.makeData().Panicf, format, data...)
}

// Warningf for warning errors
func (g *GoLog) Warningf(format string, data ...interface{}) {
	if g.minShowLevel > Warning {
		return
	}
	g.output(g.makeData().Warningf, format, data...)
}

// tarcef for trace level output
func (g *GoLog) Tracef(format string, data ...interface{}) {
	if g.minShowLevel > Trace {
		return
	}
	g.output(g.makeData().Tracef, format, data...)
}


// WithField provides setting of fields to the response
func (g *GoLog) WithField(key string, value interface{}) *GoLog {
	g.fields[key] = value
	return g
}

// makeData provides making of additional things for logger
// like fields, errors. etc
func (g *GoLog) makeData() *logrus.Entry {
	entry := logrus.NewEntry(g.logger)
	if len(g.fields) > 0 {
		entry = entry.WithFields(g.fields)
		g.fields = logrus.Fields{}
	}
	return entry
}
