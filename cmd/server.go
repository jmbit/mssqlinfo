/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/jmbit/mssqlinfo/internal/inout"
	"github.com/jmbit/mssqlinfo/internal/mssql"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "General information about the SQL server",
	Long: `Provides general info for MSSQL Servers:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		inout.Ask()
		db, err := mssql.Connect()
		if err != nil {
			log.Fatal(err)
		}
		info, err := mssql.GetServerInfo(db)
		if err != nil {
			log.Panic(err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "MSSQL Server\t%s\n", info.MachineName)
		fmt.Fprintf(w, "Version\t%s\n", info.Version)
		fmt.Fprintf(w, "ProductLevel\t%s\n", info.ProductLevel)
		fmt.Fprintf(w, "Edition\t%s\n", info.Edition)
		fmt.Fprintf(w, "EngineEdition\t%s\n", info.EngineEdition)
		fmt.Fprintf(w, "Clustered\t%v\n", info.IsClustered)
		fmt.Fprintf(w, "Default Collation\t%s\n", info.Collation)
		fmt.Fprintf(w, "BuildCLR Version\t%s\n", info.CLRVersion)
		fmt.Fprintf(w, "Processors\t%d\n", info.ProcessorCount.Int32)
		fmt.Fprintf(w, "RAM (MB)\t%d\n", info.PhysicalMemoryMB.Int32)
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
