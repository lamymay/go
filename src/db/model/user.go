package model

import(
	"db/utils"
	"fmt"
)

// User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// insert
func (iser *User) saveUser() error {
	sqlString := "insert into user (username,`password`,email) values(?,?,?);"
	fmt.Print(sqlString)

	rows,err := utils.Db.Query("select id,username from user_info");
	if err != nil{
		fmt.Printf("select fail [%s]",err)
	}

	var mapUser map[string]int
	mapUser = make(map[string]int)

	for rows.Next(){
		var id int
		var username string
		rows.Columns()
		err := rows.Scan(&id,&username)
		if err != nil{
			fmt.Printf("get user info error [%s]",err)
		}
		mapUser[username] = id
	}

	for k,v := range mapUser{
		fmt.Println(k,v);
	}

}
