package user

import (
	"github.com/gofiber/fiber/v2"
)

// If user does not exist, do not allow one to access the API.
func (h *UserHandler) checkIfUserExistsMiddleware(c *fiber.Ctx) error {
	// Create a new customized context.
	//customContext, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// Store in locals for further processing in the real handler.
	//c.Locals("userID", targetedUserID)
	return c.Next()
}
