package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type ConnDBURL struct {
	DBUrl *map[string]*sql.DB
}

func (h *ConnDBURL) ConnectDBFirst(c *fiber.Ctx) error {

	var generic *sql.DB
	var ori *sql.DB

	for k, v := range *h.DBUrl {

		if k == "GENERIC" {
			generic = v
		}
		if k == "GENERIC_ORI" {
			ori = v
		}
	}

	*generic = *ori

	return c.Next()
}

// If user does not exist, do not allow one to access the API.
func (h *ConnDBURL) ConnectDBSecond(c *fiber.Ctx) error {

	application := c.Locals("application")

	var generic *sql.DB
	var app *sql.DB

	for k, v := range *h.DBUrl {

		if k == "GENERIC" {
			generic = v
		}
		if k == application {
			app = v
		}
	}

	*generic = *app

	return c.Next()
}
