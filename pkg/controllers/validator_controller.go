package controllers

import (
	"net/http"

	"github.com/arhamj/go-commons/pkg/http_errors"
	"github.com/labstack/echo/v4"
	"github.com/shardeum/service-validator/pkg/dto"
	"github.com/shardeum/service-validator/pkg/service"
)

type ValidatorController struct {
	baseController
	validatorService service.ValidatorService
}

func NewValidatorController(validatorService service.ValidatorService) ValidatorController {
	return ValidatorController{
		baseController:   newBaseController(),
		validatorService: validatorService,
	}
}

func (v ValidatorController) GetEVMAccount(ctx echo.Context) error {
	var req dto.GetEVMAccountRequest
	if err := v.ValidateAndBind(ctx, &req); err != nil {
		return http_errors.NewBadRequestError(ctx, err.Error(), true)
	}
	resp, err := v.validatorService.GetEVMAccount(req.AddressID)
	if err != nil {
		return http_errors.ErrorCtxResponse(ctx, err, true)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (v ValidatorController) GetCodeBytes(ctx echo.Context) error {
	var req dto.GetCodeRequest
	if err := v.ValidateAndBind(ctx, &req); err != nil {
		return http_errors.NewBadRequestError(ctx, err.Error(), true)
	}
	resp, err := v.validatorService.GetCodeBytes(req.AddressID)
	if err != nil {
		return http_errors.ErrorCtxResponse(ctx, err, true)
	}
	return ctx.JSON(http.StatusOK, resp)
}
