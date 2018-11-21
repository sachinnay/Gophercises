package cmd

import (
	"fmt"

	"github.com/sachinnay/Gophercises/assignment7/task/db"
	"github.com/spf13/cobra"
)

//CLI list command used to list down the tasks
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "To list down the tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error occured  ::", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks are available in list")
			return
		}
		fmt.Println("You have following tasks: ")
		for i, task := range tasks {
			fmt.Printf("%d %s\n", i+1, task.Value)

		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
