package main

import (
	"time"

	"netty/core"
	"netty/db"
	"netty/register/router"
)

func main() {
	if err := core.SetupConfig(); err != nil {
		panic(err)
	}

	location, _ := time.LoadLocation(core.Settings.ServerTimeZone)
	time.Local = location

	databaseURL := core.BuildPgDsn()

	if err := db.DbSession(databaseURL); err != nil {
		panic(err)
	}

	router := router.SetUpRoute()

	router.Run(core.Settings.ServerAddress)
}
