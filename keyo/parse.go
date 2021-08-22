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
		entry.Seed = core.Seed(1024)
		entries := LoadEntries()
		entries.New(entry)
		entries.Write()
	case "clear":
		entries := LoadEntries()
		entries.Clear()
		entries.Write()
	default:
		fmt.Printf("Password: ")
		complement, _ := term.ReadPassword(1)
		fmt.Printf("\n")
		entries := LoadEntries()
		for _, e := range entries.Entries {
			if e.Host == a[1] || e.Host == a[2] {
				if e.Email == a[2] || e.Email == a[1] {
					fmt.Println(core.Encode(core.Derive(complement, e.Seed)))
				}
			}
		}
	}
}
