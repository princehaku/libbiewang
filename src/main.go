package main

import (
	"biewang"
	"time"
)

func main() {
	go biewang.Str2Memo("明天的下午三点 去游戏")
	time.Sleep(1)
}
