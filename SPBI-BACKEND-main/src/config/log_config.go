package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
)

func LogSetup() fiber.Handler {
	// Define file to logs
	file, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer file.Close()

	loggerHandler := logger.New(
		logger.Config{
			Next:          nil,
			Done:          nil,
			Format:        "Time & Latency: ${time} -${latency} | PID: ${pid} | From: ${ip} | HttpStatus: ${status} | ${method} | ${path} | Error: ${error}\n",
			TimeFormat:    "15:04:05",
			TimeZone:      "Local",
			TimeInterval:  500 * time.Millisecond,
			Output:        file,
			DisableColors: false,
		},
	)
	//log.SetLevel(log.LevelInfo)
	//var logger AllLogger = &defaultLogger{
	//	stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
	//	depth:  4,
	//}

	//Output to ./test.log file
	//file, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//iw := io.MultiWriter(os.Stdout, file)
	//log.SetOutput(iw)

	//commonLogger := log.WithContext(c.Context())
	//commonLogger.Info("info")

	return loggerHandler
}
