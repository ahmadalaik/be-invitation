package handler

import (
	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/service"
	"github.com/ahmadalaik/be-invitation/internal/validator"
	"github.com/ahmadalaik/be-invitation/pkg/uploader"
	"github.com/gofiber/fiber/v2"
)

type InvitationHandler struct {
	invService service.InvitationService
	uploader   *uploader.Uploader
	validator  validator.Validator
}

func NewInvitationHandler(invService service.InvitationService, uploader *uploader.Uploader, validator validator.Validator) *InvitationHandler {
	return &InvitationHandler{invService, uploader, validator}
}

func (h *InvitationHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateInvitationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var mainImageURL string

	result, err := h.uploader.UploadSingle(c, "image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	mainImageURL = result.FilePath

	// story images
	var storyImageURLs []string

	results, err := h.uploader.UploadMultiple(c, "story_images")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, r := range results {
		storyImageURLs = append(storyImageURLs, r.FilePath)
	}

	userID, ok := c.Locals("userID").(uint64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not authenticated"})
	}

	response, err := h.invService.CreateInvitation(userID, req, mainImageURL, storyImageURLs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "invitation create successfully",
		"data":    response,
	})
}

func (h *InvitationHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	response, err := h.invService.GetInvitationBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (h *InvitationHandler) GetAllByUserID(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not authenticated"})
	}

	responses, err := h.invService.GetAllInvitationsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": responses,
	})
}

func (h *InvitationHandler) Update(c *fiber.Ctx) error {
	var req dto.UpdateInvitationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var newMainImageURL string

	result, err := h.uploader.UploadSingle(c, "image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newMainImageURL = result.FilePath

	slug := c.Params("slug")

	if err := h.invService.UpdateInvitationBySlug(slug, req, newMainImageURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "invitation update successfully"})
}
