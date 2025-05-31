package adapters

import (
	"github.com/frostnzx/go-myapi-assignment/internal/core"
	"github.com/gofiber/fiber/v2"
)

type HttpProfileHandler struct {
	service core.ProfileService
}

func NewHttpProfileHandler(service core.ProfileService) *HttpProfileHandler {
	return &HttpProfileHandler{service : service}
}

func (h *HttpProfileHandler) CreateProfile(c *fiber.Ctx) error {
	var profile core.Profile
	if err := c.BodyParser(&profile) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := h.service.CreateProfile(profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(profile)
}
func (h *HttpProfileHandler) GetProfiles(c *fiber.Ctx) error {
	profiles, err := h.service.GetProfiles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(profiles)
}