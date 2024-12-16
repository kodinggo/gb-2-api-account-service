package console

import (
	"log"

	"github.com/kodinggo/gb-2-api-account-service/src/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gb-2-api-account-service",
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
