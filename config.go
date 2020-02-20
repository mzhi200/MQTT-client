package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)
func AddConfig(name string) (c *viper.Viper) {
	c = viper.New()
	c.SetConfigName(name)
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
