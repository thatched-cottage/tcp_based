package command

type CommandType byte

const (
	CommandConn   CommandType = 1
	CommandMirror CommandType = 2
)
