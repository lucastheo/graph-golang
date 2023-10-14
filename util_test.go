package graph

type TestInterface struct {
	Value int
}

func testNewListGraphEmpty(v int) *ListGraph[TestInterface, TestInterface, TestInterface] {
	return NewEmpty[TestInterface, TestInterface, TestInterface](
		TestInterface{
			Value: v,
		},
	)
}

func testNewListNode(v int) ListNode[TestInterface] {
	return ListNode[TestInterface]{
		Attributes: TestInterface{
			Value: v,
		},
	}
}

func testNewListEdge(s, d, v int) ListEdge[TestInterface] {
	return ListEdge[TestInterface]{
		Source:     s,
		Destine:    d,
		Attributes: TestInterface{Value: 0},
	}
}

func testNewListTree(s int) *ListGraph[TestInterface, TestInterface, TestInterface] {
	g := testNewListGraphEmpty(1)
	for i := 0; i < s; i++ {
		g = g.AddNode(testNewListNode(i))
	}

	for i := 0; i < s-1; i++ {
		g = g.AddEdge(testNewListEdge(i/2, i+1, i))
	}
	return g
}