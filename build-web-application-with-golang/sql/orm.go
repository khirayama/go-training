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

	var saveone Userinfo

	// INSERT
	// standard
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

	// UPDATE
	saveone.Username = "Update Username"
	saveone.Departname = "Update Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone) //現在saveoneにはプライマリキーがあります。更新操作を行います。

	t := make(map[string]interface{})
	t["username"] = "astaxie"
	orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)

	// Search
	var user Userinfo
	orm.Where("uid=?", 27).Find(&user)

	var user2 Userinfo
	orm.Where(3).Find(&user2)

	var user3 Userinfo
	orm.Where("name=?", "john").Find(&user3)

	var user4 Userinfo
	orm.Where("name=? and age < ?", "john", 88).Find(&user4)

	var allusers []Userinfo
	err = orm.Where("id > ?", "3").Limit(10, 20).FindAll(&allusers)

	var tenusers []Userinfo
	err = orm.Where("id > ?", "3").Limit(10).FindAll(&tenusers)

	var everyone []Userinfo
	err = orm.Where("uid desc, username asc").FindAll(&everyone)

	// DELETE
	orm.Delete(&saveone)

	orm.DeleteAll(&allusers)

	orm.SetTable("userinfo").Where("uid>?", 3).DeleteRow()

	// RELATION
	a, _ := orm.SetTable("userinfo").Join("LEFT", "userdeatail", "userinfo.uid=userdeatail.uid").Where("userinfo.uid=?", 1).Select("userinfo.uid,userinfo.username,userdeatail.profile").FindMap()

	// GROUP BY / HAVING
	a, _ = orm.SetTable("userinfo").GroupBy("username").Having("username='astaxie'").FindMap()
}
