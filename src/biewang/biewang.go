package biewang

import (
	"fmt"
	"strings"
)

var stop_words_arr = []string{
	"个",
	"的",
	"小",
}

var mapping_times_map = map[string]string{}
var stop_times_map = map[string]string{
	"大前天": "3天前",
	"前天":  "2天前",
	"昨天":  "1天前",
	"半天":  "12小时",
	"明天":  "1天",
	"两天":  "2天",
	"后天":  "2天",
	"大后天": "3天",
	"下周":  "1周后星期",
	"下下周": "2周后",
	"上月":  "1月前",
	"下月":  "1月后",
	"星期":  "周",
	"半":   "30",
	"一刻":  "15",
}

func Str2Memo(str string) {
	for _, w := range stop_words_arr {
		str = strings.Replace(str, w, "", 1)
	}
	fmt.Println(str)
	fmt.Println(CnStr2Int("三亿九千八百一十七万二千一百五十四"))
}
