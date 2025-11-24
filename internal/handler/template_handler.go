package handler

import (
	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/service"
	"github.com/ahmadalaik/be-invitation/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type TemplateHandler struct {
	templateService service.TemplateService
	validator       validator.Validator
}

func NewTemplateHandler(templateService service.TemplateService, validator validator.Validator) *TemplateHandler {
	return &TemplateHandler{templateService, validator}
}

func (h *TemplateHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateTemplateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.templateService.CreateTemplate(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "invitation create successfully",
	})
}
