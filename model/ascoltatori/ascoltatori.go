package ascoltatori

import (
	util "github.com/CARVIN94/go-util"
	"github.com/CARVIN94/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	database   = mgo.GetDatabase()
	collection = "ascoltatori"
)

// Model 数据模型
type Model struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Topic string        `json:"topic"`
	Value string        `json:"value"`
}

// Task 获取连接任务
func Task() *mgo.Task {
	task := mgo.Collection(database, collection)
	return &task
}

// FindAll 查询数据包中全部内容
func FindAll() (rows []Model) {
	task := Task()
	defer task.End()
	err := task.Find(bson.M{}).All(&rows)
	if err != nil {
		util.PanicOnError(err)
	}
	return rows
}

// CountTotal 查询数据包中数据数量
func CountTotal() (total int) {
	task := Task()
	defer task.End()
	total, err := task.Find(bson.M{}).Count()
	if err != nil {
		util.PanicOnError(err)
	}
	return total
}
