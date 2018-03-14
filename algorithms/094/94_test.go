package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	pre []int
	in  []int
}

type ans struct {
	one []int
}

func Test94(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{},

		{
			para{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
			},
			ans{
				[]int{1, 3, 2},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, inorderTraversal(algorithms.PreIn2Tree(p.pre, p.in)), "输入:%v", p)
	}
}
