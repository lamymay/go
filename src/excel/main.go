package main

import (
	"container/list"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

func main() {

	//testTime()
	//writeRead()
	//writeRead2()
	testGetRows()

	//read2()

}

func testTime() {

	t1 := time.Now().Year()       //年
	t2 := time.Now().Month()      //月
	t3 := time.Now().Day()        //日
	t4 := time.Now().Hour()       //小时
	t5 := time.Now().Minute()     //分钟
	t6 := time.Now().Second()     //秒
	t7 := time.Now().Nanosecond() //
	fmt.Println(time.Now())
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println(t5)
	fmt.Println(t6)
	fmt.Println(t7)

}

//获取到一个list数据，然后输出到excel
func testGetRows() {
	var rows = getList()
	for e := rows.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func getList() *list.List {
	list := list.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	return list
}

func writeRead() {
	//string到int
	//int,err:=strconv.Atoi(string)
	//#string到int64
	//int64, err := strconv.ParseInt(string, 10, 64)
	//#int到string
	//string:=strconv.Itoa(int)

	//int64到string
	var int64 = time.Now().Unix()
	int64Str := strconv.FormatInt(int64, 10)
	var filename = "测试文件" + int64Str + ".xlsx"

	write1(filename)
	read1(filename)
}

func read2() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		println(err.Error())
		return
	}

	println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}

}

func read1(filename string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		println(err.Error())
		return
	}

	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		println(err.Error())
		return
	}

	println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
}

func write1(filename string) {
	wb := excelize.NewFile()
	// 创建一个工作表
	index := wb.NewSheet("Sheet2")
	// 设置单元格的值
	wb.SetCellValue("Sheet2", "A2", "Hello world.")

	wb.SetCellValue("Sheet1", "A1", 1)
	wb.SetCellValue("Sheet1", "B2", 2)
	wb.SetCellValue("Sheet1", "C3", 3)
	wb.SetCellValue("Sheet1", "D4", 4)
	wb.SetCellValue("Sheet1", "E5", 5)

	// 设置工作簿的默认工作表
	wb.SetActiveSheet(index)
	// 根据指定路径保存文件

	if err := wb.SaveAs(filename); err != nil {
		println(err.Error())
	}
}

// 能不能帮忙评估下？？
// 排期调研  外部的大小    产品准备  技术
// 工具 加的  自测用例  代码覆盖率 就是新的代码有跑到没有
//如果是DBA运维那么果断选Ruby（Puppet等自动化运维中的神器全是Ruby）或者Python。如果做Web开发那么PHP、Java或者Ruby。如果是系统级编程以及游戏服务端那么C++、Erlang、Golang。如果是游戏以及Windows相关客户端那么C#。如果跨平台工具软件那么Java或者QT。如果IOS/MAC那么Swift， 安卓继续Java，前端继续JS。如果是黑客脚本小子那么JS、PHP、Python（SqlMap可谓神器）、国内还流行易语言，哈哈哈哈。如果像题
//
//作者：匿名用户
//链接：https://www.zhihu.com/question/36826836/answer/69257537
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
