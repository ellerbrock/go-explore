package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {

	// crypt get -plaintext -backend consul -endpoint 127.0.0.1:8500 go/info
	viper.AddRemoteProvider("consul", "127.0.0.1:8500", "crypt/mydata.json")
	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()

	if err != nil {
		log.Fatal(err)
	}

	name := viper.GetString("name")
	github := viper.GetString("github")
	fmt.Println("Name: ", name, "|  GitHub:", github)
}
