package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	TOPSTORIES  = "top-headlines"
	EVERYTHINNG = "everything"
	HOST        = "newsapi.org"
)

type Source struct {
	ID          string `json:"status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Response struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type NewsParams struct {
	Q          string
	Sources    string
	Category   string
	Language   string
	Country    string
	Domains    string
	From_param string
	Sort_by    string
	Page       string
}

func (n *NewsParams) GetTopHeadlines() (*Response, error) {

	resp, err := n.MakeRquest(TOPSTORIES)

	if err != nil {
		return nil, err
	}

	r := Response{}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (n *NewsParams) GetEverything() {
	resp, err := n.MakeRquest(EVERYTHINNG)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func (n *NewsParams) MakeRquest(endpoint string) (*http.Response, error) {
	requestUrl, err := n.GenerateRequestURL(endpoint)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NewsParams) GenerateRequestURL(endpoint string) (string, error) {

	jb, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	params := make(map[string]string)
	_ = json.Unmarshal(jb, &params)

	p := &url.Values{}
	for k, v := range params {
		if len(strings.TrimSpace(v)) > 0 {
			p.Add(strings.ToLower(k), v)
		}

	}

	p.Add("apiKey", os.Getenv("NEWSAPI"))

	url := url.URL{
		Scheme:   "https",
		Host:     HOST,
		Path:     fmt.Sprintf("v2/%s", endpoint),
		RawQuery: p.Encode(),
	}

	fmt.Println(url.String())

	return url.String(), nil
}
