package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type TestEn struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	id := "fff"
	println(id)
	dsn := "root:G2Y2LKXpig@tcp(127.0.0.1:3306)/rw_takeout"
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err == nil {
		fmt.Println("Db err:", err)
		return
	}
	fmt.Println("数据库链接成功")
}
func AddTest(t *TestEn) error {
	_, err := Db.Exec("insert into test(id, name) values (?, ?)", t.Id, t.Name)
	if err != nil {
		return err
	}
	return nil
}
func DelTest(t *TestEn) error {
	_, err := Db.Exec("delete from test where id = ?", t.Id)
	if err != nil {
		return err
	}
	return nil
}
