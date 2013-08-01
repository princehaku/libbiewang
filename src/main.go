package main

import (
	"libbiewang"
	"fmt"
)

func main() {
	HourTest()
	DayTest()
	WeekTest()
	MonthTest()
	YearTest()
}

func HourTest() {
	pTime := libbiewang.Str2Memo("一小时二十三分钟一百零九十七秒以后")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("三小时二十八分")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("10分钟后")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
}

func DayTest() {
	pTime := libbiewang.Str2Memo("26号下午三点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("后天下午十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("两天前5点一刻")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
}

func WeekTest() {
	pTime := libbiewang.Str2Memo("星期六早上3点")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("上周六晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("下周一晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("星期五晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("星期天晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
}


func MonthTest() {
	pTime := libbiewang.Str2Memo("下个月5号")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("9月23号")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
}

func YearTest() {
	pTime := libbiewang.Str2Memo("2014年晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
	pTime = libbiewang.Str2Memo("明年今天晚上十点半")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
	
}
