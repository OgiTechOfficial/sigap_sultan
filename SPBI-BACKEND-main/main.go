package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"sigap-sultan-be/src/app"
	"sigap-sultan-be/src/config"
	_ "sigap-sultan-be/src/routes"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		BodyLimit:    50 * 1024 * 1024,
	})

	fiberApp.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		config.LogSetup(),
	)
	//fiberApp.Use(
	//	cors.New(
	//		cors.Config{
	//			AllowOrigins: "http://*.sentech.id",
	//		},
	//	),
	//)

	envConfig := config.ReadEnv()
	dbConfig := config.NewDbConfig(envConfig)
	initConfig := app.InitConfig{
		App:       fiberApp,
		Db:        dbConfig.StartDbConnection(),
		Redis:     config.StartRedisConnection(),
		EnvConfig: envConfig,
	}

	app.InitModules(initConfig)

	err := fiberApp.Listen(envConfig.AddrListen)
	if err != nil {
		panic(err)
	}
}
