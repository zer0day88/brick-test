package handler

import (
	"github.com/labstack/echo/v4"
	model_pg "github.com/zer0day88/brick-test/internal/app/model"
	"github.com/zer0day88/brick-test/internal/app/service"
	pkgerr "github.com/zer0day88/brick-test/pkg/error"
)

type pgHandler struct {
	pgSrv *service.PG
}

type PGHandler interface {
	InquiryAccount(c echo.Context) error
	Transfer(c echo.Context) error
	NotifyTransfer(c echo.Context) error
}

func NewPgHandler(pgSrv *service.PG) PGHandler {
	return &pgHandler{pgSrv: pgSrv}
}

func (h *pgHandler) InquiryAccount(c echo.Context) error {
	var req model_pg.InquiryAccountRequest

	if err := c.Bind(&req); err != nil {
		return res.JSON(c, nil, pkgerr.InquiryAccountBadParam)
	}

	data, err := h.pgSrv.InquiryAccount(c.Request().Context(), req)
	if data == nil {
		return res.JSON(c, nil, err)
	}

	return res.JSON(c, data, err)
}

func (h *pgHandler) Transfer(c echo.Context) error {
	var req model_pg.TransferRequest

	if err := c.Bind(&req); err != nil {
		return res.JSON(c, nil, pkgerr.TransferBadParam)
	}

	data, err := h.pgSrv.Transfer(c.Request().Context(), req)
	if data == nil {
		return res.JSON(c, nil, err)
	}

	return res.JSON(c, data, err)
}

// NotifyTransfer waiting to hit by bank
func (h *pgHandler) NotifyTransfer(c echo.Context) error {
	var req model_pg.NotifyTransferRequest

	if err := c.Bind(&req); err != nil {
		return res.JSON(c, nil, pkgerr.NotifyTransferBadParam)
	}

	err := h.pgSrv.NotifyTransfer(c.Request().Context(), req)

	return res.JSON(c, nil, err)
}
