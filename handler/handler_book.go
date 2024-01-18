package handler

import (
	"echo-ci-cd/usecase"
	"echo-ci-cd/utils"

	"github.com/labstack/echo/v4"
)

type BookHandler interface {
	GetBookByID(ctx echo.Context) error
}

type handler struct {
	bookHandler usecase.BookUsecase
}

func NewHandler(bookHandler usecase.BookUsecase) *handler {
	return &handler{
		bookHandler: bookHandler,
	}
}

func (h *handler) GetBookByID(ctx echo.Context) error {
	uuid := ctx.Param("uuid")
	result, err := h.bookHandler.GetBookByID(ctx.Request().Context(), uuid)

	return utils.Response(ctx, result, err)
}

func (h *handler) GetBooks(ctx echo.Context) error {
	result, err := h.bookHandler.GetBooks(ctx.Request().Context())

	return utils.Response(ctx, result, err)
}
