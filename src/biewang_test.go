package biewang_test

import (
	"biewang"
	"testing"
	"fmt"
)

func TestMemo(t *testing.T) {
	biewang.Str2Memo("明天下午三点")
}

func TestCNumber2Int(t *testing.T) {
	 //c:= biewang.CnStr2Int("三亿九千八百一十七万二千一百五十四")
}

func BenchmarkCNumber2Int(b *testing.B) {
    for i := 0; i < b.N; i++ {
		biewang.CnStr2Int("三亿九千八百一十七万二千一百五十四")
    }
}