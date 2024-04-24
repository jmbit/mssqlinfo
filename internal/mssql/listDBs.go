package mssql

import (
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

type DBListEntry struct {
	Name          string
	Owner         string
	Size          int
	RecoveryModel string
}

func ListDatabases(db *sqlx.DB) []DBListEntry {
	result, err := db.Queryx(`
    SELECT 
      name AS DatabaseName,
      SUSER_SNAME(owner_sid) AS Owner,
      CAST((size * 8.0 / 1024) AS DECIMAL(18,2)) AS SizeMB,
      recovery_model_desc AS RecoveryModel 
    FROM sys.databases`,
	)
	if err != nil {
		log.Fatal("Could not get Databases:", err)
	}
	dbList := []DBListEntry{}
	for result.Next() {
		var entry DBListEntry
		result.StructScan(&entry)
		dbList = append(dbList, entry)

	}

	return dbList

}
