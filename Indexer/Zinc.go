package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Zinc struct {
	Server   string
	Index    string
	User     string
	Password string
}

func (z *Zinc) LoadDataDoc(data []byte) ([]byte, error) {
	client := &http.Client{}
	zinc_url := z.Server + "/api/" + z.Index + "/_doc"
	request, err := http.NewRequest("POST", zinc_url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(z.User, z.Password)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (z *Zinc) LoadDataMulti(data []byte) ([]byte, error) {
	client := &http.Client{}
	zinc_url := z.Server + "/api/" + z.Index + "/_multi"
	request, err := http.NewRequest("POST", zinc_url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth(z.User, z.Password)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
