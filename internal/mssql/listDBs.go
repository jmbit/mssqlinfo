package mssql

import (
	"log"

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
      sys.databases.name AS name,
      SUSER_SNAME(sys.databases.owner_sid) AS owner,
      CAST((sys.master_files.size * 8.0 / 1024) AS INT) AS size,
      sys.databases.recovery_model_desc AS recoverymodel
    FROM sys.databases, sys.master_files
    WHERE sys.databases.name = sys.master_files.name`,
	)
	if err != nil {
		log.Fatal("Could not get Databases:", err)
	}
	dbList := []DBListEntry{}
	for result.Next() {
		var entry DBListEntry
		err := result.StructScan(&entry)
		if err != nil {
			log.Println(err)
		}
		dbList = append(dbList, entry)

	}

	return dbList

}
