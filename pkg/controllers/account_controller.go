package controllers

import (
	"net/http"

	"github.com/arhamj/go-commons/pkg/http_errors"
	"github.com/labstack/echo/v4"
	"github.com/shardeum/service-validator/pkg/dto"
	"github.com/shardeum/service-validator/pkg/service"
)

type AccountController struct {
	baseController
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return AccountController{
		baseController: newBaseController(),
		accountService: accountService,
	}
}

func (a AccountController) GetAccount(ctx echo.Context) error {
	var req dto.GetAccountRequest
	if err := a.ValidateAndBind(ctx, &req); err != nil {
		return http_errors.NewBadRequestError(ctx, err.Error(), true)
	}
	account, err := a.accountService.GetAccount(req.AccountId)
	if err != nil {
		return http_errors.ErrorCtxResponse(ctx, err, true)
	}
	return ctx.JSON(http.StatusOK, account)
}

func (a AccountController) GetAccounts(ctx echo.Context) error {
	var req dto.GetAccountsRequest
	if err := a.ValidateAndBind(ctx, &req); err != nil {
		return http_errors.NewBadRequestError(ctx, err.Error(), true)
	}
	accounts, err := a.accountService.GetAccounts(req.AccountIds)
	if err != nil {
		return http_errors.ErrorCtxResponse(ctx, err, true)
	}
	return ctx.JSON(http.StatusOK, accounts)
}
