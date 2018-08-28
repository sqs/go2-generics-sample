package main

// This code is taken from the Go 2 generics proposal at https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md.

contract Graph(n Node, e Edge) {
	var edges []Edge = n.Edges()
	var nodes []Node = e.Nodes()
}

func ShortestPath(type N, E Graph)(src, dst N) []E {
	panic("TODO")
}
