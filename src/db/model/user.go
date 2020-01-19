package model

import (
	"database/sql"
	"fmt"
)

// User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	//1、写sql
	sqlString := "insert into user (username,password,email) values(?,?,?)"
	//2、预编译

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test3?charset=utf8")
	insertStm, err := db.Prepare(sqlString)

	//编译异常
	if err != nil {
		fmt.Println(err)
		return err
	}
	//3、执行
	_, err2 := insertStm.Exec("root2", "ppp", "123qq.com")

	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}

func sayHello() error {
	fmt.Println("sayHello -------------")
	return nil
}

func sayHello2() error {
	fmt.Println("sayHello222 -------------")
	//1、写sql
	sqlString := "insert into user (username,password,email) values(?,?,?)"
	//2、预编译

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test3?charset=utf8")
	insertStm, err := db.Prepare(sqlString)

	//编译异常
	if err != nil {
		fmt.Println(err)
		return err
	}
	//3、执行
	_, err2 := insertStm.Exec("root2", "ppp", "123qq.com")

	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}

func (user *User) GetUserByID(*User, error) {

	selectUserById := "select id,username,email from user where  id =?"

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test3?charset=utf8")

	row := db.QueryRow(selectUserById, user.ID)

	var id int
	var username string
	var password string
	var email string

	err2 := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err2
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}

	return u, nil

}
