package console

import (
	"account-service/account-service/src/repository"

	"github.com/spf13/cobra"
)

var startServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the service server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := dbmysql.InitDBConn()

		accountRepository := repository.NewAccountRepository(dbConn)

	},
}

func init() {
	rootCmd.AddCommand(startServeCmd)
}
