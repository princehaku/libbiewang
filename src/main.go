package main

import (
	"libbiewang"
)

func main() {
	HourTest()
	//DayTest()
	//WeekTest()
	//MonthTest()
	//YearTest()
}

func HourTest() {
	libbiewang.Str2Memo("一小时二十三分钟一百零九十七秒以后")
	libbiewang.Str2Memo("三小时二十八分")
	libbiewang.Str2Memo("10分钟后")
}

func DayTest() {
	libbiewang.Str2Memo("26号下午三点半")
	libbiewang.Str2Memo("后天下午十点半")
	libbiewang.Str2Memo("两天前5点一刻")
}

func WeekTest() {
	libbiewang.Str2Memo("星期六早上3点")
	libbiewang.Str2Memo("上周六晚上十点半")
	libbiewang.Str2Memo("下周一晚上十点半")
	libbiewang.Str2Memo("星期五晚上十点半")
	libbiewang.Str2Memo("星期天晚上十点半")
}


func MonthTest() {
	libbiewang.Str2Memo("下个月5号")
	libbiewang.Str2Memo("9月23号")
}

func YearTest() {
	libbiewang.Str2Memo("2014年晚上十点半")
	libbiewang.Str2Memo("明年今天晚上十点半")
}
