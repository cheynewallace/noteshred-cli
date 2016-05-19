package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
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
	t := time.Now()
	title := fmt.Sprintf("%s %v Note", t.Month(), t.Day())
	noteMarshalled, err := json.Marshal(Note{
		Title:       title,
		Password:    *password,
		Content:     *content,
		ShredMethod: 1,
		Recipients:  []string{*email}})

	bs := bytes.NewBuffer(noteMarshalled)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.noteshred.com/v1/notes"), bs)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	req.Header.Set("Content-Type", "application/json")
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
