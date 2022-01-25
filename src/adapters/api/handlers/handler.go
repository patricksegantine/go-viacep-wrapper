package handlers

import (
	"net/http"

	s "go-viacep-wrapper/internal/application/services"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Handler struct {
	profileService s.ProfileService
}

func NewHandler(profileService s.ProfileService) *Handler {
	return &Handler{profileService: profileService}
}

func (h Handler) SetUpRoutes(app fiber.Router) {
	app.Get("/find-address-by-zipcode/:zipcode", h.findAddressByZipcode)
}

func (h Handler) findAddressByZipcode(ctx *fiber.Ctx) error {
	zipcode := ctx.Params("zipcode")
	result, err := h.profileService.FindAddress(zipcode)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			Response{
				Message: err.Error(),
				Data:    nil,
			})
	}

	ctx.Set("Content-Type", "application/json")
	return ctx.Status(200).JSON(&Response{Data: result})
}
