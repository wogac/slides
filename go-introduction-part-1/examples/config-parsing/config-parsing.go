package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	v "github.com/spf13/viper"
)

func setDefaultConfiguration() {
	v.SetConfigType("yaml")
	v.SetDefault("timeout", 500)
	v.SetDefault("host", "localhost")
	v.SetDefault("port", 8080)

	v.SetConfigName("config")
	v.AddConfigPath(".")
	// v.ReadInConfig()
}

func writeConfig() {
	v.WriteConfigAs(".config")
}

func extractConfigAndPrint() {
	fmt.Printf("Timeout: %d s, Host: %s:%d\n",
		v.Get("timeout"),
		v.Get("host"),
		v.Get("port"))
}

func main() {
	// Initialize configuration
	setDefaultConfiguration()
	// Show defaults
	extractConfigAndPrint()

	// Watching config for changes
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Configuration change detected:", e.Name)
		fmt.Println("Redisplaying configuration summary...")
		extractConfigAndPrint()
	})

	for {
		time.Sleep(time.Second * 5)
		fmt.Println("Still running...")
	}
}
