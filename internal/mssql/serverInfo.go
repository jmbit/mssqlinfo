package mssql

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type ServerInfo struct {
	Version          string
	ProductLevel     string
	Edition          string
	EngineEdition    string
	MachineName      string
	IsClustered      bool
	Collation        string
	CLRVersion       string
	ProcessorCount   sql.NullInt32
	PhysicalMemoryMB sql.NullInt32
}

func GetServerInfo(db *sqlx.DB) (ServerInfo, error) {
	var info ServerInfo
	result, err := db.Queryx(`SELECT 
    SERVERPROPERTY('ProductVersion') AS version,
    SERVERPROPERTY('ProductLevel') AS productlevel,
    SERVERPROPERTY('Edition') AS edition,
    SERVERPROPERTY('EngineEdition') AS engineedition,
    SERVERPROPERTY('MachineName') AS machinename,
    SERVERPROPERTY('IsClustered') AS isclustered,
    SERVERPROPERTY('Collation') AS collation,
    SERVERPROPERTY('BuildClrVersion') AS clrversion,
    SERVERPROPERTY('ProcessorCount') AS processorcount,
    SERVERPROPERTY('PhysicalMemoryMB') AS physicalmemorymb;
`)
	if err != nil {
		log.Printf("Could not get server info: %v", err)
		return info, err
	}

	for result.Next() {
		err = result.StructScan(&info)
	}
	if err != nil {
		log.Printf("Could not get server info: %v", err)
		return info, err
	}
	return info, nil
}
