package get_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tg_cs/logger"
)

type Privilege struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Flags       string `json:"flags"`
	Days        []struct {
		DayID int64 `json:"day_id"`
		Day   int   `json:"day"`
		Price int   `json:"price"`
	} `json:"days"`
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

	err = json.Unmarshal(data, &privileges) // TODO: Проверить айди элементов
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

func GetPrivilegeFromID(privilegeID int64) (Privilege, error) {
	var out Privilege
	fmt.Printf("privilegeID %d", privilegeID)

	var check bool = false
	for _, privilege := range privileges.Privilege {
		if privilegeID == privilege.ID {
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

func GetDayIDFromString(privilege Privilege, buttonName string) (int64, error) {
	numbers := strings.Split(buttonName, " ")

	day, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, err
	}

	dayID, err := getDayIDFromDay(privilege, day)
	if err != nil {
		return 0, err
	}

	return dayID, nil
}

func getDayIDFromDay(privilege Privilege, day int) (int64, error) {
	for _, days := range privilege.Days {
		if days.Day == day {
			return days.DayID, nil
		}
	}

	return 0, errors.New("DayIDNotFound")
}

func GetDayFromDayID(privilege Privilege, DayID int64) int {
	for _, days := range privilege.Days {
		if DayID == days.DayID {
			return days.Day
		}
	}

	return 0
}

func GetPrivileges() Privileges {
	return privileges
}
