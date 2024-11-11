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
	groupping := e.Group("auth")

	groupping.POST("/register", handler.register)
	groupping.POST("/login", handler.login)
}

func (handler *AccountHandler) register(c echo.Context) error {
	var body model.Register

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := handler.accountUsecase.CreateAccount(c.Request().Context(), body)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}

func (handler *AccountHandler) login(c echo.Context) error {
	var body model.Login

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := handler.accountUsecase.Login(c.Request().Context(), body)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}
