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

	mainImage, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "image is required"})
	}

	var mainImageURL string

	result, err := h.uploader.UploadSingle(c, mainImage)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	mainImageURL = result.FilePath

	// story images
	form, _ := c.MultipartForm()
	storyFiles := form.File["story_images"]
	var storyImageURLs []string

	if storyFiles != nil {
		results, err := h.uploader.UploadMultiple(c, storyFiles)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		for _, r := range results {
			storyImageURLs = append(storyImageURLs, r.FilePath)
		}
	}

	userID, ok := c.Locals("userID").(int64)
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
