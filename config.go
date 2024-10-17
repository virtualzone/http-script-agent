package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Port     int
	Command  string
	Username string
	Password string
}

var _configInstance *Config
var _configOnce sync.Once

func GetConfig() *Config {
	_configOnce.Do(func() {
		_configInstance = &Config{}
		_configInstance.ReadConfig()
	})
	return _configInstance
}

func (c *Config) ReadConfig() {
	port, err := strconv.Atoi(c.getEnv("PORT", "8080"))
	if err != nil {
		log.Panicln("PORT must be numeric")
	}
	c.Port = port
	c.Command = c.getEnv("SHELL_COMMAND", "/app/demo.sh")
	c.Username = c.getEnv("USERNAME", "")
	c.Password = c.getEnv("PASSWORD", "")
}

func (c *Config) getEnv(key, defaultValue string) string {
	res := os.Getenv(key)
	if res == "" {
		return defaultValue
	}
	return res
}
