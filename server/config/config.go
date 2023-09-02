package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	Cfg configStruct
)

type configStruct struct {
	Zookeeper struct {
		Schema   string   `yaml:"schema"`
		ZkAddr   []string `yaml:"address"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"zookeeper"`

	Mysql struct {
		Address       []string `yaml:"address"`
		Username      string   `yaml:"username"`
		Password      string   `yaml:"password"`
		Database      string   `yaml:"database"`
		MaxOpenConn   int      `yaml:"maxOpenConn"`
		MaxIdleConn   int      `yaml:"maxIdleConn"`
		MaxLifeTime   int      `yaml:"maxLifeTime"`
		LogLevel      int      `yaml:"logLevel"`
		SlowThreshold int      `yaml:"slowThreshold"`
	} `yaml:"mysql"`

	Mongo struct {
		Uri         string   `yaml:"uri"`
		Address     []string `yaml:"address"`
		Database    string   `yaml:"database"`
		Username    string   `yaml:"username"`
		Password    string   `yaml:"password"`
		MaxPoolSize int      `yaml:"maxPoolSize"`
	} `yaml:"mongo"`

	Redis struct {
		Address  []string `yaml:"address"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"redis"`

	Kafka struct {
		Username         string   `yaml:"username"`
		Password         string   `yaml:"password"`
		Addr             []string `yaml:"addr"`
		LatestMsgToRedis struct {
			Topic string `yaml:"topic"`
		} `yaml:"latestMsgToRedis"`
		MsgToMongo struct {
			Topic string `yaml:"topic"`
		} `yaml:"offlineMsgToMongo"`
		MsgToPush struct {
			Topic string `yaml:"topic"`
		} `yaml:"msgToPush"`
		ConsumerGroupID struct {
			MsgToRedis string `yaml:"msgToRedis"`
			MsgToMongo string `yaml:"msgToMongo"`
			MsgToMySql string `yaml:"msgToMySql"`
			MsgToPush  string `yaml:"msgToPush"`
		} `yaml:"consumerGroupID"`
	} `yaml:"kafka"`
	APIServerGrpcPort int `yaml:"APIServerGrpcPort"`
	APIServerWsPort   int `yaml:"APIServerWsPort"`
}

func InitConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var cfg configStruct
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	Cfg = cfg
	return nil
}
