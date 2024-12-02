package httphandler

import (
	"account-service/src/model"
	"net/http"
	"strconv"

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

	g.GET("/account/:id", handler.show)
	g.POST("/register", handler.register)
	g.POST("/login", handler.login)
	g.PUT("/account/:id/update", handler.update)

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

func (handler *AccountHandler) show(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	account, err := handler.accountUsecase.FindById(c.Request().Context(), model.Account{}, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Data: account,
	})
}

func (handler *AccountHandler) update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}
	var body model.Account
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	account, err := handler.accountUsecase.Update(c.Request().Context(), body, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())

	}

	return c.JSON(http.StatusOK, Response{
		Data: account,
	})
}
