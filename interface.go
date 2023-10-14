package graph

type Graph interface {
	Nodes() []int
	FilterEgdesReturnID(n int) []int
	FilterEgdesReturnNodeID(n int) []int
	Print()
}
