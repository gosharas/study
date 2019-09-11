package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/findById/party"

type Dates struct {
	Value string
	Kpp   string
	Ogrn  string
	Name  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getByInn/{inn}", GetByInn).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func GetByInn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	respBody, err := simpleReqDadata(params["inn"])
	if err != nil {
		panic(err)
	}
	dates := parsingJson(respBody)
	json.NewEncoder(w).Encode(dates)
}

func parsingJson(body []byte) *Dates {
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error")
		panic(err)
	}
	data := dat["suggestions"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})
	value := dat["suggestions"].([]interface{})[0].(map[string]interface{})["value"].(string)
	kpp := data["kpp"].(string)
	ogrn := data["ogrn"].(string)
	name := data["management"].(map[string]interface{})["name"].(string)

	dates := &Dates{value, kpp, ogrn, name}

	return dates
}

func simpleReqDadata(query string) (respBody []byte, err error) {
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

	return body, nil

}
