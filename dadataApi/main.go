package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/findById/party"

func main() {
	SimpleReqDadata()
}

func SimpleReqDadata() {
	reqBody, _ := json.Marshal(map[string]string{
		"query":       "7707083893",
		"branch_type": "MAIN",
	})
	req, errReq := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if errReq != nil {
		fmt.Println("Error Request")
		panic(errReq)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 84eeb31c0764a70db046f2f2b0641eb07bad71d2")
	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		fmt.Println("Error Response")
		panic(errResp)
		return
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println("Error Read")
		panic(errRead)
		return
	}
	fmt.Println("response Body:", string(body))

}
