/*
# ------------------------------------------------------------
# -- conf.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package global

import (
	"fmt"
	"omciAnalyzer/utils"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     string        `yaml:"httpPort"`
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type MinIOConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Secure          bool   `yaml:"secure"`
	Bucket          string `yaml:"bucket"`
}

type MysqlConfig struct {
	User        string        `yaml:"user"`
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	DefaultDB   string        `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}

type KafkaConfig struct {
	Addr                   string `yaml:"addr"`
	TopicCollectorRequest  string `yaml:"topicCollectorRequest"`
	TopicCollectorResponse string `yaml:"topicCollectorResponse"`
	GroupCollectorRequest  string `yaml:"groupCollectorRequest"`
	GroupCollectorResponse string `yaml:"groupCollectorResponse"`
}

type Config struct {
	Server ServerConfig
	MinIO  MinIOConfig
	Mysql  MysqlConfig
	Kafka  KafkaConfig
}

var (
	AppConf Config
)

func InitConf(fileName string) error {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		utils.Log("open file failed" + fmt.Sprintf("%s", err))
		return err
	}

	err = yaml.UnmarshalStrict(yamlFile, &AppConf)
	if err != nil {
		utils.Log("get AppConf failed" + fmt.Sprintf("%s", err))
		return err
	}

	return nil
}
