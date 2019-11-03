package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Zaruba",
	Long:  `All software has versions. This is Zaruba's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Zaruba v0.0.0 -- [prototype]")
	},
}