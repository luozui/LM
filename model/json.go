package model

type Json struct {
	Users    map[string]*User    `json:"users"`
	Machines map[string]*Machine `json:"machines"`
	Token    string              `json:"token"`
}
