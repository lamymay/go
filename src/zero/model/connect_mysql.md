


下载驱动包：

$ go get github.com/go-sql-driver/mysql
最后导入包即可：

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

2.连接至数据库
db, err := sql.Open("mysql", "root:root@/uestcbook")
3.执行查询
（1）Exec 

result, err := db.Exec(
    "INSERT INTO users (name, age) VALUES (?, ?)",
    "gopher",
    27,
)
 

（2）Query 

rows, err := db.Query("SELECT name FROM users WHERE age = ?", age)
if err != nil {
    log.Fatal(err)
}
for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s is %d\n", name, age)
}
if err := rows.Err(); err != nil {
    log.Fatal(err)
}
 

（3）QueryRow

var age int64
row := db.QueryRow("SELECT age FROM users WHERE name = ?", name)
err := row.Scan(&age)
 

 

（4）Prepared statements 

age := 27
stmt, err := db.Prepare("SELECT name FROM users WHERE age = ?")
if err != nil {
    log.Fatal(err)
}
rows, err := stmt.Query(age)
// process rows
 

 

4. 事务
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}
 

 

5. 各种方式效率分析
问题：db.exec和statement.exec和tx.exec的区别？

实例如下：

package main

import (
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "time"
    "log"
)

var db = &sql.DB{}

func init(){
    db,_ = sql.Open("mysql", "root:root@/book")
} 

func main() {
    insert()
    query()
    update()
    query()
    delete()
}

func update(){
    //方式1 update
    start := time.Now()
    for i := 1001;i<=1100;i++{
        db.Exec("UPdate user set age=? where uid=? ",i,i)
    }
    end := time.Now()
    fmt.Println("方式1 update total time:",end.Sub(start).Seconds())
    
    //方式2 update
    start = time.Now()
    for i := 1101;i<=1200;i++{
        stm,_ := db.Prepare("UPdate user set age=? where uid=? ")
        stm.Exec(i,i)
        stm.Close()
    }
    end = time.Now()
    fmt.Println("方式2 update total time:",end.Sub(start).Seconds())
    
    //方式3 update
    start = time.Now()
    stm,_ := db.Prepare("UPdate user set age=? where uid=?")
    for i := 1201;i<=1300;i++{
        stm.Exec(i,i)
    }
    stm.Close()
    end = time.Now()
    fmt.Println("方式3 update total time:",end.Sub(start).Seconds())
    
    //方式4 update
    start = time.Now()
    tx,_ := db.Begin()
    for i := 1301;i<=1400;i++{
        tx.Exec("UPdate user set age=? where uid=?",i,i)
    }
    tx.Commit()
    
    end = time.Now()
    fmt.Println("方式4 update total time:",end.Sub(start).Seconds())
    
    //方式5 update
    start = time.Now()
    for i := 1401;i<=1500;i++{
        tx,_ := db.Begin()
        tx.Exec("UPdate user set age=? where uid=?",i,i)
        tx.Commit()
    }
    end = time.Now()
    fmt.Println("方式5 update total time:",end.Sub(start).Seconds())
    
}

func delete(){
    //方式1 delete
    start := time.Now()
    for i := 1001;i<=1100;i++{
        db.Exec("DELETE FROM USER WHERE uid=?",i)
    }
    end := time.Now()
    fmt.Println("方式1 delete total time:",end.Sub(start).Seconds())
    
    //方式2 delete
    start = time.Now()
    for i := 1101;i<=1200;i++{
        stm,_ := db.Prepare("DELETE FROM USER WHERE uid=?")
        stm.Exec(i)
        stm.Close()
    }
    end = time.Now()
    fmt.Println("方式2 delete total time:",end.Sub(start).Seconds())
    
    //方式3 delete
    start = time.Now()
    stm,_ := db.Prepare("DELETE FROM USER WHERE uid=?")
    for i := 1201;i<=1300;i++{
        stm.Exec(i)
    }
    stm.Close()
    end = time.Now()
    fmt.Println("方式3 delete total time:",end.Sub(start).Seconds())
    
    //方式4 delete
    start = time.Now()
    tx,_ := db.Begin()
    for i := 1301;i<=1400;i++{
        tx.Exec("DELETE FROM USER WHERE uid=?",i)
    }
    tx.Commit()
    
    end = time.Now()
    fmt.Println("方式4 delete total time:",end.Sub(start).Seconds())
    
    //方式5 delete
    start = time.Now()
    for i := 1401;i<=1500;i++{
        tx,_ := db.Begin()
        tx.Exec("DELETE FROM USER WHERE uid=?",i)
        tx.Commit()
    }
    end = time.Now()
    fmt.Println("方式5 delete total time:",end.Sub(start).Seconds())
    
}

func query(){
    
    //方式1 query
    start := time.Now()
    rows,_ := db.Query("SELECT uid,username FROM USER")
    defer rows.Close()
    for rows.Next(){
         var name string
         var id int
        if err := rows.Scan(&id,&name); err != nil {
            log.Fatal(err)
        }
        //fmt.Printf("name:%s ,id:is %d\n", name, id)
    }
    end := time.Now()
    fmt.Println("方式1 query total time:",end.Sub(start).Seconds())
    
    
    
    
    
    
    
    
    
    