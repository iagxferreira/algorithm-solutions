package hackerrank_test

import (
	"hackerrank"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	root := &hackerrank.Node{Data: 1}
	root.Right = &hackerrank.Node{Data: 2}
	root.Right.Right = &hackerrank.Node{Data: 5}
	root.Right.Right.Left = &hackerrank.Node{Data: 3}
	root.Right.Right.Right = &hackerrank.Node{Data: 6}
	root.Right.Right.Left.Right = &hackerrank.Node{Data: 4}

	order := hackerrank.LevelOrder(root) // Output: 1 2 5 3 6 4

	assert.Equal(t, order, []int{1, 2, 5, 3, 6, 4})
}
