package biewang

import (
	"regexp"
	"strings"
)

var chinesnumber_arr_map = map[string]int{
	"零": 0, "一": 1, "二": 2, "两": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
}

var cn_quality_map = map[string]int{
	"(.*?)亿(.*)": 100000000, "(.*?)万(.*)": 10000, "(.*?)千(.*)": 1000, "(.*?)百(.*)": 100,
	"(.*?)十(.*)": 10, "(.*)": 1,
}

func CnStr2Int(cn_str string) int {
	s := -1
	// hack解决方案
	if strings.HasPrefix(cn_str, "十") {
		cn_str = "一" + cn_str
	}
	for cn_regxp, quality := range cn_quality_map {
		var regxpPattern = regexp.MustCompile(cn_regxp)
		m := regxpPattern.FindStringSubmatch(cn_str)
		if len(m) > 2 {
			m1 := CnStr2Int(m[1])
			m1 = m1 * quality
			m2 := CnStr2Int(m[2])
			s = m1 + m2
			break
		}
		if len(m) == 2 {
			s = chinesnumber_arr_map[m[1]]
			break
		}
	}
	return s
}
