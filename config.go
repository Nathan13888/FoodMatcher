package main

import (
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("foods", map[string]string{
		"fried chicken": string(MAIN_COURSE),
		"mango pudding": string(DESSERT),
	})
}

func setupConfig() {
	setDefaults()

	viper.SetConfigName("data")      // name of config file (without extension)
	viper.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.fm") // call multiple times to add many search paths
	viper.AddConfigPath(".")         // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			viper.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	typesOfFoods = viper.GetStringMapString("foods")
}
