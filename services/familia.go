package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TokenizeResp struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

func TokenizeString(text string) []string {
	req, err := http.NewRequest("GET", os.Getenv("FAMILIA_URL")+"/tokenize", nil)
	if err != nil {
		log.Print(err)
	}

	q := req.URL.Query()
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := http.Post(req.URL.String(), "Application/json", nil)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var result TokenizeResp
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	return result.Data
}
