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

type ContextPrlg struct {
	ID          int
	ChatID      int64
	PrivilegeID sql.NullInt64
	SteamID     sql.NullString
	Nick        sql.NullString
}

type Context struct {
	ID        int
	ChatID    int64
	Stage     Stage
	ServerID  sql.NullInt64
	ServiceID sql.NullInt64
	Prvg      ContextPrlg
}

type Server struct {
	ID   int64
	Name string
	IP   string
}
