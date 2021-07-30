package graph

type Graph interface {
	AddVertex(vertex int)
	AddEdge(src, target int)
	Display()

	BFS()
	DFS()
}

// GraphWithWeight representation the graph with weight
type GraphWithWeight interface {
	Graph

	SpanningTree()
}
