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
	readDir(config.PathData, "")
	fmt.Printf("End: %v\n", time.Now())
}

/*func LoadJson(jsonData string) {
	res, err := zinc.LoadDataMulti([]byte(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}*/

func LoadDataBulkV2(jsonData []byte) {
	res, err := zinc.LoadDataBulkV2(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	_ = res
	fmt.Println(string(res))
}

func LoadJsonDoc(jsonData []byte) {
	res, err := zinc.LoadDataDoc(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	_ = res
	//fmt.Println(string(res))
}

func LoadJsonMulti(jsonData []byte) {
	res, err := zinc.LoadDataMulti(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	_ = res
	//fmt.Println(string(res))
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
			fmt.Printf("Start: %v\n", time.Now())
			fmt.Println(fileName + "\\" + archivo.Name())
			readDir(pathFile, fileName+"\\"+archivo.Name())
		} else {
			mailMessage, err := NewMailMessageFromFile(pathFile, archivo.Name())
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
			fmt.Printf("Error: %s", err.Error())
		}
		LoadDataBulkV2(jsonDataBytes)
		/*fmt.Println(len(arrayMessage))
		for _, mMessage := range arrayMessage {
			jsonDataBytes, err := json.Marshal(mMessage)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			LoadJsonDoc(jsonDataBytes)
		}*/

		/*limit := 500
		for i := 0; i < len(arrayMessage); i += limit {
			batch := arrayMessage[i:min(i+limit, len(arrayMessage))]
			jsonDataBytes, err := json.Marshal(batch)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			LoadJsonMulti(jsonDataBytes)
		}*/

	}

	fmt.Printf("End: %v\n", time.Now())
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
