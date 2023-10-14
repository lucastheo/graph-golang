package graph

func IsTree(g Graph) bool {
	return !NewDFS(
		g.Nodes(),
		g.FilterEgdesReturnNodeID,
	).HasCyclo()
}
