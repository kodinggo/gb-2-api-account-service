package console

import (
	"log"
	"net"
	"net/http"

	"github.com/kodinggo/gb-2-api-account-service/src/config"
	httphandler "github.com/kodinggo/gb-2-api-account-service/src/handler/http"
	"github.com/kodinggo/gb-2-api-account-service/src/repository"
	"github.com/kodinggo/gb-2-api-account-service/src/usecase"
	"google.golang.org/grpc"

	pb "github.com/kodinggo/gb-2-api-account-service/pb/account"
	grpchandler "github.com/kodinggo/gb-2-api-account-service/src/handler/grpc"

	dbmysql "github.com/kodinggo/gb-2-api-account-service/src/database/mysql"

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

		go func() {
			grpcServer := grpc.NewServer()

			accountgRPCHandler := grpchandler.NewAccountgRPCHandler(accountUsecase)

			pb.RegisterAccountServiceServer(grpcServer, accountgRPCHandler)

			httpListener, err := net.Listen("tcp", ":6000")
			if err != nil {
				log.Fatalf("failed to create gRPC listener: %v", err)
			}

			log.Println("gRPC server is running on port: 6000")

			if err := grpcServer.Serve(httpListener); err != nil {
				log.Fatalf("failed to start gRPC server: %v", err)
			}
		}()

		<-quitCh
	},
}

func init() {
	rootCmd.AddCommand(startServeCmd)
}
