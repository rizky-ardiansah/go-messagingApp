package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/ws"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/database"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/env"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/router"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	database.SetupMongoDB()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	go ws.ServerWSMessaging(app)

	router.InstallRouter(app)

	return app
}
