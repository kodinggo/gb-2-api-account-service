package console

import (
	"account-service/src/config"
	httphandler "account-service/src/handler/http"
	"account-service/src/repository"
	"account-service/src/usecase"
	"net/http"

	dbmysql "account-service/src/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var startServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the service server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := dbmysql.InitDBConn()

		accountRepository := repository.NewAccountRepository(dbConn)
		accountUsecase := usecase.NewAccountUsecase(accountRepository)
		quitCh := make(chan bool, 1)

		go func() {
			e := echo.New()

			e.GET("/ping", func(c echo.Context) error {
				return c.String(http.StatusOK, "pong!")
			})

			accountHandler := httphandler.NewAccountHandler(accountUsecase)

			accountHandler.RegisterRoute(e)

			e.Start(":" + config.Port())
		}()
		<-quitCh
	},
}

func init() {
	rootCmd.AddCommand(startServeCmd)
}
