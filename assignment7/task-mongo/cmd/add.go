package cmd

import (
	"fmt"
	"strings"

	"github.com/sachinnay/Gophercises/Assignemnt7/task/model"

	"github.com/sachinnay/Gophercises/Assignemnt7/task/service"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add the task in list",
	Run: func(cmd *cobra.Command, args []string) {
		taskInput := strings.Join(args, " ")
		var task model.Tasks
		task.Task = taskInput
		// Do Stuff Here
		service.Add(task)

		fmt.Println("Task added in the list ===>", taskInput)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
