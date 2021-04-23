package model

import (
	"time"
)

type User struct {
	DockerName  string    `json:"dockername"`
	DockerTag   string    `json:"dockertag"`
	Names       string    `json:"names"`
	Description string    `json:"description"`
	IP          string    `json:"ip"`
	Machineip   string    `json:"machineip"`
	CPUs        string    `json:"cpus"`
	GPUs        string    `json:"gpus"`
	Mem         string    `json:"mem"`
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:"endtime"`
	Status      int       `json:"status"`
	Password    []byte    `json:"password"`
}
