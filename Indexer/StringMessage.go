package main

import (
	"bufio"
	"encoding/json"
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
	Header
}

type Header struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	XFrom                   string
	XTo                     string
	Xcc                     string
	Xbcc                    string
	XFolder                 string
	XOrigin                 string
	XFileName               string
}

func NewStringMessageFromFile(path string, fileName string) (*StringMessage, error) {
	sm := &StringMessage{
		FileName: fileName,
	}
	sl, err := sm.FileToLines(path)
	if err != nil {
		return nil, err
	}
	_ = sl

	return sm, nil
}

/*func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}*/

func (s *StringMessage) FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s.AssignToProperty(line)
		lines = append(lines, line)
	}
	err = scanner.Err()
	return
}

func (s *StringMessage) AssignToProperty(line string) {
	before, after, found := strings.Cut(line, ":")
	after = strings.TrimSpace(after)
	if found {
		switch before {
		case "Message-ID":
			{
				s.Header.MessageID = after
			}
		case "Date":
			{
				s.Header.Date = after
			}
		case "From":
			{
				s.Header.From = after
			}
		case "To":
			{
				s.Header.To = after
			}
		case "Subject":
			{
				s.Header.Subject = after
			}
		case "Mime-Version":
			{
				s.Header.MimeVersion = after
			}
		case "Content-Type":
			{
				s.Header.ContentType = after
			}
		case "Content-Transfer-Encoding":
			{
				s.Header.ContentTransferEncoding = after
			}
		case "X-From":
			{
				s.Header.XFrom = after
			}
		case "X-To":
			{
				s.Header.XTo = after
			}
		case "X-cc":
			{
				s.Header.Xcc = after
			}
		case "X-bcc":
			{
				s.Header.Xbcc = after
			}
		case "X-Folder":
			{
				s.Header.XFolder = after
			}
		case "X-Origin":
			{
				s.Header.XOrigin = after
			}
		case "X-FileName":
			{
				s.Header.XFileName = after
			}
		}
	}

}

func (m *StringMessage) ToJson() ([]byte, error) {
	datajson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return datajson, nil
}
