package game

import "regexp"

func IsSteamIDValid(steamID string) bool {
	pattern := `^STEAM_[01]:\d+:\d+$`
	match, _ := regexp.MatchString(pattern, steamID)
	return match
}
