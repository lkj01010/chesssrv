package com
import "errors"

const (
	E_Success = 0
	E_Fail = 1
	E_Unknown = 2
)

const (
	E_AgentAccountNotExist = 1000
	E_AgentAccountExist = 1001
	E_AgentPasswordIncorrect = 1002
	E_AgentPasswordCannotBeNull = 1003
)

// game
const (
	E_GameAlreadyInGame = 2000
)
var (
	ErrClientTimeout = errors.New("client timeout")
	ErrAgentNotFound = errors.New("agent not found")
	ErrCommandWithoutLogin = errors.New("command without login")


	// game
	ErrAlreadyInGame = errors.New("player already in game")
)