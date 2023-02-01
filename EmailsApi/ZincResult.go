package main

import (
	"encoding/json"
	"fmt"
)

type HitsTotalResult struct {
	Value int `json:"value"`
}

type SourceRowResult struct {
	Subject string   `json:"subject"`
	From    []string `json:"from"`
	To      []string `json:"to"`
	Body    string   `json:"body"`
}

type HitsRowResult struct {
	Id     string           `json:"_id"`
	Source *SourceRowResult `json:"_source"`
}

type HitsResult struct {
	Total *HitsTotalResult `json:"total"`
	Data  []*HitsRowResult `json:"hits"`
}

type ZincSearchResult struct {
	Records *HitsResult `json:"hits"`
}

func SearchZincResult(text string, from int, size int) (*ZincSearchResult, error) {
	var config *Config
	config, err := NewConfig("config,json")
	var zinc Zinc = Zinc{
		Server:   config.ZincServer,
		Index:    config.ZincIndex,
		User:     config.ZincUser,
		Password: config.ZincPassword,
	}

	datos, err := zinc.Search(text, from, size)
	if err != nil {
		return nil, err
	}

	r := ZincSearchResult{}
	err = json.Unmarshal(datos, &r)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", r)

	return &r, nil
}
