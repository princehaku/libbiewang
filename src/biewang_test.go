package libbiewang_test

import (
	"github.com/princehaku/libbiewang"
	"testing"
	"fmt"
)

func TestStr2Memo(t *testing.T) {
	pTime := libbiewang.Str2Memo("一个月三天一小时二十三分钟一百零九十七秒以后")
	fmt.Println(pTime)
	fmt.Println(pTime.Time())
}

func TestCNumber2Int(t *testing.T) {
	 c:= libbiewang.CnStr2Int("三亿九千八百一十七万二千一百五十四")
	 fmt.Println(c)
}

func BenchmarkStr2Memo(b *testing.B) {
    for i := 0; i < b.N; i++ {
	libbiewang.Str2Memo("一个月三天一小时二十三分钟一百零九十七秒以后")
    }
}

func BenchmarkCNumber2Int(b *testing.B) {
    for i := 0; i < b.N; i++ {
		libbiewang.CnStr2Int("三亿九千八百一十七万二千一百五十四")
    }
}