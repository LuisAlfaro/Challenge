package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
	"time"
)

type MailMessage struct {
	FileName   string    `json:"fileName"`
	Message_Id string    `json:"message_id"`
	Date       time.Time `json:"date"`
	To         []string  `json:"to"`
	From       []string  `json:"from"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
}

func NewMailMessageFromFile(path string, fileName string) (*MailMessage, error) {
	datosComoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	msg, err := mail.ReadMessage(bytes.NewBuffer(datosComoBytes))
	if err != nil {
		return nil, err
	}
	m := &MailMessage{
		FileName: fileName,
	}
	m.LoadFromMailMessage(msg)
	return m, nil
}

func (m *MailMessage) LoadFromMailMessage(mailMessage *mail.Message) {
	m.Message_Id = mailMessage.Header.Get("Message-ID")
	date, err := mail.ParseDate(mailMessage.Header.Get("Date"))
	if err != nil {
		log.Fatal(err)
	}
	m.Date = date
	m.Subject = mailMessage.Header.Get("Subject")
	to := mailMessage.Header.Get("To")
	m.To = strings.Split(to, ",")
	from := mailMessage.Header.Get("From")
	m.From = strings.Split(from, ",")

	m.Body = StreamToString(mailMessage.Body)
}

func (m *MailMessage) ToJson() ([]byte, error) {
	datajson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return datajson, nil
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
