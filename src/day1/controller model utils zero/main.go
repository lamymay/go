package main

import (
	"fmt"
)

func main() {
	defer_call()
}

// defer 的执行顺序是倒序执行（同入栈先进后出） 栈结构，FILO
//1.defer的生效顺序

//2.defer与return,函数返回值之间的顺序		return最先执行->return负责将结果写入返回值中->接着defer开始执行一些收尾工作->最后函数携带当前返回值退出
//
//返回值的表达方式，我们知道根据是否提前声明有两种方式：一种是func controller model utils zero() int 另一种是 func controller model utils zero() (i int)，所以两种情况都来说说

//3.defer定义和执行两个步骤，做的事情。
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
