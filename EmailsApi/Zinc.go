package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Zinc struct {
	Server   string
	Index    string
	User     string
	Password string
}

func getQuery(text string) string {
	if text != "" {
		return fmt.Sprintf(`
			"search_type": "term",			
			"query":
			{
				"term": "%s"
			},`, text)
	} else {
		return ""
	}
}

func (z *Zinc) Search(text string, from int, size int) ([]byte, error) {
	query := getQuery(text)
	bodyQuery := fmt.Sprintf(`{        
        %s
        "from": %d,
        "max_results": %d,
        "_source": [
			"subject",
        	"to",
        	"from",
        	"date",        	
        	"body"
		]}`, query, from, size)

	client := &http.Client{}
	zinc_url := z.Server + "/api/" + z.Index + "/_search"
	request, err := http.NewRequest("POST", zinc_url, strings.NewReader(bodyQuery))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(z.User, z.Password)
	request.Header.Add("Content-type", "application/json")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

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
