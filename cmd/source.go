package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sourceCmd)
}

var sourceCmd = &cobra.Command{
	Use: "server",
	Run: sourceHandle,
}

func sourceHandle(cmd *cobra.Command, args []string) {
	fmt.Println("imail source manager")
}
