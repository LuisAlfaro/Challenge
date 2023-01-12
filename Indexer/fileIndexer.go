package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var config Config = Config{}
var zinc Zinc = Zinc{}

func main() {
	config, err := NewConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	zinc = Zinc{
		Server:   config.ZincServer,
		Index:    config.ZincIndex,
		User:     config.ZincUser,
		Password: config.ZincPassword,
	}
	fmt.Printf("Start: %v\n", time.Now())
	readDir(config.PathData, "")
	fmt.Printf("End: %v\n", time.Now())
}

func readDir(path string, fileName string) {
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		pathFile := path + "\\" + archivo.Name()
		if archivo.IsDir() {
			readDir(pathFile, fileName+"\\"+archivo.Name())
		} else {
			mailMessage, err := NewMailMessageFromFile(pathFile, archivo.Name())
			if err != nil {
				log.Println(err)
			} else {
				dataJson, err := mailMessage.ToJson()
				if err != nil {
					log.Fatal(err)
				}
				res, err := zinc.LoadDataDoc(dataJson)
				if err != nil {
					log.Fatal(err)
				}
				_ = res
			}
		}
	}
}
