package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmbit/mssqlinfo/internal/mssql"
)

func serverHandler(w http.ResponseWriter, r *http.Request) {
	result, err := mssql.GetServerInfo(dbcon)
	jsonResp, err := json.Marshal(result)

	if err != nil {
		log.Printf("error handling JSON marshal. Err: %v", err)
		r.Response.StatusCode = http.StatusInternalServerError

		fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
	}

	_, _ = w.Write(jsonResp)
}
