package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func get(endpoint string, cookie string) ([]byte, error) {
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

func post(endpoint string, cookie string, level string, answer string) (string, error) {
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

	if err != nil {
		return "", fmt.Errorf("performing request: %v", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("reading body: %v", err)
	}

	body := string(bytes)
	var result string
	if strings.Contains(body, "That's the right answer") {
		result = "Correct, you nailed it :)"
	} else if strings.Contains(body, "Did you already complete it") {
		result = "Oops, you've already sent that before :("
	} else if strings.Contains(body, "That's not the right answer") {
		result = "Oops, that's incorrect :("
	}
	return result, nil
}

// args: input|answer cookie year day level answer
func main() {
	if len(os.Args) != 5 && len(os.Args) != 7 {
		fmt.Println("Error, wrong arguments.\nUsage: go run script.go \"input|answer\" \"cookie\" \"year\" \"day\" [level] [answer]")
		return
	}

	var err error
	var input []byte
	action := os.Args[1]
	cookie := os.Args[2]
	year := os.Args[3]
	day := os.Args[4]
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/%s", year, day, action)

	switch action {
	case "answer":
		level := os.Args[5]
		answer := os.Args[6]
		result, err := post(url, cookie, level, answer)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
		return
	case "input":
		input, err = get(url, cookie)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Error: Wrong action.\nUsage: go run script.go \"input|answer\" \"cookie\" \"year\" \"day\" [level] [answer]")
		os.Exit(1)
	}

	path := fmt.Sprintf("%s/input/%s.in", year, day)

	err = os.MkdirAll(filepath.Dir(path), 0755)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(path, input, 0644)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Created AoC input file in " + path)
}
