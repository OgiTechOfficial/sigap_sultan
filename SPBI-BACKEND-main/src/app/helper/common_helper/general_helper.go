package common_helper

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/common"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateResponseJSON(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}

func ProcessErrorDomain(c *fiber.Ctx, errorDomain *common.ErrorDomain) error {
	if errorDomain.Message == "no rows in result set" {
		return GenerateResponseJSON(c, http.StatusOK, models.Set().Success("Data is not found", nil))
	} else {
		return GenerateResponseJSON(c, http.StatusInternalServerError, models.Set().Success(errorDomain.Message, nil))
	}
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[random.Int64()]
	}
	return string(b)
}
