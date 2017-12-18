package Problem0567

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s1, s2 string
	ans    bool
}{

	{
		"ab",
		"ab",
		true,
	},

	{
		"ab",
		"eidbaooo",
		true,
	},

	{
		"ab",
		"eidboaoo",
		false,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkInclusion(tc.s1, tc.s2), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkInclusion(tc.s1, tc.s2)
		}
	}
}
