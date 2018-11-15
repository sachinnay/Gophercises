package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd :: is definding the root command for CLI task manager
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI task manager",
	Long:  "CLI task manager for managing tasks",
}
