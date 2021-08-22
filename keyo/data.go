package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var DataFilePath = os.Getenv("HOME") + "/.keyo"

type Entries struct {
	Entries []Entry
}

func LoadEntries() *Entries {
	b, err := os.ReadFile(DataFilePath)
	if errors.Is(err, os.ErrNotExist) {
		os.Create(DataFilePath)
	}
	entries := []Entry{}
	json.Unmarshal(b, &entries)
	return &Entries{entries}
}

func (e *Entries) New(newEntry Entry) {
	for _, entry := range e.Entries {
		if entry.Email == newEntry.Email &&
			entry.Host == newEntry.Host &&
			entry.Username == newEntry.Username &&
			entry.Name == newEntry.Name {
			fmt.Println("Cannot create duplicate entry")
			return
		}
		if entry.Name == newEntry.Name {
			fmt.Println("Entry must have unique Name")
		}
	}
	e.Entries = append(e.Entries, newEntry)
}

func (e *Entries) Clear() {
	e.Entries = nil
}

func (e *Entries) Write() error {
	b, err := json.MarshalIndent(e.Entries, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(DataFilePath, b, 0667)
}
