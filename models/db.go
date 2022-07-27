package models

import (
	"log"
	"github.com/globalsign/mgo"
)

const (
	host   = "localhost:27017"
	source = ""
	user   = ""
	pass   = ""

	db = "blog"
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   source,
		Username: user,
		Password: pass,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalS = s
}

// db:操作的数据库
// collection:操作的文档(表)
// query:查询条件
// selector:需要过滤的数据(projection)
// result:查询到的结果

// 连接数据库
func connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

// // 插入一条数据
func Insert(collection string, docs ...interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

// 查询是否存在
func IsExist(collection string, query interface{}) bool {
	ms, c := connect(collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func FindOne(collection string, query, selector, result interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(collection string, query, selector, result interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(collection string, query, update interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(collection string, query interface{}) error {
	ms, c := connect(collection)
	defer ms.Close()
	return c.Remove(query)
}
