package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func FixSubject(datos string) (string, string, error) {
	beforeSubject, afterSubject, found := strings.Cut(string(datos), "Subject:")
	if found {
		before, after, found := strings.Cut(afterSubject, "Mime-Version:")
		if found {
			var buffer bytes.Buffer
			buffer.WriteString(beforeSubject + "Subject:" + "\n" + "Mime-Version:" + after)
			newDatos := buffer.String()
			return newDatos, before, nil
		}
	}
	err := fmt.Errorf("No se encontro el Subject")
	return "", "", err
}

func NewMailMessageFromFile(path string, fileName string) (*MailMessage, error) {
	datosComoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := &MailMessage{
		FileName: fileName,
	}
	var msg *mail.Message
	msg, err = mail.ReadMessage(bytes.NewBuffer(datosComoBytes))
	if err != nil {
		newDatos, subject, err := FixSubject(string(datosComoBytes))
		if err != nil {
			return nil, err
		}
		s := strings.NewReader(newDatos)
		msg, err := mail.ReadMessage(s)
		if err != nil {
			return nil, err
		}
		m.LoadFromMailMessage(msg)
		m.Subject = subject
		return m, nil
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

func (m *MailMessage) ToJson() (string, error) {
	datajson, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(datajson), nil
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(stream)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
