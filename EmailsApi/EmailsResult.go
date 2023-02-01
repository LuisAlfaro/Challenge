package main

import (
	"encoding/json"
	"log"
)

type SearchResult struct {
	FromCount int          `json:"fromCount"`
	ToCount   int          `json:"toCount"`
	Total     int          `json:"total"`
	Size      int          `json:"size"`
	Data      []DataResult `json:"data"`
}

type DataResult struct {
	Id      string   `json:"id"`
	Subject string   `json:"subject"`
	From    []string `json:"from"`
	To      []string `json:"to"`
	Body    string   `json:"body"`
}

func Search(text string, from int, size int) ([]byte, error) {
	zincResult, err := SearchZincResult(text, from, size)
	if err != nil {
		return nil, err
	}
	_ = zincResult

	var dataResult = make([]DataResult, 0)

	for _, hit := range zincResult.Records.Data {
		data := DataResult{
			Id:      hit.Id,
			Subject: hit.Source.Subject,
			From:    hit.Source.From,
			To:      hit.Source.To,
			Body:    hit.Source.Body,
		}

		nuevoSlice := append(dataResult, data)
		dataResult = nuevoSlice
	}

	total := zincResult.Records.Total.Value
	fromCount := from + size

	result := SearchResult{}
	result.Total = total
	result.Data = dataResult
	result.FromCount = from + 1
	if fromCount >= total {
		result.ToCount = total
	} else {
		result.ToCount = from + size
	}
	result.Size = size

	datosJson, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	return datosJson, nil
}
