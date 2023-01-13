package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type BulkMessage struct {
	Index   string        `json:"Index"`
	Records []MailMessage `json:"records"`
}

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
	var arrayMessage []MailMessage
	arrayMessage = readDir(config.PathData, "", arrayMessage)
	fmt.Printf("End: %v\n", time.Now())

	if len(arrayMessage) > 0 {
		bulk := BulkMessage{
			Index:   zinc.Index,
			Records: arrayMessage,
		}
		jsonDataBytes, err := json.Marshal(bulk)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		res, err := zinc.LoadDataBulkV2(jsonDataBytes)
		if err != nil {
			log.Fatal(err)
		}
		_ = res
		fmt.Println(string(res))
	}

}

func readDir(path string, fileName string, arrayMessage []MailMessage) []MailMessage {
	//var arrayMessage []MailMessage
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		pathFile := path + "\\" + archivo.Name()
		if archivo.IsDir() {
			arrayMessage = readDir(pathFile, fileName+"\\"+archivo.Name(), arrayMessage)
		} else {
			mailMessage, err := NewMailMessageFromFile(pathFile, fileName+"_"+archivo.Name())
			if err != nil {
				log.Println(err)
			} else {
				arrayMessage = append(arrayMessage, *mailMessage)
			}
		}
	}
	return arrayMessage
	/*if len(arrayMessage) > 0 {
		bulk := BulkMessage{
			Index:   zinc.Index,
			Records: arrayMessage,
		}
		jsonDataBytes, err := json.Marshal(bulk)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		res, err := zinc.LoadDataBulkV2(jsonDataBytes)
		if err != nil {
			log.Fatal(err)
		}
		_ = res
		//fmt.Println(string(res))
	}*/
}
