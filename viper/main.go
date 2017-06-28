package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	setDefaults()
	setConfig()
	watchConfig()
	confEnv()
	setEnv()

	// get the default values
	devDir := viper.Get("devDir")
	fmt.Println("devDir:", devDir)

	// get environment variable
	fmt.Println("repo:", getEnv("repo"))
}

func setDefaults() {
	viper.SetDefault("devDir", "dev")
	viper.SetDefault("prodDir", "prod")
}

func setConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./dev-viper")
	err := viper.ReadInConfig()
	checkError(err)
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func confEnv() {
	viper.SetEnvPrefix("myapp")
	viper.BindEnv("repo")
	viper.BindEnv("type")
}

func setEnv() {
	os.Setenv("MYAPP_REPO", "https://github.com/ellerbrock")
	os.Setenv("MYAPP_TYPE", "git")
}

func getEnv(ev string) string {
	return viper.GetString(ev)
}

func checkError(err error) {
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
