package function

import (
	"graph/src/graph"
	"graph/src/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewListGraphEmpty(v int) *graph.ListGraph[util.TestInterface, util.TestInterface, util.TestInterface] {
	return graph.NewEmpty[util.TestInterface, util.TestInterface, util.TestInterface](
		util.TestInterface{
			Value: v,
		},
	)
}

func testNewListNode(v int) graph.ListNode[util.TestInterface] {
	return graph.ListNode[util.TestInterface]{
		Attributes: util.TestInterface{
			Value: v,
		},
	}
}

func testNewListEdge(s, d, v int) graph.ListEdge[util.TestInterface] {
	return graph.ListEdge[util.TestInterface]{
		Source:     s,
		Destine:    d,
		Attributes: util.TestInterface{Value: 0},
	}
}

func testNewListTree(s int) *graph.ListGraph[util.TestInterface, util.TestInterface, util.TestInterface] {
	g := testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(testNewListNode(i))
	}

	for i := 0; i < s-1; i++ {
		g = g.AddEdge(testNewListEdge(i/2, i+1, i))
	}
	return g
}

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
