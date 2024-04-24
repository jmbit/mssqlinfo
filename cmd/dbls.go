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
	Use:   "dbls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
