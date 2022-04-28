package config

import "time"

type DB struct{
	Host string
	Port int
	ConnectionTimeoutInSeconds time.Duration
}