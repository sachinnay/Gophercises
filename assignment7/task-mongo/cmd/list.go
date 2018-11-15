package cmd

import (
	"fmt"

	"github.com/sachinnay/Gophercises/Assignemnt7/task/service"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "To listdown the tasks",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		tasks := service.List()
		for i, taskItem := range tasks {
			fmt.Printf("%d  %s  \n", i+1, taskItem.Task)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
