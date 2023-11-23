package cmd

import (
	"github.com/Aj002Th/imail/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: ServerHandle,
}

func ServerHandle(cmd *cobra.Command, args []string) {
	server.RunMain(ConfigPath)
}
