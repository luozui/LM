package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/luozui/LM/model"
)

var Data model.Json

const (
	CONFIG = "./conf.json"
)

func Init() {
	Data = model.Json{
		Users:    make(map[string]*model.User),
		Machines: make(map[string]*model.Machine),
		Token:    "",
	}
	data, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		fmt.Print(err)
	}
	// unmarshall it
	err = json.Unmarshal(data, &Data)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func Write() error {
	file, err := json.Marshal(&Data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(CONFIG, file, 0644)
}

func init() {
	_, err := ioutil.ReadFile(CONFIG)
	if err == nil {
		return
	}
	Data = model.Json{
		Users:    make(map[string]*model.User),
		Machines: make(map[string]*model.Machine),
		IP:       make(map[string]bool),
		Token:    "101111",
	}
	Data.Machines["172.22.103.6"] = &model.Machine{
		IP:   "172.22.103.6",
		Use:  [4]int{0, 0, 0, 0},
		Load: [4]int{0, 0, 0, 0},
	}
	Data.IP = map[string]bool{
		"172.22.103.50":  true,
		"172.22.103.51":  true,
		"172.22.103.52":  true,
		"172.22.103.53":  true,
		"172.22.103.54":  true,
		"172.22.103.55":  true,
		"172.22.103.56":  true,
		"172.22.103.57":  true,
		"172.22.103.58":  true,
		"172.22.103.59":  true,
		"172.22.103.60":  true,
		"172.22.103.61":  true,
		"172.22.103.62":  true,
		"172.22.103.63":  true,
		"172.22.103.64":  true,
		"172.22.103.65":  true,
		"172.22.103.66":  true,
		"172.22.103.67":  true,
		"172.22.103.68":  true,
		"172.22.103.69":  true,
		"172.22.103.70":  true,
		"172.22.103.71":  true,
		"172.22.103.72":  true,
		"172.22.103.73":  true,
		"172.22.103.74":  true,
		"172.22.103.75":  true,
		"172.22.103.76":  true,
		"172.22.103.77":  true,
		"172.22.103.78":  true,
		"172.22.103.79":  true,
		"172.22.103.80":  true,
		"172.22.103.81":  true,
		"172.22.103.82":  true,
		"172.22.103.83":  true,
		"172.22.103.84":  true,
		"172.22.103.85":  true,
		"172.22.103.86":  true,
		"172.22.103.87":  true,
		"172.22.103.88":  true,
		"172.22.103.89":  true,
		"172.22.103.90":  true,
		"172.22.103.91":  true,
		"172.22.103.92":  true,
		"172.22.103.93":  true,
		"172.22.103.94":  true,
		"172.22.103.95":  true,
		"172.22.103.96":  true,
		"172.22.103.97":  true,
		"172.22.103.98":  true,
		"172.22.103.99":  true,
		"172.22.103.100": true,
	}
	file, err := json.Marshal(&Data)
	if err != nil {
		return
	}
	ioutil.WriteFile(CONFIG, file, 0644)
}
