package config

import "log"

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
