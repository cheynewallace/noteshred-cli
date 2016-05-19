package main

import (
	"fmt"
	"os"
	"strings"
)

var action, token, password, apiKey, message, email string

func init() {
	apiKey = strings.ToLower(os.Getenv("NOTESHRED_API_KEY"))
	if apiKey == "" {
		fmt.Println("NOTESHRED_API_KEY not set")
		fmt.Println("")
		os.Exit(0)
	}

	if len(os.Args) < 3 {
		showInstructions()
		os.Exit(0)
	}

	// show or new
	action = strings.ToLower(os.Args[1])

	if action == "show" {
		if len(os.Args) < 4 {
			showInstructions()
			os.Exit(0)
		}
		token = os.Args[2]
		password = os.Args[3]

	} else if action == "new" {
		message = os.Args[2]
		if len(os.Args) == 4 && os.Args[3] != "" {
			email = os.Args[3]
		}
	}
}

func main() {
	if action == "show" {
		note := getNote(&token, &password)
		if note.Message != "" {
			fmt.Println(note.Message)
		} else {
			fmt.Println("############################################")
			fmt.Println(fmt.Sprintf("Title:   %s", note.Title))
			fmt.Println(fmt.Sprintf("From:    %s", note.CreatedBy))
			fmt.Println(fmt.Sprintf("Message: %s", note.Content))
			fmt.Println("############################################")
		}
	}

	if action == "new" {
		uPass := randStr(16, "mixed")
		note := newNote(&message, &uPass, &email)
		if note.Message != "" {
			fmt.Println(note.Message)
		} else {
			fmt.Println("############################################")
			fmt.Println("Note Successfully Created!")
			fmt.Println(fmt.Sprintf("Token:    %s", note.Token))
			fmt.Println(fmt.Sprintf("Password: %s", uPass))
			fmt.Println(fmt.Sprintf("URL:      shred.io/%s", note.Token))
			fmt.Println(fmt.Sprintf("CLI:      noteshred show %s %s", note.Token, uPass))
			if email != "" {
				fmt.Println(fmt.Sprintf("Notified: %s", email))
			}
			fmt.Println("############################################")
		}
	}
}

func showInstructions() {
	fmt.Println("############################################")
	fmt.Println("NoteShred - Auto Shredding, Encrypted Notes")
	fmt.Println("############################################")
	fmt.Println("Usage:")
	fmt.Println("noteshred new \"secret message here\"")
	fmt.Println("noteshred show <note token> <password>")
	fmt.Println("############################################")
}
