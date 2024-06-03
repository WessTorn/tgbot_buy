package get_data

import (
	"encoding/json"
	"os"
	"tg_cs/logger"
)

type Privilege struct {
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        []struct {
		Days  int `json:"days"`
		Price int `json:"price"`
	} `json:"cost"`
}

type Privileges struct {
	Privilege []Privilege `json:"privileges"`
}

var privileges Privileges

func ReadPrivilege() {
	data, err := os.ReadFile("privileges.json")
	if err != nil {
		logger.Log.Fatalf("(ReadFile) %v", err)
		return
	}

	err = json.Unmarshal(data, &privileges)
	if err != nil {
		logger.Log.Fatalf("(Unmarshal) %v", err)
		return
	}
}

func GetPrivileges() Privileges {
	return privileges
}
