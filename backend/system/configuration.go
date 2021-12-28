package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration_t struct {
	AdminPassword string `json:"admin_password"`
	Port          int    `json:"port"`
	SecretKey     string `json:"secret_key"`
}

var conf Configuration_t

func LoadConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteContent, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteContent, &conf)

	fmt.Println("Configuration loaded successfully.")
}

func GetConfiguration() Configuration_t {
	return conf
}
