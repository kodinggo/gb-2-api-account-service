package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}

func ENV() string {
	return viper.GetString("env")
}

func Port() string {
	return viper.GetString("port")
}

func MySQLDBHost() string {
	return viper.GetString("mysql.dbhost")
}

func MySQLDBPort() string {
	return viper.GetString("mysql.dbport")
}

func MySQLDBUser() string {
	return viper.GetString("mysql.dbuser")
}

func MySQLDBPass() string {
	return viper.GetString("mysql.dbpass")
}

func MySQLDBName() string {
	return viper.GetString("mysql.dbname")
}

func JWTSigningKey() string {
	return viper.GetString("jwt.signing_key")
}

func JWTExp() time.Duration {
	return viper.GetDuration("jwt.exp")
}
