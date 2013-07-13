package main

import (
	"fmt"
	"time"
	"biewang"
)
			
func server(i int) {
	for {
		fmt.Println(i)
		time.Sleep(1)
	}
}

func main() {
	go server(1)
	go server(2)
	biewang.Str2Todo("明天下午三点 去游戏")
}
