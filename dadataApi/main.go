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
	str := "7707083893"
	respBody, err := SimpleReqDadata(str)
	if err != nil {
		panic(err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}
	data := dat["suggestions"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})
	value := dat["suggestions"].([]interface{})[0].(map[string]interface{})["value"]
	kpp := data["kpp"]
	ogrn := data["ogrn"]
	name := data["management"].(map[string]interface{})["name"]

	fmt.Println("\n\n\n")
	fmt.Println("Value: ", value)
	fmt.Println("KPP: ", kpp)
	fmt.Println("OGRN: ", ogrn)
	fmt.Println("Name: ", name)
}

func SimpleReqDadata(query string) (respBody []byte, err error) {
	reqBody, _ := json.Marshal(map[string]string{
		"query":       query,
		"branch_type": "MAIN",
	})

	req, errReq := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if errReq != nil {
		fmt.Println("Error Request")
		//panic(errReq)
		return []byte{}, errReq
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 84eeb31c0764a70db046f2f2b0641eb07bad71d2")

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		fmt.Println("Error Response")
		//panic(errResp)
		return []byte{}, errResp
	}

	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println("Error Read")
		//panic(errRead)
		return []byte{}, errRead
	}
	fmt.Println(string(body) + "\n\n\n")

	return body, nil

}
