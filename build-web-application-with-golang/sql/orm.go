package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "test/xiemengjun/123456")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)

	type Userinfo struct {
		Uid        int `PK`
		Username   string
		Departname string
		Created    time.Time
	}

	// standard
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)

	// map
	add := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2012-12-02"
	orm.SetTable("userinfo").Insert(add)

	// insert multi data
	addslice := make([]map[string]interface{}, 0)
	add = make(map[string]interface{})
	add2 := make(map[string]interface{})
	add["username"] = "astaxie"
	add["departname"] = "cloud develop"
	add["created"] = "2012-12-02"
	add2["username"] = "astaxie2"
	add2["departname"] = "cloud develop2"
	add2["created"] = "2012-12-02"
	addslice = append(addslice, add, add2)
	orm.SetTable("userinfo").InsertBatch(addslice)
}
