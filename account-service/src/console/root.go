package console

import (
	"account-service/account-service/src/config"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "account-service",
	Short: "api service for account api gateway",
}

func init() {
	config.InitConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
