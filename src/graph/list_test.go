package graph

import (
	"fmt"
	"graph/src/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewListGraphEmpty(v int) *ListGraph[util.TestInterface, util.TestInterface, util.TestInterface] {
	return NewEmpty[util.TestInterface, util.TestInterface, util.TestInterface](
		util.TestInterface{
			Value: v,
		},
	)
}

func testNewListNode(v int) ListNode[util.TestInterface] {
	return ListNode[util.TestInterface]{
		Attributes: util.TestInterface{
			Value: v,
		},
	}
}

func testNewListEdge(s, d, v int) ListEdge[util.TestInterface] {
	return ListEdge[util.TestInterface]{
		Source:     s,
		Destine:    d,
		Attributes: util.TestInterface{Value: 0},
	}
}

func testNewListTree(s int) *ListGraph[util.TestInterface, util.TestInterface, util.TestInterface] {
	g := testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(testNewListNode(i))
	}

	for i := 0; i < s-1; i++ {
		g = g.AddEdge(testNewListEdge(i/2, i+1, i))
	}
	return g
}

func TestListCreateNewEmptyWithAttributes(t *testing.T) {
	g := testNewListGraphEmpty(1)

	assert.Equal(t, 1, g.Attributes.Value)
}

func TestListAddNode(t *testing.T) {
	v := 3
	g := testNewListGraphEmpty(1)
	ng := g.AddNode(ListNode[util.TestInterface]{
		Attributes: util.TestInterface{Value: v},
	})

	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.nodes)+1, len(ng.nodes))
	assert.Equal(t, len(g.edges), len(ng.edges))
	assert.Equal(t, ng.nodes[0].Attributes.Value, v)
}

func TestListAddEdge(t *testing.T) {
	sId := 5
	dId := 6
	v := 7
	g := testNewListGraphEmpty(1)
	ng := g.AddEdge(ListEdge[util.TestInterface]{
		Source:     sId,
		Destine:    dId,
		Attributes: util.TestInterface{Value: v},
	})

	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.nodes), len(ng.nodes))
	assert.Equal(t, len(g.edges)+1, len(ng.edges))
	assert.Equal(t, ng.edges[0].Source, sId)
	assert.Equal(t, ng.edges[0].Destine, dId)
	assert.Equal(t, ng.edges[0].Attributes.Value, v)
}

func TestListRemoveEdge(t *testing.T) {
	g := testNewListGraphEmpty(1)
	for i := 0; i < 10; i++ {
		g = g.AddEdge(
			ListEdge[util.TestInterface]{
				Source:     i,
				Destine:    i + 1,
				Attributes: util.TestInterface{Value: i},
			},
		)
	}

	ng, err := g.RemoveEdge(1)
	assert.Nil(t, err)
	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.edges)-1, len(ng.edges))
	for _, e := range ng.edges {
		assert.NotEqual(t, 1, e.Attributes.Value)
	}

	ng, err = g.RemoveEdge(0)
	assert.Nil(t, err)
	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.edges)-1, len(ng.edges))
	for _, e := range ng.edges {
		assert.NotEqual(t, 0, e.Attributes.Value)
	}

	ng, err = g.RemoveEdge(len(g.edges) - 1)
	assert.Nil(t, err)
	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.edges)-1, len(ng.edges))
	for _, e := range ng.edges {
		assert.NotEqual(t, len(g.edges)-1, e.Attributes.Value)
	}

	ng, err = g.RemoveEdge(len(ng.edges) - 1)
	assert.Nil(t, err)
	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.edges)-1, len(ng.edges))
	for _, e := range ng.edges {
		assert.NotEqual(t, len(ng.edges)-1, e.Attributes.Value)
	}

	ng, err = g.RemoveEdge(0)
	assert.Nil(t, err)
	assert.NotEqual(t, g, ng)
	assert.Equal(t, len(g.edges)-1, len(ng.edges))
	for _, e := range ng.edges {
		assert.NotEqual(t, 0, e.Attributes.Value)
	}

	_, err = g.RemoveEdge(len(g.edges))
	assert.Nil(t, err)
}

