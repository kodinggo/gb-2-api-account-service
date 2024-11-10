package httphandler

import (
	"account-service/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountUsecase model.AccountUsecase
}

func NewAccountHandler(accountUsecase model.AccountUsecase) *AccountHandler {
	return &AccountHandler{accountUsecase: accountUsecase}
}

func (handler *AccountHandler) RegisterRoute(e *echo.Echo) {
	groupping := e.Group("account")

	groupping.POST("/register", handler.register)
}

func (handler *AccountHandler) register(c echo.Context) error {
	var body model.Register

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("picture")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "picture file is required")
	}

	accessToken, err := handler.accountUsecase.CreateNewAccountData(c.Request().Context(), body, file)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}
