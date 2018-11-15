package cmd

import (
	"fmt"
	"strings"

	"github.com/sachinnay/Gophercises/assignment7/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add the task in list",
	Run: func(cmd *cobra.Command, args []string) {
		taskInput := strings.Join(args, " ")
		_, err := db.CreateTask(taskInput)
		if err != nil {
			fmt.Println("Error occured ", err.Error())
			return
		}
		fmt.Println("Task added in the list ===>", taskInput)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
