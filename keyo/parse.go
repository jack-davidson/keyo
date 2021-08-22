package main

import (
	"fmt"
	"strings"

	"github.com/jack-davidson/keyo/core"
	"golang.org/x/term"
)

type Args []string

func parseEntry(args []string) Entry {
	entry := Entry{}
	for _, arg := range args {
		pair := strings.Split(arg, ":")
		keyArg := pair[0]
		valueArg := pair[1]
		switch keyArg {
		case "name":
			entry.Name = valueArg
		case "host":
			entry.Host = valueArg
		case "email":
			entry.Email = valueArg
		case "username":
			entry.Username = valueArg
		}
	}
	return entry
}

func (a Args) Parse() {
	switch a[1] {
	case "new":
		entry := parseEntry(a[2:])
		if entry.Name == "" {
			fmt.Println("Field 'name' required")
			return
		}
		entry.Seed = core.Seed(1024)
		entries := LoadEntries()
		entries.New(entry)
		entries.Write()
	case "clear":
		entries := LoadEntries()
		entries.Clear()
		entries.Write()
	case "list":
		entries := LoadEntries()
		for _, e := range entries.Entries {
			fmt.Printf("%s:\n\t%s\n\t%s\n\t%s\n", e.Name, e.Host, e.Email, e.Username)
		}
	default:
		fmt.Printf("Password: ")
		complement, _ := term.ReadPassword(1)
		fmt.Printf("\n")
		entries := LoadEntries()
		for _, e := range entries.Entries {
			if e.Name == a[1] {
				fmt.Println(core.Encode(core.Derive(complement, e.Seed)))
			}
		}
	}
}
