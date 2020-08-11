package bootstrap

import (
	"github.com/vidar-team/Cardinal/conf"
	"github.com/vidar-team/Cardinal/internal/asteroid"
	"github.com/vidar-team/Cardinal/internal/db"
	"github.com/vidar-team/Cardinal/internal/dynamic_config"
	"github.com/vidar-team/Cardinal/internal/livelog"
	"github.com/vidar-team/Cardinal/internal/misc"
	"github.com/vidar-team/Cardinal/internal/route"
	"github.com/vidar-team/Cardinal/internal/store"
	"github.com/vidar-team/Cardinal/internal/timer"
	"log"
)

// LinkStart starts the Cardinal.
func LinkStart() {
	// Install

	// Sentry
	misc.Sentry()

	// Init MySQL database.
	db.InitMySQL()
	autoMigrate()

	// Refresh the dynamic config from the database.
	dynamic_config.Init()

	// Game timer.
	timer.Init()

	// Cache
	store.Init()

	// Unity3D Asteroid
	asteroid.Init()

	// Live log
	livelog.Init()

	// Web router.
	router := route.Init()

	log.Fatalln(router.Run(conf.Get().Port))
}
