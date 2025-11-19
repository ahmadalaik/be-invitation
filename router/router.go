package router

import (
	"github.com/ahmadalaik/be-invitation/internal/handler"
	"github.com/ahmadalaik/be-invitation/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	AuthHandler       *handler.AuthHandler
	User              *handler.UserHandler
	InvitationHandler *handler.InvitationHandler
}

func SetupRoutes(app *fiber.App, h *Handlers) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	public := v1.Group("/auth")
	{
		public.Post("/register", h.AuthHandler.Register)
		public.Post("/login", h.AuthHandler.Login)
	}

	protected := v1.Group("", middleware.AuthMiddleware)
	{
		users := protected.Group("/users")
		users.Post("/", h.User.Create)

		invitations := protected.Group("/invitations")
		invitations.Post("/", h.InvitationHandler.Create)
	}
}
