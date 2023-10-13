package graph

import (
	"fmt"
)

type EdgeAttributes = interface{}
type NodeAttributes = interface{}
type GraphAttributes = interface{}
type NodeIdType = int
type EdgeIdType = int

type ListGraph[GA GraphAttributes, NA NodeAttributes, EA EdgeAttributes] struct {
	Attributes GA
	nodes      []ListNode[NA]
	edges      []ListEdge[EA]
}

type ListNode[NA NodeAttributes] struct {
	Attributes NA
}

type ListEdge[EA EdgeAttributes] struct {
	Attributes EA
	Source     NodeIdType
	Destine    NodeIdType
}

func NewEmpty[GA GraphAttributes, NA NodeAttributes, EA EdgeAttributes](
	a GraphAttributes,
) *ListGraph[GA, NA, EA] {
	return &ListGraph[GA, NA, EA]{
		Attributes: a.(GA),
	}
}

func (g *ListGraph[GA, NA, EA]) AddNode(n ListNode[NA]) *ListGraph[GA, NA, EA] {
	nodes := make([]ListNode[NA], len(g.nodes)+1)
	copy(nodes, g.nodes)
	nodes[len(g.nodes)] = n

	edges := make([]ListEdge[EA], len(g.edges))
	copy(edges, g.edges)

	return &ListGraph[GA, NA, EA]{
		Attributes: g.Attributes,
		nodes:      nodes,
		edges:      edges,
	}
}

func (g *ListGraph[GA, NA, EA]) RemoveNode(id NodeIdType) *ListGraph[GA, NA, EA] {
	lastNode := len(g.nodes) - 1
	nodes := make([]ListNode[NA], lastNode)
	copy(nodes, g.nodes[:len(g.nodes)])
	if id != lastNode {
		nodes[id] = g.nodes[lastNode]
	}

	useRemoveID := func(e ListEdge[EA]) bool {
		return e.Source == id || e.Destine == id
	}

	s := 0
	for _, e := range g.edges {
		if !useRemoveID(e) {
			s++
		}
	}
	edges := make([]ListEdge[EA], s)

	i := 0
	for _, e := range g.edges {
		if !useRemoveID(e) {
			if e.Source == lastNode {
				e.Source = id
			}
			if e.Destine == lastNode {
				e.Destine = id
			}
			edges[i] = e
			i++
		}
	}

	return &ListGraph[GA, NA, EA]{
		Attributes: g.Attributes,
		nodes:      nodes,
		edges:      edges,
	}
}

func (g *ListGraph[GA, NA, EA]) RemoveEdge(id EdgeIdType) (*ListGraph[GA, NA, EA], error) {
	if len(g.edges) <= id {
		return nil, nil
	}
	nodes := make([]ListNode[NA], len(g.nodes)+1)
	copy(nodes, g.nodes)

	edges := make([]ListEdge[EA], len(g.edges)-1)
	copy(edges, g.edges[:len(edges)])

	if id != len(g.edges)-1 {
		edges[id] = g.edges[len(edges)]
	}

	return &ListGraph[GA, NA, EA]{
		Attributes: g.Attributes,
		nodes:      nodes,
		edges:      edges,
	}, nil
}

func (g *ListGraph[GA, NA, EA]) AddEdge(e ListEdge[EA]) *ListGraph[GA, NA, EA] {
	nodes := make([]ListNode[NA], len(g.nodes))
	copy(nodes, g.nodes)

	edges := make([]ListEdge[EA], len(g.edges)+1)
	copy(edges, g.edges)
	edges[len(g.edges)] = e

	return &ListGraph[GA, NA, EA]{
		Attributes: g.Attributes,
		nodes:      nodes,
		edges:      edges,
	}
}

func (g *ListGraph[GA, NA, EA]) IsGraph() bool {
	if len(g.nodes) == 0 {
		return len(g.edges) == 0
	}

	for _, e := range g.edges {
		smallerThanId := len(g.nodes) > e.Source && len(g.nodes) > e.Destine
		if !smallerThanId {
			return false
		}
	}
	return true
}

func (g *ListGraph[GA, NA, EA]) FilterEgdesReturnID(n int) []int {
	edges := make([]int, 0)

	for i, e := range g.edges {
		if e.Source == n {
			edges = append(edges, i)
		}
	}
	return edges
}

func (g *ListGraph[GA, NA, EA]) FilterEgdesReturnNodeID(n int) []int {
	edges := make([]int, 0)

	for _, e := range g.edges {
		if e.Source == n {
			edges = append(edges, e.Destine)
		}
	}
	return edges
}

func (g *ListGraph[GA, NA, EA]) FilterEgdes(n int) []ListEdge[EA] {
	edgesID := g.FilterEgdesReturnID(n)

	edges := make([]ListEdge[EA], len(edgesID))
	for i, eID := range edgesID {
		edges[i] = g.edges[eID]
	}
	return edges
}

func (g *ListGraph[GA, NA, EA]) Print() {
	for i, _ := range g.nodes {
		for _, e := range g.edges {
			if i == e.Source {
				fmt.Printf("\t  %d->%d\n", i, e.Destine)
			}
		}
	}
}

func (g *ListGraph[GA, NA, EA]) Nodes() []int {
	return make([]int, len(g.nodes))
}