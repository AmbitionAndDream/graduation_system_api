package global

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Options struct {
	Log bool
}

type ServerConfig struct {
	EnvConfig *EnvConfig `toml:"env"`
	DBConfig  *DBConfig  `toml:"database"`
}

func (c *ServerConfig) toString() string {
	b, _ := json.Marshal(map[string]interface{}{"env": c.EnvConfig, "database": c.DBConfig})
	return string(b)
}

type EnvConfig struct {
	Env string
}

type DBConfig struct {
	Server     string
	Ports      []int
	PassWord   string
	Connection *Connection
}

type Connection struct {
	ConnectionMax int
	ConnectionMin int
	IdleTime      int
}

func must(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

var conf ServerConfig

func InitConfig(filePath string, options *Options) error {
	logrus.Infof("the config filepath is  %s", filePath)

	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		return err
	}
	logrus.Infof("sever config is %s", conf.toString())
	if options.Log {
		must(func() error {
			return initLog()
		})
	}
	return nil
}

func getConfig() *ServerConfig {
	return &conf
}

func initLog() error {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)
	return nil
}
