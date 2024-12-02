package httphandler

import (
	"net/http"

	"github.com/kodinggo/gb-2-api-account-service/src/model"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountUsecase model.AccountUsecase
}

func NewAccountHandler(accountUsecase model.AccountUsecase) *AccountHandler {
	return &AccountHandler{accountUsecase: accountUsecase}
}

func (handler *AccountHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("v1/auth")

	g.POST("/register", handler.register)
	g.POST("/login", handler.login)
}

func (handler *AccountHandler) register(c echo.Context) error {
	var body model.Register
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := handler.accountUsecase.Create(c.Request().Context(), body)
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
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}
