package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

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
			}
		}
		fmt.Println("do called", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
