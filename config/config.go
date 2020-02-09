package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

const configPath = "config/yig-collector.toml"
type Config struct {
	LogPath                string        `toml:"log_path"`
	BucketLogPath          string        `toml:"bucket_log_path"`
	TidbInfo               string        `toml:"tidb_info"`
	DbMaxIdleConns         int           `toml:"db_max_open_conns"`
	DbMaxOpenConns         int           `toml:"db_max_idle_conns"`
	DbConnMaxLifeSeconds   int           `toml:"db_conn_max_life_seconds"`
	Producer               DummyProducer `toml:"producer"`
}

type DummyProducer struct {
	EndPoint  string
	AccessKey string
	SecretKey string
}
var Conf Config
func ReadConfig() {
	data, err := ioutil.ReadFile(configPath)  //读取配置文件
	if err != nil {    //两个err嵌套
		if err != nil {
			panic("[ERROR] Cannot open /etc/yig/yig-collector.toml")
		}
	}
	fmt.Println(string(data))
	_, err = toml.Decode(string(data), &Conf)  //解析配置文件
	if err != nil {
		panic("[ERROR] Load yig-collector.toml error: " + err.Error())
	}
}