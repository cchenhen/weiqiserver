package main

import (
	"server/game"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
)

func main() {
	// lconf.LogLevel = conf.Server.LogLevel
	// lconf.LogPath = conf.Server.LogPath
	// lconf.LogFlag = conf.LogFlag
	// lconf.ConsolePort = conf.Server.ConsolePort
	// lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
