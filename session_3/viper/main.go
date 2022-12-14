package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// set defaults
	viper.SetDefault("test_config", "default")

	// get config from files
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; bind from environment variable instead
			viper.BindEnv("test_config")
			// this will take TEST_CONFIG environment variable value
		} else {
			panic(err)
		}
	}

	// get the config
	fmt.Println(viper.GetString("test_config"))
}
