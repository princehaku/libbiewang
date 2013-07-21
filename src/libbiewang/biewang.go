package libbiewang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	"两天":  "2天",
	"后天":  "2天后",
	"大后天": "3天后",
	"上周":  "1周前周",
	"下周":  "1周后周",
	"下下周": "2周后",
	"上月":  "1月前",
	"下月":  "1月后",
	"周日" : "周7",
	"星期天" : "周7",
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

func (t *TimeMention) String() string {
	s := " Minute:" + t.minute + " Second:" + t.second
	s += " Day:" + t.day + " Hour:" + t.hour
	s += " Year:" + t.year + " Month:" + t.month

	return s
}
func MatchAndReplace(str string, pattern string) (string, []string) {
	var regxpPattern = regexp.MustCompile(pattern)
	m := regxpPattern.FindStringSubmatch(str)
	if len(m) > 1 {
		str = strings.Replace(str, m[0], "", 1)
		fmt.Println(str)
	}
	return str, m
}
func ParseSecond(str string, pt *TimeMention) string {
	str, m := MatchAndReplace(str, "(\\d+)秒([前后]?)")
	if len(m) >= 3 {
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
	return str
}

func ParseMinute(str string, pt *TimeMention) string {
	str, m := MatchAndReplace(str, "(\\d+)分([前后]?)")
	if len(m) >= 3 {
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
	return str
}

func ParseHour(str string, pt *TimeMention) string {
	str, m := MatchAndReplace(str, "(\\d+)时([前后]?)")
	if len(m) >= 3 {
		if m[2] == "后" {
			pt.hour = "+" + m[1]
		}
		if m[2] == "前" {
			pt.hour = "-" + m[1]
		}
		if m[2] == "" {
			pt.hour = "+" + m[1]
		}
	}
	str, m = MatchAndReplace(str, "(\\d+)点")
	if len(m) >= 2 {
		intval, _ := strconv.ParseInt(m[1], 0, 32)
		base_h := int(intval)
		if strings.Contains(str, "早上") {
			str = strings.Replace(str, "早上", "", 1)
		}
		if strings.Contains(str, "下午") {
			base_h = base_h + 12
			str = strings.Replace(str, "下午", "", 1)
		}
		if strings.Contains(str, "晚上") {
			base_h = base_h + 12
			str = strings.Replace(str, "晚上", "", 1)
		}
		pt.hour = "=" + strconv.Itoa(base_h)
	}
	return str
}

func ParseDay(str string, pt *TimeMention) string {
	str, m := MatchAndReplace(str, "(\\d+)天([前后]?)")
	if len(m) >= 3 {
		if m[2] == "后" {
			pt.day = "+" + m[1]
		}
		if m[2] == "前" {
			pt.day = "-" + m[1]
		}
		if m[2] == "" {
			pt.day = "+" + m[1]
		}
	}
	str, m = MatchAndReplace(str, "(\\d+)号")
	if len(m) == 2 {
		pt.day = "=" + m[1]
	}
	return str
}

func ParseWeek(str string, pt *TimeMention) string {
	//N周
	str, m := MatchAndReplace(str, "(\\d+)周([前后]?)")
	if len(m) >= 3 {
		if m[2] == "后" {
			temp_i, _ := strconv.Atoi(m[1])
			pt.day = "+" + strconv.Itoa(temp_i*7)
		}
		if m[2] == "前" {
			temp_i, _ := strconv.Atoi(m[1])
			pt.day = "-" + strconv.Itoa(temp_i*7)
		}
		if m[2] == "" {
			temp_i, _ := strconv.Atoi(m[1])
			pt.day = "=" + strconv.Itoa(temp_i*7)
		}
	}
	// 周1-六
	str, m = MatchAndReplace(str, "周(\\d+)")
	if len(m) >= 2 {
		temp_i, _ := strconv.Atoi(m[1])
		temp_i += 0
		t := time.Now()
		// 需要把weekday的0转成7
		weekday := int(t.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		t_new := temp_i - weekday
		if pt.day != "" {
			temp_i, _ := strconv.Atoi(pt.day)
			pt.day = strconv.Itoa(temp_i + t_new)
		} else {
			pt.day = strconv.Itoa(t_new)
		}
	}

	return str
}

func parseMonth(str string, pt *TimeMention) {

}

func parseYear(str string, pt *TimeMention) {

}

func ReplaceEnTime(str string) string {

	return str
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
	str = ReplaceEnTime(str)

	// 然后从秒开始..挨个去解析
	fmt.Println(str)
	pTime := new(TimeMention)
	//pDuratime := new(DurationMention)
	str = ParseSecond(str, pTime)
	str = ParseMinute(str, pTime)
	str = ParseHour(str, pTime)
	str = ParseDay(str, pTime)
	str = ParseWeek(str, pTime)
	fmt.Println(pTime)
}
