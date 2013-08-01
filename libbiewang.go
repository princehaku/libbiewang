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
	"今天",
}

var MappingTimesMap = map[string]string{}

var FormatTimesMap = [][]string{
	[]string{"大前天", "3天前"},
	[]string{"前天", "2天前"},
	[]string{"昨天", "1天前"},
	[]string{"半天", "12小时后"},
	[]string{"明天", "1天后"},
	[]string{"两天", "2天"},
	[]string{"后天", "2天后"},
	[]string{"大后天", "3天后"},
	[]string{"上周", "1周前周"},
	[]string{"下周", "1周后周"},
	[]string{"下下周", "2周后"},
	[]string{"周日", "周7"},
	[]string{"星期天", "周7"},
	[]string{"星期", "周"},
	[]string{"上月", "1月前"},
	[]string{"下月", "1月后"},
	[]string{"去年", "1年前"},
	[]string{"明年", "1年后"},
	[]string{"半", "30分"},
	[]string{"一刻", "15分"},
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
	s += " Hour:" + t.hour + " Day:" + t.day
	s += " Month:" + t.month + " Year:" + t.year
	return s
}

// 转换成系统可用的Time对象
func (t *TimeMention) Time() time.Time {
	now := time.Now()
	sec := now.Second()
	min := now.Minute()
	hour := now.Hour()
	day := now.Day()
	month := now.Month()
	month_i := int(month)
	year := now.Year()
	duration := 0
	duration += 0
	lastopt := ""
	sec, lastopt, duration = ReParserTime(t.second, lastopt, 1, sec, duration)
	min, lastopt, duration = ReParserTime(t.minute, lastopt, 60, min, duration)
	hour, lastopt, duration = ReParserTime(t.hour, lastopt, 3600, hour, duration)

	duration_day := 0
	duration_month := 0
	duration_year := 0
	day, lastopt, duration_day = ReParserTime(t.day, "", 1, day, duration_day)
	month_i, lastopt, duration_month = ReParserTime(t.month, "", 1, month_i, duration_month)
	year, lastopt, duration_year = ReParserTime(t.year, "", 1, year, duration_year)

	tc := time.Date(year, time.Month(month_i), day, hour, min, sec, 0, time.UTC)
	tc = tc.Add(time.Duration(duration) * time.Second)
	tc = tc.AddDate(duration_year, duration_month, duration_day)
	return tc
}

func ReParserTime(timeSpec string, lastopt string, square int, resignment int, duration int) (int, string, int) {
	opt, qua := SplitTime(timeSpec)

	switch opt {
	case "+":
		duration += qua * square
		break
	case "-":
		duration -= qua * square
		break
	case "=":
		if lastopt == "+" {
			opt = "+"
			duration += qua * square
		} else {
			resignment = qua
		}
		break
	}
	return resignment, opt, duration
}

func SplitTime(str string) (string, int) {
	opt := "@"
	qua := 0
	if len(str) > 0 {
		opt = strings.Split(str, "")[0]
		qua_s := strings.TrimLeft(str, opt)
		qua, _ = strconv.Atoi(qua_s)
	}
	return opt, qua
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
		if pt.day != "0" && !strings.HasSuffix(pt.day, "+") {
			pt.day = "+" + pt.day
		}
	}

	return str
}

func ParseMonth(str string, pt *TimeMention) string {
	//N月
	str, m := MatchAndReplace(str, "(\\d+)月([前后]?)")
	if len(m) >= 3 {
		if m[2] == "后" {
			pt.month = "+" + m[1]
		}
		if m[2] == "前" {
			pt.month = "-" + m[1]
		}
		if m[2] == "" {
			pt.month = "=" + m[1]
		}
	}

	return str
}

func ParseYear(str string, pt *TimeMention) string {
	//N年
	str, m := MatchAndReplace(str, "(\\d+)年([前后]?)")
	if len(m) >= 3 {
		if m[2] == "后" {
			pt.year = "+" + m[1]
		}
		if m[2] == "前" {
			pt.year = "-" + m[1]
		}
		if m[2] == "" {
			pt.year = "=" + m[1]
		}
	}

	return str
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
	for _, t := range FormatTimesMap {
		str = strings.Replace(str, t[0], t[1], -1)
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
	str = ParseMonth(str, pTime)
	str = ParseYear(str, pTime)
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
}
