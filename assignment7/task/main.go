package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment7/task/cmd"
	"github.com/sachinnay/Gophercises/assignment7/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
