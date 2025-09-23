package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adii1203/video-cred/internals/service"
	"github.com/adii1203/video-cred/internals/storage"
	"github.com/adii1203/video-cred/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct {
	service *service.UserService
	logger  *slog.Logger
}

func NewUserHandler(service *service.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) ClerkHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		headers := c.GetReqHeaders()
		payload := c.Body()

		h.logger.Info("clerk webhook handler")

		// verify webhook
		wh := pkg.InitSvix()

		if err := wh.Verify(payload, headers); err != nil {
			h.logger.Warn("Webhook verification failed", "errors", err.Error(), "headers", headers)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid webhook signature",
			})
		}

		var clerkEvt pkg.ClerkUserCreated

		if clerkEvt.Type == "user.created" {
			if err := json.Unmarshal(payload, &clerkEvt); err != nil {
				h.logger.Error("failed to parse clerk event payload", "error", err.Error(),
					"payload", string(payload))
				return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "invalid event payload",
				})
			}

			if len(clerkEvt.Data.EmailAddresses) == 0 || clerkEvt.Data.EmailAddresses[0].EmailAddress == "" {
				return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "invalid event payload",
				})
			}

			pramps := storage.CreateUserParams{
				Name:    fmt.Sprintf("%s %s", clerkEvt.Data.FirstName, clerkEvt.Data.LastName),
				Clerkid: pgtype.Text{String: clerkEvt.Data.Id, Valid: true},
				Email:   clerkEvt.Data.EmailAddresses[0].EmailAddress,
			}

			err := h.service.CreateUserWithClerk(c.Context(), pramps)
			if err != nil {
				if ctx.Err() != nil {
					h.logger.Warn("request context canceled",
						"clerkID", clerkEvt.Data.Id,
						"error", err.Error())

					return c.Status(http.StatusRequestTimeout).JSON(fiber.Map{
						"error": "request context canceled",
					})
				}

				h.logger.Error("Failed to create user from Clerk event",
					"clerkID", clerkEvt.Data.Id,
					"error", err.Error(),
				)
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to process webhook",
				})
			}

			h.logger.Info("successfully processed clerk webhook",
				"clerkID", clerkEvt.Data.Id,
				"email", clerkEvt.Data.EmailAddresses[0].EmailAddress,
			)
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"status": "ok",
			})
		}
		return nil
	}
}
