package mssql

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/spf13/viper"
)

func buildURL() string {
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(viper.GetString("user"), viper.GetString("password")),
		Host:     fmt.Sprintf("%s:%d", viper.GetString("server"), viper.GetInt("port")),
		RawQuery: "database=master&connection+timeout=30",
		// Path:  instance, // if connecting to an instance instead of a port
	}
	return u.String()
}

func Connect() *sqlx.DB {
	db, err := sqlx.Open("sqlserver", buildURL())
	if err != nil {
		log.Fatal("Failed to connect to Database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to Database:", err)
	}
	return db

}
