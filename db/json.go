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
		Token:    "101111",
	}
	Data.Machines["172.22.103.6"] = &model.Machine{
		IP:   "172.22.103.6",
		Use:  [4]int{0, 0, 0, 0},
		Load: [4]int{0, 0, 0, 0},
	}
	file, err := json.Marshal(&Data)
	if err != nil {
		return
	}
	ioutil.WriteFile(CONFIG, file, 0644)
}
