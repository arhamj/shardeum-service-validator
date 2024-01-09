package controllers

import (
	"github.com/labstack/echo/v4"
)

type baseController struct {
}

func newBaseController() baseController {
	return baseController{}
}

func (b baseController) ValidateAndBind(ctx echo.Context, req any) error {
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return err
	}
	return nil
}