func TestListRemoveNode(t *testing.T) {
	// Create a circular graph
	s := 30
	g := testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(ListNode[util.TestInterface]{
			Attributes: util.TestInterface{Value: i},
		})
	}
	for i := 0; i < s-1; i++ {
		g = g.AddEdge(
			ListEdge[util.TestInterface]{
				Source:     i,
				Destine:    i + 1,
				Attributes: util.TestInterface{Value: i},
			},
		)
	}
	g = g.AddEdge(
		ListEdge[util.TestInterface]{
			Source:     len(g.nodes) - 1,
			Destine:    0,
			Attributes: util.TestInterface{Value: len(g.nodes) - 1},
		},
	)

	for id := range g.nodes {
		g1 := g.RemoveNode(id)
		errOutput := fmt.Sprintf("%d -> %+v", id, g1.edges)
		assert.Equal(t, len(g.nodes)-1, len(g1.nodes), errOutput)
		assert.Equal(t, len(g.edges)-2, len(g1.edges), errOutput)

		// Verify if not exist a relation equal in g
		for _, e := range g1.edges {
			assert.False(t, e.Source == id-1 && e.Destine == id, errOutput)
			assert.False(t, e.Source == id && e.Destine == id+1, errOutput)
		}

		// Validation if id was removed
		containID := false
		if len(g1.nodes) > id { // g1 doesn't have id = len(g.nodes) -1 because a node was removed
			for _, e := range g1.edges {
				if e.Source == id || e.Destine == id {
					containID = true
				}
			}
			assert.True(t, containID, errOutput)
		}

		// Verify if exist a edge between the id removed and last element because in g has a edge between last and antipnlast
		if len(g1.nodes)-1 > id { // The last and antipnlast element louse this edge
			containID = false
			for _, e := range g1.edges {
				if e.Source == len(g1.nodes)-1 && e.Destine == id {
					containID = true
				}
			}
			assert.True(t, containID, errOutput)
		}

		// Verify if exist a edge between the id removed and id zero because in g has a edge between last and first
		if id == 0 {
			for _, e := range g1.edges {
				if e.Source == id && e.Source == 0 {
					containID = false
				}
			}
			assert.True(t, containID, errOutput)
		}
	}
}

func TestListIsGraph(t *testing.T) {
	// Correct casa
	s := 30
	g := testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(ListNode[util.TestInterface]{
			Attributes: util.TestInterface{Value: i},
		})
	}
	for i := 0; i < s-1; i++ {
		g = g.AddEdge(
			ListEdge[util.TestInterface]{
				Source:     i,
				Destine:    i + 1,
				Attributes: util.TestInterface{Value: i},
			},
		)
	}
	assert.True(t, g.IsGraph())

	// Edge with Destive bigger than number of nodes is invalid
	g = testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(ListNode[util.TestInterface]{
			Attributes: util.TestInterface{Value: i},
		})
	}
	for i := 0; i < s; i++ {
		g = g.AddEdge(
			ListEdge[util.TestInterface]{
				Source:     i,
				Destine:    i + 1,
				Attributes: util.TestInterface{Value: i},
			},
		)
	}
	assert.False(t, g.IsGraph())

	// A empty graph with edges id invalid
	g = testNewListGraphEmpty(1)
	g = g.AddEdge(ListEdge[util.TestInterface]{
		Source:     0,
		Destine:    0,
		Attributes: util.TestInterface{Value: 0},
	})
	assert.False(t, g.IsGraph())

	// A empty graph with edges is valid
	g = testNewListGraphEmpty(1)
	g = g.AddNode(ListNode[util.TestInterface]{
		Attributes: util.TestInterface{Value: 0},
	})
	fmt.Printf(" %+v", g)
	assert.True(t, g.IsGraph())
}

func TestListFilterEgdesReturnID(t *testing.T) {
	s := 100
	g := testNewListTree(s)
	for i := 0; i < s/2; i++ {
		filterEgdes := g.FilterEgdesReturnID(i)
		for _, eID := range filterEgdes {
			e := g.edges[eID]
			assert.True(t, i == e.Source)
			assert.True(t, e.Destine == (i+1)*2-1 || e.Destine == (i+1)*2)
		}
	}
	for i := s / 2; i < s; i++ {
		filterEgdes := g.FilterEgdesReturnID(i)
		assert.Equal(t, 0, len(filterEgdes))
	}
}

func TestListFilterEgdesReturnNodeID(t *testing.T) {
	s := 100
	g := testNewListTree(s)
	for i := 0; i < s/2; i++ {
		destines := g.FilterEgdesReturnNodeID(i)
		for _, destine := range destines {
			assert.True(t, destine == (i+1)*2-1 || destine == (i+1)*2)
		}
	}
	for i := s / 2; i < s; i++ {
		filterEgdes := g.FilterEgdesReturnID(i)
		assert.Equal(t, 0, len(filterEgdes))
	}
}

