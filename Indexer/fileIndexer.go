package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	//fmt.Printf("Start: %v\n", time.Now())
	readDir(config.PathData, "")
	//fmt.Printf("End: %v\n", time.Now())

}

func readDir(path string, fileName string) {
	var arrayMessage []MailMessage
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		pathFile := path + "\\" + archivo.Name()
		if archivo.IsDir() {
			readDir(pathFile, fileName+"\\"+archivo.Name())
		} else {
			mailMessage, err := NewMailMessageFromFile(pathFile, fileName+"_"+archivo.Name())
			if err != nil {
				log.Println(err)
			} else {
				arrayMessage = append(arrayMessage, *mailMessage)
			}
		}
	}
	if len(arrayMessage) > 0 {
		bulk := BulkMessage{
			Index:   zinc.Index,
			Records: arrayMessage,
		}
		jsonDataBytes, err := json.Marshal(bulk)
		if err != nil {
			log.Fatal(err)
		}
		res, err := zinc.LoadDataBulkV2(jsonDataBytes)
		if err != nil {
			log.Fatal(err)
		}
		_ = res
		//fmt.Println(string(res))
	}
}
