package config

import (
	"encoding/json"
	"log"
	"os"
)

var Configs Config

type Config struct {
	App      app  `json:"app"`
	Job      jobs `json:"jobs"`
	ProdMode bool `json:"prod"`
}

type jobs struct {
	Notifier uint64 `json:"notifierTime"`
}

type app struct {
	ServerName string `json:"host"`
	Port       int64  `json:"port"`
	Logfile    string `json:"logfile"`
}

func Setup(F string) {
	byteValue, err := os.ReadFile(F)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	err = json.Unmarshal(byteValue, &Configs)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	log.Printf("cdonfig: %+v", Configs)
	// for logging
	//file, err := os.OpenFile(Configs.App.Logfile, os.O_WRONLY, 0777)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.SetOutput(file)
}
