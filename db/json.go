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
		"172.22.103.50":  false,
		"172.22.103.51":  false,
		"172.22.103.52":  false,
		"172.22.103.53":  false,
		"172.22.103.54":  false,
		"172.22.103.55":  false,
		"172.22.103.56":  false,
		"172.22.103.57":  false,
		"172.22.103.58":  false,
		"172.22.103.59":  false,
		"172.22.103.60":  false,
		"172.22.103.61":  false,
		"172.22.103.62":  false,
		"172.22.103.63":  false,
		"172.22.103.64":  false,
		"172.22.103.65":  false,
		"172.22.103.66":  false,
		"172.22.103.67":  false,
		"172.22.103.68":  false,
		"172.22.103.69":  false,
		"172.22.103.70":  false,
		"172.22.103.71":  false,
		"172.22.103.72":  false,
		"172.22.103.73":  false,
		"172.22.103.74":  false,
		"172.22.103.75":  false,
		"172.22.103.76":  false,
		"172.22.103.77":  false,
		"172.22.103.78":  false,
		"172.22.103.79":  false,
		"172.22.103.80":  false,
		"172.22.103.81":  false,
		"172.22.103.82":  false,
		"172.22.103.83":  false,
		"172.22.103.84":  false,
		"172.22.103.85":  false,
		"172.22.103.86":  false,
		"172.22.103.87":  false,
		"172.22.103.88":  false,
		"172.22.103.89":  false,
		"172.22.103.90":  false,
		"172.22.103.91":  false,
		"172.22.103.92":  false,
		"172.22.103.93":  false,
		"172.22.103.94":  false,
		"172.22.103.95":  false,
		"172.22.103.96":  false,
		"172.22.103.97":  false,
		"172.22.103.98":  false,
		"172.22.103.99":  false,
		"172.22.103.100": false,
	}
	file, err := json.Marshal(&Data)
	if err != nil {
		return
	}
	ioutil.WriteFile(CONFIG, file, 0644)
}
