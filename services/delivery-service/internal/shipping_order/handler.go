package shipping_order

import (
	"context"
	"delivery-service/internal/middleware"
	"delivery-service/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Represents our handler with our use-case / service.
type ShippingOrderHandler struct {
	shippingOrderService ShippingOrderService
}

// Creates a new handler.
func NewShippingOrderHandler(shippingOrderRoute fiber.Router, us ShippingOrderService) {
	// Create a handler based on our created service / use-case.
	handler := &ShippingOrderHandler{
		shippingOrderService: us,
	}

	// We will restrict this route with our JWT middleware.
	// You can inject other middlewares if you see fit here.
	shippingOrderRoute.Use(middleware.JWTProtected(), middleware.ExtractTokenMetadata)

	// Declare routing endpoints for general routes.
	shippingOrderRoute.Post("", handler.createShippingOrder)
	shippingOrderRoute.Get("/:shippingOrderID", handler.getShippingOrder)
	shippingOrderRoute.Put("/:shippingOrderID", handler.updateShippingOrder)
	shippingOrderRoute.Delete("/:shippingOrderID", handler.cancelShippingOrder)

	// Declare routing endpoints for specific routes.
	shippingOrderRoute.Post("/sender", handler.createShippingOrder)
	shippingOrderRoute.Get("/:shippingOrderID/sender/:senderID", handler.checkIfShippingOrderExistsMiddleware, handler.getShippingOrder)
	shippingOrderRoute.Delete("/:shippingOrderID/sender/:senderID", handler.checkIfShippingOrderExistsMiddleware, handler.cancelShippingOrder)
}

// Gets a single shippingOrder.
func (h *ShippingOrderHandler) getShippingOrder(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedShippingOrderID, err := c.ParamsInt("shippingOrderID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid shippingOrder ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Get one shippingOrder.
	shippingOrder, err := h.shippingOrderService.GetShippingOrder(customContext, targetedShippingOrderID)
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

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "shippingOrder obtained succesfully",
		"http_code": fiber.StatusOK,
		"data":      shippingOrder,
	})
}

// Creates a single shippingOrder.
func (h *ShippingOrderHandler) createShippingOrder(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables and Create a new shippingOrder auth struct.
	shippingOrderInsert := &ShippingOrderInsert{}

	// Parse request body.
	if err := c.BodyParser(shippingOrderInsert); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create a new validator for a ShippingOrder model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(shippingOrderInsert); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    "fail",
			"message":   utils.ValidatorErrors(err),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create one shippingOrder.
	shippingOrder, err := h.shippingOrderService.CreateShippingOrder(customContext, shippingOrderInsert)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":    "success",
		"message":   "ShippingOrder has been created successfully!",
		"http_code": fiber.StatusCreated,
		"data":      shippingOrder,
	})
}

// Updates a single shippingOrder.
func (h *ShippingOrderHandler) updateShippingOrder(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables.
	shippingOrderUpdate := &ShippingOrderUpdate{}

	// Fetch parameter.
	targetedShippingOrderID, err := c.ParamsInt("shippingOrderID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid shippingOrder ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Parse request body.
	if err := c.BodyParser(shippingOrderUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create a new validator for a ShippingOrder model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(shippingOrderUpdate); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    "fail",
			"message":   utils.ValidatorErrors(err),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Update one shippingOrder.
	shippingOrder, err := h.shippingOrderService.UpdateShippingOrder(customContext, targetedShippingOrderID, shippingOrderUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "ShippingOrder has been updated successfully!",
		"http_code": fiber.StatusOK,
		"data":      shippingOrder,
	})
}

// Deletes a single shippingOrder.
func (h *ShippingOrderHandler) cancelShippingOrder(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables.
	shippingOrderCancel := &ShippingOrderCancel{}

	// Fetch parameter.
	targetedShippingOrderID, err := c.ParamsInt("shippingOrderID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid shippingOrder ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Parse request body.
	if err := c.BodyParser(shippingOrderCancel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create a new validator for a ShippingOrder model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(shippingOrderCancel); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    "fail",
			"message":   utils.ValidatorErrors(err),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Update one shippingOrder.
	err = h.shippingOrderService.CancelShippingOrder(customContext, targetedShippingOrderID, shippingOrderCancel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "Shipping Order has been cancelled successfully!",
		"http_code": fiber.StatusOK,
	})
}
