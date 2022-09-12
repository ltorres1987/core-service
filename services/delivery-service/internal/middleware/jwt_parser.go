package middleware

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"delivery-service/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) error {
	token, tokenString, err := verifyToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":    "fail",
			"message":   err.Error(),
			"http_code": fiber.StatusInternalServerError,
		})
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		parsedUserID := claims["id"].(string)

		// User name
		userName := claims["user"].(string)

		// User name
		application := claims["application"].(string)

		//Audience
		audience := claims["Audience"].(string)

		// Expires time.
		expires := int64(claims["expires"].(float64))

		userID, err := strconv.Atoi(parsedUserID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":    "fail",
				"message":   err.Error(),
				"http_code": fiber.StatusInternalServerError,
			})
		}

		// Get now time.
		now := time.Now().Unix()

		// Checking, if now time greather than expiration from JWT.
		if now > expires {
			// Return status 401 and unauthorized error message.
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":    "fail",
				"message":   fmt.Sprintf("unauthorized, check expiration time of your token"),
				"http_code": fiber.StatusUnauthorized,
			})
		}

		// Create a new Redis connection.
		connRedis, err := utils.RedisConnection()
		if err != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":    "fail",
				"message":   err.Error(),
				"http_code": fiber.StatusInternalServerError,
			})
		}

		// Get token to Redis.
		val, err := connRedis.Get(context.Background(), userName).Result()
		if err != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":    "fail",
				"message":   err.Error(),
				"http_code": fiber.StatusInternalServerError,
			})
		}

		// valid token vs redis
		if val != tokenString {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":    "fail",
				"message":   "invalid token - redis not registered",
				"http_code": fiber.StatusInternalServerError,
			})
		}

		// Go to next.
		c.Locals("userid", userID)
		c.Locals("username", userName)
		c.Locals("application", application)
		c.Locals("audience", audience)
		c.Locals("expires", expires)
		return c.Next()
	}

	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
		"status":    "fail",
		"message":   err.Error(),
		"http_code": fiber.StatusInternalServerError,
	})
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, string, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, "", err
	}

	return token, tokenString, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
