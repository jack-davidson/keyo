package main

import (
	"fmt"
	"os"
)

var KeyFilePath = os.Getenv("HOME") + "/.keyo"

type KeyData struct {
	Seed     []byte `json:"seed"`
	Host     string `json:"host"`
	Email    string `json:"email"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

func main() {
	for i, arg := range os.Args {
		switch os.Args[1] {
		case "new":
			fmt.Printf("%s, %d", arg, i)
		}
	}
	/*
		b, err := os.ReadFile(KeyFilePath)
		if errors.Is(err, os.ErrNotExist) {
			os.WriteFile(KeyFilePath, []byte("{}"), 0667)
		}
		var keyData []KeyData
		json.Unmarshal(b, &keyData)
		fmt.Println(keyData)
		os.WriteFile(KeyFilePath, b, 0667)
	*/
}
