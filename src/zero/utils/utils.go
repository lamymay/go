package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test3?charset=utf8")
	//	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test3")
	//Db, err := sql.Open("mysql", "root@tcp(122.51.110.127:3306/zero)")
	//lamymay12345677Z!
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Db)
	//if err != nil{
	//	fmt.Printf("connect mysql fail ! [%s]",err)
	//}else{
	//	fmt.Println("connect to mysql success")
	//}

}
