package dao

import (
	"fmt"
	"log"

	"github.com/sachinnay/Gophercises/Assignemnt7/task/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//DbDao :: Db strct
type DbDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	//COLLECTION :: constant
	COLLECTION = "task"
)

// Connect :: creates connection
func (m *DbDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert a task into database
func (m *DbDao) Insert(task model.Tasks) error {
	log.Println("Insert  :: ")
	err := db.C(COLLECTION).Insert(&task)
	fmt.Println("in Insert", err)
	return err
}

// FindAll tasks from database
func (m *DbDao) FindAll() ([]model.Tasks, error) {
	log.Println("FindAll ::")
	var tasks []model.Tasks
	err := db.C(COLLECTION).Find(bson.M{}).All(&tasks)
	return tasks, err
}

// Delete an existing employee
func (m *DbDao) Delete(task string) error {
	err := db.C(COLLECTION).Remove(bson.M{"task": task})
	return err
}
