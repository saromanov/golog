package main

import (
	"os"

	"github.com/saromanov/golog"
	"github.com/sirupsen/logrus"
)

func main() {
	g := golog.New(&golog.Config{
		MinShowLevel: golog.Debug,
	})
	g.Before(func(l *logrus.Logger) {
		l.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		l.SetOutput(os.Stdout)
	})
	g.Infof("test")
	g.WithField("key", "value").Infof("DATA")
}
