package main

import (
	"flag"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/teachain/leafserver/conf"
	"github.com/teachain/leafserver/game"
	"github.com/teachain/leafserver/gate"
	"github.com/teachain/leafserver/login"
)

func main() {
	flag.Parse()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
