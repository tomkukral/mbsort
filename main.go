package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("mailsort")
	viper.AddConfigPath("$HOME/.mutt")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to read config file or defaults: %s", err)
	}

	viper.SetDefault("file.input", "$HOME/.mutt/mailboxes_raw")
	viper.SetDefault("file.output", "$HOME/.mutt/mailboxes")
	viper.SetDefault("metrics", []string{})
	viper.SetDefault("default_metric", 1000)
	viper.SetDefault("debug", false)

	var (
		debug   bool     = viper.GetBool("debug")
		metrics []string = viper.GetStringSlice("metrics")
		//		defMetric int      = viper.GetInt("metrics")
	)

	// print debug
	if debug {
		viper.Debug()
	}

}
