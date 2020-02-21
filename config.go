package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

//.Cache(config)
type Server struct {
	Host string `json:"host" mapstructure:"host"`
	Port uint32 `json:"port" mapstructure:"port"`
}
type TokenCfg struct {
	Version string `json:"version" mapstructure:"version"`
	Et      string `json:"et" mapstructure:"et"`
	Method  string `json:"method" mapstructure:"method"`
}
type OneNet struct {
	ProductId   string `json:"Product-id" mapstructure:"Product-id"`
	Server      Server `json:"Server" mapstructure:"Server"`
	EquipName   string `json:"Device-name" mapstructure:"Device-name"`
	EquipId     string `json:"Device-id" mapstructure:"Device-id"`
	EquipKey    string `json:"Device-key" mapstructure:"Device-key"`
	KeepAlive   uint32 `json:"KeepAlive" mapstructure:"KeepAlive"`
	PingTimeout uint32 `json:"PingTimeout" mapstructure:"PingTimeout"`
	Token       TokenCfg `json:"Token" mapstructure:"Token"`
}
type Configuration struct {
	OneNet   OneNet `json:"OneNet" mapstructure:"OneNet"`
	LogLevel string `json:"log-level" mapstructure:"log-level"`
}

func AddConfig(name string) (c *viper.Viper) {
	c = viper.New()
	c.SetConfigName(name)
	viper.SetConfigType("json")
	c.AddConfigPath("/etc/config")
	c.AddConfigPath("/etc")
	c.AddConfigPath(".")
	c.ReadInConfig()
	c.WatchConfig()
	c.OnConfigChange(func(e fsnotify.Event) {
		//c.Viper.Unmarshal(s)
		fmt.Println("Config file changed:", e.Name)
	})
	return
}

func InitConfig() error {
	cfg := AddConfig("config")
	err := cfg.Unmarshal(&config)
	return err
}