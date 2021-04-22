package model

type Json struct {
	Users    map[string]*User    `json:"users"`
	Machines map[string]*Machine `json:"machines"`
	IP       map[string]bool     `json:"IP"`
	Token    string              `json:"token"`
}
