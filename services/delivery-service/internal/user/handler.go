package user

import (
	"context"
	"delivery-service/internal/middleware"
	"delivery-service/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Represents our handler with our use-case / service.
type UserHandler struct {
	userService UserService
}

// Creates a new handler.
func NewUserHandler(userRoute fiber.Router, us UserService) {
	// Create a handler based on our created service / use-case.
	handler := &UserHandler{
		userService: us,
	}

	// Declare routing endpoints for general routes.
	userRoute.Get("", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.getUsers)

	// Declare routing endpoints for sign user
	userRoute.Post("/sign/up", handler.createUser)
	userRoute.Post("/sign/in", handler.UserSignIn)
	userRoute.Post("/sign/out", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.UserSignOut)
	userRoute.Get("/sign/private", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.privateRoute)

	// Declare routing endpoints for specific routes.
	userRoute.Get("/:userID", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.getUser)
	userRoute.Put("/:userID", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.updateUser)
	userRoute.Delete("/:userID", middleware.JWTProtected(), middleware.ExtractTokenMetadata, handler.deleteUser)
}

// Gets all users.
func (h *UserHandler) getUsers(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get all users.
	users, err := h.userService.GetUsers(customContext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "user obtained succesfully",
		"http_code": fiber.StatusOK,
		"data":      users,
	})
}

// Gets a single user.
func (h *UserHandler) getUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid user ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Get one user.
	user, err := h.userService.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":    "fail",
			"message":   fmt.Sprintf("user of ID {%s} does not exist.", targetedUserID),
			"http_code": fiber.StatusNotFound,
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "user obtained succesfully",
		"http_code": fiber.StatusOK,
		"data":      user,
	})
}

// Creates a single user.
func (h *UserHandler) createUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables and Create a new user auth struct.
	signUp := &SignUp{}

	// Parse request body.
	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    "fail",
			"message":   utils.ValidatorErrors(err),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Create one user.
	user, err := h.userService.CreateUser(customContext, signUp)
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
		"message":   "User has been created successfully!",
		"http_code": fiber.StatusCreated,
		"data":      user,
	})
}

// Updates a single user.
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables.
	userUpdate := &UserUpdate{}

	// Fetch parameter.
	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid user ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Parse request body.
	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Update one user.
	user, err := h.userService.UpdateUser(customContext, targetedUserID, userUpdate)
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
		"message":   "User has been updated successfully!",
		"http_code": fiber.StatusOK,
		"data":      user,
	})
}

// Deletes a single user.
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize variables.
	userDelete := &UserDelete{}

	// Fetch parameter.
	targetedUserID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   "Please specify a valid user ID!",
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Parse request body.
	if err := c.BodyParser(userDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Update one user.
	err = h.userService.DeleteUser(customContext, targetedUserID, userDelete)
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
		"message":   "User has been deleted successfully!",
		"http_code": fiber.StatusOK,
	})
}

// Sign in users
func (h *UserHandler) UserSignIn(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new user auth struct.
	signIn := &SignIn{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signIn); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusBadRequest,
		})
	}

	// Get user by user name.
	tokens, err := h.userService.UserSignIn(customContext, signIn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return result 200 OK.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "User has been created successfully!",
		"http_code": fiber.StatusOK,
		"data":      tokens,
	})

}

// Sign out user
func (h *UserHandler) UserSignOut(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	userName := c.Locals("username").(string)

	// Get user by user name.
	err := h.userService.UserSignOut(customContext, userName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Return result 200 OK.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "session deleted",
		"http_code": fiber.StatusOK,
	})
}

// Private user
func (h *UserHandler) privateRoute(c *fiber.Ctx) error {
	// Create cancellable context.
	// customContext, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// Give form to our output response.
	type jwtResponse struct {
		UserID      interface{} `json:"user_id"`
		UserName    interface{} `json:"user_name"`
		Application interface{} `json:"application"`
		Audience    interface{} `json:"audience"`
		Expires     interface{} `json:"expires"`
	}

	// Shape output response.
	jwtResp := &jwtResponse{
		UserID:      c.Locals("userid"),
		UserName:    c.Locals("username"),
		Application: c.Locals("application"),
		Audience:    c.Locals("audience"),
		Expires:     c.Locals("expires"),
	}

	fmt.Println(jwtResp)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":    "success",
		"message":   "Welcome to the private route!",
		"http_code": fiber.StatusOK,
		"data":      jwtResp,
	})
}
