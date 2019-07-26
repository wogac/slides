package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func setDefaultConfiguration() {
	viper.SetConfigType("yaml")
	viper.SetDefault("timeout", 500)
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 8080)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// viper.ReadInConfig()
}

func writeConfig() {
	viper.WriteConfigAs(".config")
}

func extractConfigAndPrint() {
	fmt.Printf("Timeout: %d s, Host: %s:%d\n",
		viper.Get("timeout"),
		viper.Get("host"),
		viper.Get("port"))
}

func main() {
	// Initialize configuration
	setDefaultConfiguration()
	// Show defaults
	extractConfigAndPrint()

	// Watching config for changes
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Configuration change detected:", e.Name)
		fmt.Println("Redisplaying configuration summary...")
		extractConfigAndPrint()
	})

	for {
		time.Sleep(time.Second * 5)
		fmt.Println("Still running...")
	}
}
