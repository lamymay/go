package model

import (
	"fmt"
	"testing"
)

//run package tests | run file tests

/*//run controller model utils zero |debug controller model utils zero
func TestSaveUser(t *testing.T) {

	fmt.Println("\n ")
	fmt.Println(" ########## START ########## ")
	usr :=   &User {}
	fmt.Println(usr)
	//
	//调用User中的方法
	usr.SaveUser()
	fmt.Println(" ########## END ########## ")

}
*/

func TestUser(t *testing.T) {
	fmt.Println("###### main controller model utils zero")

	t.Run("ss", testSaveUser1111)
	t.Run("ss", testSaveUser2222)
	t.Run("testSaveUser", testSaveUser)

}

func testSaveUser1111(t *testing.T) {
	fmt.Println(" ########## 1111 ########## ")
}

func testSaveUser2222(t *testing.T) {
	fmt.Println(" ########## 2222 ########## ")
}

func testSaveUser(t *testing.T) {
	fmt.Println(" ########## usr.AddUser() ########## ")
	sayHello()

	//报错  runtime error: invalid memory address or nil pointer dereference [recovered]
	//        panic: runtime error: invalid memory address or nil pointer dereference
	//[signal 0xc0000005 code=0x1 addr=0x20 pc=0x5133a8]
	//
	//goroutine 10 [running]:
	//sayHello2()
}

func TestGetUserByID(t *testing.T) {
	fmt.Println("\n ")
	fmt.Println(" ########## START  selectOneById ########## ")
	//var usr = &User{}
	var usr = User{
		ID: 1,
	}
	fmt.Println(usr)
	//
	//调用User中的方法
	//user.saveUser()

	//u, _ := usr.GetUserByID()
	//fmt.Println(u)

	fmt.Println(" ########## END ########## ")

}
