package main

import (
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/user", User)
	app.Post("/api/logout", Logout)

}
