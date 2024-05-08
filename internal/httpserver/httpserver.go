package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmbit/mssqlinfo/internal/mssql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var dbcon *sqlx.DB

func StartServer() {
	var err error
	address := viper.GetString("address")
	if address == "" {
		address = "localhost:3341"
	}
	log.Println("Starting server on", address)
	server := http.Server{
		Addr:    address,
		Handler: registerRoutes(),
	}
	dbcon, err = mssql.Connect()
	if err != nil {
		log.Fatalln("Could not connect to MSSQL server:", err)

	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	fmt.Fprint(w, "This is my website!\n")
}
