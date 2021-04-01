package model

type Machine struct {
	IP   string `json:"ip"`
	Use  [4]int `json:"use"`
	Load [4]int `json:"load"` // avg all load, avg cpu load, avg gpu load, avg mem load
}
