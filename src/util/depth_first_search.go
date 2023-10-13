package util

type DeapthFIrstSearch struct {
	recursiveStack []bool
	visit          []bool
	nodes          []int
	neighbors      func(v int) []int
}

func NewDFS(nodes []int, neighbour func(v int) []int) DeapthFIrstSearch {
	return DeapthFIrstSearch{
		make([]bool, len(nodes)),
		make([]bool, len(nodes)),
		nodes,
		neighbour,
	}
}

func (dfs DeapthFIrstSearch) HasCyclo() bool {
	for _, node := range dfs.nodes {
		if dfs.hasCycloVerify(node) {
			return true
		}
	}
	return false
}

func (dfs DeapthFIrstSearch) hasCycloVerify(node int) bool {
	if dfs.recursiveStack[node] {
		return true
	}
	if dfs.visit[node] {
		return false
	}

	dfs.visit[node] = true
	dfs.recursiveStack[node] = true

	for _, neighbor := range dfs.neighbors(node) {
		if dfs.hasCycloVerify(neighbor) {
			return true
		}
	}
	dfs.recursiveStack[node] = false
	return false
}
