package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
	Id struct{
		Username string `json:"username"`
		Password string `json:"password"`
	} //`json:"id"`
}

func LoadConfig(f string) (Config, error){
	var config Config
	cFile, _ := os.Open(f)
	defer cFile.Close()

	jsonPaser := json.NewDecoder(cFile)
	err := jsonPaser.Decode(&config)
	return config, err
}

func main(){
	fmt.Println("Loading Config file config.json ....")
	config,_ := LoadConfig("config.json")
	fmt.Println("User =",config.Id)
}