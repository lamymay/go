package utils

import "database/sql"

var (
	Db  *sql.DB
	err error
)

func init() {
	//Db,err:=sql.Open("mysql","root@tcp(127.0.0.1:3306/dev)")
	Db, err := sql.Open("mysql", "root@tcp(122.51.110.127:3306/zero)")
	//lamymay12345677Z!
	if err != nil {
		panic(err.Error())
	}
}
