package service

import (
	"log"

	"github.com/sachinnay/Gophercises/Assignemnt7/task/dao"
	"github.com/sachinnay/Gophercises/Assignemnt7/task/model"
)

var dbo = dao.DbDao{
	Server:   "localhost",
	Database: "testdb",
}

func init() {
	dbo.Connect()
}

//Add :: To call db method to add the task in the list
func Add(task model.Tasks) {
	log.Println("Controller :: Add")
	err := dbo.Insert(task)
	if err != nil {
		log.Println("Error occurred :: ", err)
	}
}

//List :: To call db method to listdown tasks
func List() []model.Tasks {
	log.Println("Controller :: List")

	tasks, err := dbo.FindAll()
	if err != nil {
		log.Println("Error occurred :: ", err)
	}
	return tasks
}

//Do :: To call db method to mark the task as completed
func Do() []model.Tasks {
	log.Println("Controller :: Do")

	tasks, err := dbo.FindAll()
	if err != nil {
		log.Println("Error occurred :: ", err)
	}
	return tasks
}
