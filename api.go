package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func getNote(token *string, password *string) *Note {
	data := url.Values{}
	data.Set("password", *password)
	bs := bytes.NewBufferString(data.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.noteshred.com/v1/notes/%s", *token), bs)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Failed To Retrieve Note: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var note = new(Note)
	err = json.Unmarshal(body, &note)
	if err != nil {
		fmt.Println("Marshalling Error:", err)
	}

	return note
}

func newNote(content *string, password *string, email *string) *Note {
	data := url.Values{}
	data.Set("title", "Command Line Note")
	data.Set("shred_method", "1")
	data.Set("password", *password)
	data.Set("content", *content)
	data.Set("recipients", *email)
	bs := bytes.NewBufferString(data.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.noteshred.com/v1/notes"), bs)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Failed To Create Note: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var note = new(Note)
	err = json.Unmarshal(body, &note)
	if err != nil {
		fmt.Println("Marshalling Error:", err)
	}

	//fmt.Println(fmt.Sprintf("%s", body))
	return note
}
