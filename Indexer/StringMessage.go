package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/mail"
	"os"
	"strings"
	"time"
)

type StringMessage struct {
	FileName   string    `json:"fileName"`
	Message_Id string    `json:"message_id"`
	Date       time.Time `json:"date"`
	To         []string  `json:"to"`
	From       []string  `json:"from"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
}

func NewStringMessageFromFile(path string, fileName string) (*MailMessage, error) {
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

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
