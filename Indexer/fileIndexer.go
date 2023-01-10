package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

var pathDB string = "D:\\Projects\\Truora\\Proyecto\\Base Datos\\enron_mail_20110402\\enron_mail_20110402\\maildir\\baughman-d\\calendar"

var zinc_server string = "http://localhost:4080"
var zinc_uid string = "admin"
var zinc_pwd string = "Complexpass#123"

var index string = "EMails"

type eMail struct {
	FileName   string    `json:"fileName"`
	Message_Id string    `json:"message_id"`
	Date       time.Time `json:"date"`
	To         []string  `json:"to"`
	From       []string  `json:"from"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
}

func main() {

	fmt.Println(time.Now())
	readFiles(pathDB, "")
	fmt.Println(time.Now())
}

func readFiles(path string, fileName string) {
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		pathFile := path + "\\" + archivo.Name()
		if archivo.IsDir() {
			readFiles(pathFile, fileName+"\\"+archivo.Name())
		} else {
			readFile(pathFile, fileName+"\\"+archivo.Name())
		}

	}
}

func readFile(path string, fileName string) {
	datosComoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.NewReader(string(datosComoBytes))
	//msg, err := mail.ReadMessage(bytes.NewBuffer(datosComoBytes))
	msg, err := mail.ReadMessage(s)
	if err != nil {
		log.Println(err)
		fmt.Println(path)
		return
	}
	emailData(msg, fileName)
}

func emailData(mailMessage *mail.Message, fileName string) {
	id := mailMessage.Header.Get("Message-ID")
	date, err := mail.ParseDate(mailMessage.Header.Get("Date"))
	if err != nil {
		log.Fatal(err)
	}
	subject := mailMessage.Header.Get("Subject")
	to := mailMessage.Header.Get("To")
	toList := strings.Split(to, ",")
	from := mailMessage.Header.Get("From")
	fromList := strings.Split(from, ",")

	body := StreamToString(mailMessage.Body)
	data := eMail{
		FileName:   fileName,
		Message_Id: id,
		Date:       date,
		To:         toList,
		From:       fromList,
		Subject:    subject,
		Body:       body,
	}
	datajson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	_ = datajson
	loadData(datajson)
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func loadData(data []byte) {
	client := &http.Client{}
	zinc_url := zinc_server + "/api/" + index + "/_doc"
	request, err := http.NewRequest("POST", zinc_url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(zinc_uid, zinc_pwd)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = body
	//log.("%s\n", body)
	//fmt.Printf("%s\n", body)
}
