package function

import (
	"graph/src/graph"
	"graph/src/util"
)

func IsTree(g graph.Graph) bool {
	return !util.NewDFS(
		g.Nodes(),
		g.FilterEgdesReturnNodeID,
	).HasCyclo()
}
