package cmd

import (
	"fmt"
	"strconv"

	"github.com/sachinnay/Gophercises/assignment7/task/db"
	"github.com/spf13/cobra"
)

//CLI do command used to mark the task as completed
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "To mark task as completed",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, input := range args {
			id, err := strconv.Atoi(input)
			if err == nil {
				ids = append(ids, id)
			} else {
				fmt.Println("Error while parsing inputs :: ", id)
				return

			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error occured  ::", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number ::", id)
				return
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark task: \"%d\" as completed. Error is :: %s  \n ", id, err)
				return
			}
			fmt.Printf("Marked task :: \"%d\"  as completed. \n", id)
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
