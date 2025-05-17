package core

import "log"

type Core struct {
	Port     int
	Hostname string
	Logger   *log.Logger
}
