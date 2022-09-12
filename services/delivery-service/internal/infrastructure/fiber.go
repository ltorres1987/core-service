package infrastructure

import (
	"delivery-service/internal/configs"
	"delivery-service/internal/middleware"
	"delivery-service/internal/misc"
	"delivery-service/internal/package_size"
	"delivery-service/internal/shipping_order"
	"delivery-service/internal/user"
	"delivery-service/internal/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
)

// Run our Fiber webserver.
func Run() {
	// Try to connect to our database as the initial part.
	mariadb, err := ConnectToMariaDB()
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	// Define Fiber config.
	config := configs.FiberConfig()

	// Creates a new Fiber instance.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Create repositories.
	userRepository := user.NewUserRepository(mariadb)
	shippingOrderRepository := shipping_order.NewShippingOrderRepository(mariadb)
	packageSizeRepository := package_size.NewPackageSizeRepository(mariadb)

	// Create all of our services.
	userService := user.NewUserService(userRepository)
	shippingOrderService := shipping_order.NewShippingOrderService(shippingOrderRepository, packageSizeRepository)

	// Prepare our endpoints for the API.
	misc.NewMiscHandler(app.Group("/api/v1"))
	user.NewUserHandler(app.Group("/api/v1/users"), userService)
	shipping_order.NewShippingOrderHandler(app.Group("/api/v1/order"), shippingOrderService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
			"code":    fiber.StatusNotFound,
			"data":    "",
		})
	})

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
