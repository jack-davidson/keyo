package main

type Entry struct {
	Name     string `json:"name"`
	Len      int    `json:"len"`
	Seed     []byte `json:"seed"`
	Host     string `json:"host"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
