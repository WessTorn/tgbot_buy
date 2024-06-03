package database

import "database/sql"

type Stage int

const (
	ServerStg Stage = iota + 1
	ServiceStg
	PrivilegeStg
	PrlgSteamStg
	PrlgNickStg
)

type Context struct {
	ID       int
	ChatID   int64
	Stage    Stage
	ServerID sql.NullInt64
	SteamID  sql.NullString
	Nick     sql.NullString
}

type Server struct {
	ID   int
	Name string
	IP   string
}
