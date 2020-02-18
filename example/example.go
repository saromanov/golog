package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/saromanov/golog"
)
func main(){
	g := golog.New()
	g.Before(func(l *logrus.Logger){
		l.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		l.SetOutput(os.Stdout)
	})
	g.Infof("test")
}