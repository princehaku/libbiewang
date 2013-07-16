package biewang

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var ChinesNumberMap = map[string]int{
	"零": 0, "一": 1, "二": 2, "两": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
}

var ChinesNumberArr = []string{
	"零", "一", "二", "两", "三", "四", "五", "六", "七", "八", "九", "十", "百", "千", "万", "亿",
}

var CnQualityMap = map[string]int{
	"(.*?)亿(.*)": 100000000, "(.*?)万(.*)": 10000, "(.*?)千(.*)": 1000, "(.*?)百(.*)": 100,
	"(.*?)十(.*)": 10, "(.*)": 1,
}

func pos(arr []string, s string) int {
	for p, v := range arr {
		if v == s {
			return p
		}
	}
	return -1
}
type MapSorter []Item
                 
type Item struct {
    Key string
    Val int
}
                 
func ReMapSorter(m map[string]int) MapSorter {
    ms := make(MapSorter, 0, len(m))
                 
    for k, v := range m {
        ms = append(ms, Item{k, v})
    }
                 
    return ms
}
                 
func (ms MapSorter) Len() int {
    return len(ms)
}
                 
func (ms MapSorter) Less(i, j int) bool {
    return ms[i].Val > ms[j].Val
}
                 
func (ms MapSorter) Swap(i, j int) {
    ms[i], ms[j] = ms[j], ms[i]
}

func ReplaceCnNumber(str string) string {
	r := []rune(str)
	inMap := false
	tmps := ""
	i := 0
	cnnums := map[string] int{}
	for _, v := range r {
		if pos(ChinesNumberArr, string(v)) != -1 {
			inMap = true
			tmps += string(v)
		} else {
			inMap = false
			if tmps != "" {
				cnnums[tmps] = len([]rune(tmps))
				tmps = ""
				i = i + 1
			}
		}
	}
	// 按长度重排大小
	ms := ReMapSorter(cnnums)
    sort.Sort(ms)
	for _,cnnum := range ms {
		cnint := CnStr2Int(cnnum.Key)
		str = strings.Replace(str, cnnum.Key, strconv.Itoa(cnint), -1)
	}
	inMap = inMap && false
	return str
}

func CnStr2Int(cnstr string) int {
	s := -1
	// 十可以默认省略最前面的一
	if strings.HasPrefix(cnstr, "十") {
		cnstr = "一" + cnstr
	}
	for cnregxp, quality := range CnQualityMap {
		var regxpPattern = regexp.MustCompile(cnregxp)
		m := regxpPattern.FindStringSubmatch(cnstr)
		if len(m) > 2 {
			m1 := CnStr2Int(m[1])
			m1 = m1 * quality
			m2 := CnStr2Int(m[2])
			s = m1 + m2
			break
		}
		if len(m) == 2 {
			dst := m[1]
			// 这里的处理是为了规避零九的情况
			runeStr := []rune(m[1])
			if len(runeStr) != 1 {
				i := len(runeStr) - 1
				dst = string(runeStr[i])
			}
			s = ChinesNumberMap[dst]
			break
		}
	}
	return s
}
