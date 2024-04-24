/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jmbit/mssqlinfo/internal/inout"
	"github.com/jmbit/mssqlinfo/internal/mssql"
	"github.com/spf13/cobra"
)

// dblsCmd represents the dbls command
var dblsCmd = &cobra.Command{
	Use: "dbls",

	Short: "List the Databases on the Server",
	Long: `List Databases for a MSSQL server:
mssqlinfo dbls --user sa --password verysecretpassword --server mssql.example.com
Name         Owner        Size(MB)  Recovery
master       sa           4         SIMPLE
testusteron  testusteron  8         SIMPLE`,
	Run: func(cmd *cobra.Command, args []string) {
		inout.Ask()
		db := mssql.Connect()
		results := mssql.ListDatabases(db)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "Name\tOwner\tSize(MB)\tRecModel")

		for _, p := range results {
			fmt.Fprintf(w, "%s\t%s\t%d\t%s", p.Name, p.Owner, p.Size, p.RecoveryModel)

		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(dblsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dblsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dblsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
