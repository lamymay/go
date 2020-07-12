package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(time.Now())
	fmt.Fprint(w, time.Now(), "  Hello word !", r.URL.Path)

}

//注意main 方法所在的包要是main包
//fmt.Print("day2_file word!")
