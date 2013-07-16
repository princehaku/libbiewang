package libbiewang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var StopWordsArr = []string{
	"个",
	"的",
	"小",
	"以",
	"钟",
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
	second  string
	minute  string
	hour    string
	day     string
	month   string
	year    string
	defined int
}

func parseSecond(str string, pt *TimeMention) {
	var regxpPattern = regexp.MustCompile("(\\d*?)秒(后?)")
	m := regxpPattern.FindStringSubmatch(str)
	if len(m) == 3 {
		if m[2] == "后" {
			pt.second = "+" + m[1]
		}
		if m[2] == "前" {
			pt.second = "-" + m[1]
		}
		if m[2] == "" {
			pt.second = "=" + m[1]
		}
	}
}

func parseMinute(str string, pt *TimeMention) {
	var regxpPattern = regexp.MustCompile("(\\d*?)分(后?)")
	m := regxpPattern.FindStringSubmatch(str)
	if len(m) == 3 {
		if m[2] == "后" {
			pt.minute = "+" + m[1]
		}
		if m[2] == "前" {
			pt.minute = "-" + m[1]
		}
		if m[2] == "" {
			pt.minute = "=" + m[1]
		}
	}
}

func parseHour(str string, pt *TimeMention) {
	var regxpPattern = regexp.MustCompile("(\\d*?)时(后?)")
	m := regxpPattern.FindStringSubmatch(str)
	if len(m) == 3{
		if m[2] == "后" {
			pt.hour = "+" + m[1]
		}
		if m[2] == "前" {
			pt.hour = "-" + m[1]
		}
		if m[2] == "" {
			pt.minute = "+" + m[1]
		}
	}
	regxpPattern = regexp.MustCompile("(\\d*?)点(后?)")
	m = regxpPattern.FindStringSubmatch(str)
	if len(m) == 3 {
		intval, _ := strconv.ParseInt(m[1], 0, 32)
		base_h := int(intval)
		if strings.Contains(str, "下午") {
			base_h = base_h + 12
		}
		pt.hour = "=" + strconv.Itoa(base_h)
	}
}

func parseDay(str string, pt *TimeMention) {

}

func parseWeek(str string, pt *TimeMention) {

}

func parseMonth(str string, pt *TimeMention) {

}

func parseYear(str string, pt *TimeMention) {

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
	parseMinute(str, pTime)
	parseHour(str, pTime)

}
