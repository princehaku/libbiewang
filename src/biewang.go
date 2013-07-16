package biewang

import (
	"fmt"
	"regexp"
	"strings"
)

var StopWordsArr = []string{
	"个",
	"的",
	"小",
}

var MappingTimesMap = map[string]string{}

var FormatTimesMap = map[string]string{
	"大前天": "3天前",
	"前天":  "2天前",
	"昨天":  "1天前",
	"半天":  "12小时后",
	"明天":  "1天后",
	"两天":  "2天后",
	"后天":  "2天后",
	"大后天": "3天后",
	"下周":  "1周后星期",
	"下下周": "2周后",
	"上月":  "1月前",
	"下月":  "1月后",
	"星期":  "周",
	"半":   "30分",
	"一刻":  "15分",
}

type TimeMention struct {
	second  int
	minute  int
	hour    int
	day     int
	month   int
	year    int
	defined int
}

type DurationMention struct {
	second   int
	minute   int
	hour     int
	day      int
	month    int
	year     int
	timeType int
}

func parseSecond(str string, pt *TimeMention) {
	var regxpPattern = regexp.MustCompile("(\\d*?)秒")
	m := regxpPattern.FindStringSubmatch(str)
	fmt.Println(m)
}

func parseMinute(str string, t *TimeMention) {
	var regxpPattern = regexp.MustCompile("秒")
	m := regxpPattern.FindStringSubmatch(str)
	fmt.Println(m)

}

func parseHour(str string, t *TimeMention) {

}

func parseDay(str string, t *TimeMention) {

}

func parseWeek(str string, t *TimeMention) {

}

func parseMonth(str string, t *TimeMention) {

}

func parseYear(str string, t *TimeMention) {

}

func Str2Memo(str string) {
	// 清除一些终止词
	for _, w := range StopWordsArr {
		str = strings.Replace(str, w, "", -1)
	}
	// 转义一些中文词，规整化
	for f, t := range FormatTimesMap {
		str = strings.Replace(str, f, t, -1)
	}
	// 把中文描述的数字全部转成英文
	str = ReplaceCnNumber(str)

	// 然后从秒开始..挨个去解析
	fmt.Println(str)
	pTime := new(TimeMention)
	//pDuratime := new(DurationMention)
	parseSecond(str, pTime)

}
