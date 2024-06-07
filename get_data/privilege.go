package get_data

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"tg_cs/logger"
)

type Privilege struct {
	ID          int64
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

func GetPrivilegeFromName(privilegeName string) (Privilege, error) {
	var out Privilege

	var check bool = false
	for _, privilege := range privileges.Privilege {
		if strings.EqualFold(privilegeName, privilege.Name) {
			out = privilege
			check = true
			break
		}
	}

	if !check {
		return out, errors.New("PrivilegeNotFound")
	}

	return out, nil
}

func GetPrivileges() Privileges {
	return privileges
}
