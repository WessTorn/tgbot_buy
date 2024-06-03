package database

import "database/sql"

type Stage int

const (
	ServerStage Stage = iota + 1
)

type Context struct {
	ID       int
	ChatID   int64
	Stage    Stage
	ServerID sql.NullInt64
}

type Server struct {
	ID   int
	Name string
	IP   string
}
