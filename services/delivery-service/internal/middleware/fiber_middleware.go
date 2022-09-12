package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// registra detalles de solicitud/respuesta HTTP.
		logger.New(),
		// Fiber que comprimirá la respuesta mediante compresión gzip, deflate y brotli según el encabezado Accept-Encoding.
		compress.New(),
		// permite que los cachés sean más eficientes y ahorren ancho de banda,
		// ya que un servidor web no necesita volver a enviar una respuesta completa si el contenido no ha cambiado.
		etag.New(),
		// Los agentes de usuario solicitan favicon.ico con frecuencia e indiscriminadamente,
		//por lo que es posible que desee excluir estas solicitudes
		favicon.New(),
		// recupera de pánicos en cualquier parte de la cadena de pila y maneja el control al ErrorHandler centralizado.
		recover.New(),
		// RequestID para Fiber que agrega un identificador a la respuesta.
		requestid.New(),
		// se usa para limitar las solicitudes repetidas a las API
		limiter.New(limiter.Config{
			Max: 100000,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
					"status":  "fail",
					"message": "You have requested too many in a single time-frame! Please wait another minute!",
				})
			},
		}),
	)
}
