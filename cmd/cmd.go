package cmd

import (
	"time"
)

type user struct {
	guid          int
	name          string `json:"name"`
	username      string
	telephone     string
	password_hash string    `json:"password_hash"`
	token_hash    string    `json:"token_hash"`
	creation_ts   time.Time `json:"creation_ts"`
	upgrade_ts    time.Time `json:"upgrade_ts"`
	deactivated   bool
}

func Execute() {
	print("OK")
}
