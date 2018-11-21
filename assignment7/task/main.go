package main

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment7/task/cmd"
	"github.com/sachinnay/Gophercises/assignment7/task/db"
)

//HomedirVar Using function as variable
var HomedirVar = homedir.Dir

//DbFileName const for DB file name
const dbFileName = "tasks.db"

//Starting point for application
func main() {
	home, err := HomedirVar()
	if err != nil {
		fmt.Println("Error occured :: ", err)
		return
	}
	dbPath := filepath.Join(home, dbFileName)
	err = db.Init(dbPath)
	if err != nil {
		fmt.Println("Error occured :: ", err)
		return
	}
	cmd.RootCmd.Execute()

}
