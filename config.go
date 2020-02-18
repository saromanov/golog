package golog

import "github.com/sirupsen/logrus"

// Config defines configuration for logger
type Config struct {
	MinShowLevel Level
	Hooks        []logrus.Hook
}
