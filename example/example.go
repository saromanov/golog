package main

import (
	"github.com/saromanov/golog"
)
func main(){
	g := golog.New()
	g.Before(nil)
	g.Infof("test")
}