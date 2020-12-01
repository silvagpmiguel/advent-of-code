package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Get returns the input of a puzzle
func Get(endpoint string, cookie string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return nil, fmt.Errorf("creating request: %v", err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", cookie))
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("performing request: %v", err)
	}

	input, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("reading body: %v", err)
	}

	return input, nil
}

// Post sends the answer to AoC
func Post(endpoint string, cookie string, level string, answer string) (string, error) {
	client := &http.Client{}
	form := url.Values{
		"level":  {level},
		"answer": {answer},
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))

	if err != nil {
		return "", fmt.Errorf("creating request: %v", err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", cookie))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", fmt.Errorf("performing request: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing POST response as html: %v", err)
	}

	selection := doc.Find("body > :not(script):not(header):not(#sidebar)")

	return strings.TrimSpace(selection.Text()), nil
	return "", nil
}
