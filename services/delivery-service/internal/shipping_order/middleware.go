package shipping_order

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// If user does not exist, do not allow one to access the API.
func (h *ShippingOrderHandler) checkIfShippingOrderExistsMiddleware(c *fiber.Ctx) error {
	// Create a new customized context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedSenderID := c.Params("senderID")

	targetedShippingOrderID, err := c.ParamsInt("shippingOrderID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid shippingOrder ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Check if order exists.
	shippingOrder, err := h.shippingOrderService.GetSenderShippingOrder(customContext, targetedShippingOrderID, targetedSenderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	if shippingOrder == nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":    "fail",
			"message":   fmt.Sprintf("shippingOrder of ID {%d} does not exist.", targetedShippingOrderID),
			"http_code": fiber.StatusNotFound,
		})
	}

	return c.Next()
}
