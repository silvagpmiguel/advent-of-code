package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func get(url string, cookie string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("getting new request: %v", err)
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

// args: year day cookie
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Error, wrong arguments.\nUsage: go run script.go \"year\" \"day\" \"cookie\" ")
		return
	}
	year := os.Args[1]
	day := os.Args[2]
	cookie := os.Args[3]

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	fmt.Println(url)

	input, err := get(url, cookie)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(255)
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
		os.Exit(2)
	}

	fmt.Println("Created AoC input file in " + path)
}
