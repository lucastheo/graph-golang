package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListIsTree(t *testing.T) {
	s := 100
	g := testNewListTree(s)

	assert.True(t, IsTree(g))
	for i := 1; i < s-1; i++ {
		assert.True(t, IsTree(g.AddEdge(testNewListEdge(0, i, s))))
	}

	g = testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(testNewListNode(i))
	}

	for i := 0; i < s-1; i++ {
		g = g.AddEdge(testNewListEdge(i+1, i, i))
	}

	assert.True(t, IsTree(g))
	assert.False(t, IsTree(g.AddEdge(testNewListEdge(0, s-1, s))))
}
