package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

func Test35(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{1, 3, 5, 6}, 5},
			ans{2},
		},

		{
			para{[]int{1, 3, 5, 6}, 2},
			ans{1},
		},

		{
			para{[]int{1, 3, 5, 6}, 7},
			ans{4},
		},

		{
			para{[]int{1, 3, 5, 6}, 0},
			ans{0},
		},

		{
			para{[]int{1, 3, 5, 6}, 6},
			ans{3},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, searchInsert(p.one, p.two), "输入:%v", p)
	}
}
