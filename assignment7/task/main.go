package main

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment7/task/cmd"
	"github.com/sachinnay/Gophercises/assignment7/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		fmt.Println("Error occured :: ", err)
		return
	}
	cmd.RootCmd.Execute()

}
