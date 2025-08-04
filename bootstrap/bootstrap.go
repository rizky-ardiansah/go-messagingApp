package bootstrap

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/ws"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/database"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/env"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/router"
	"go.elastic.co/apm"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	SetupLogfile()

	database.SetupDatabase()
	database.SetupMongoDB()

	apm.DefaultTracer.Service.Name = "simple-messaging-app"
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	go ws.ServerWSMessaging(app)

	router.InstallRouter(app)

	return app
}

func SetupLogfile() {
	logFile, err := os.OpenFile("./logs/simple_messaging_app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
