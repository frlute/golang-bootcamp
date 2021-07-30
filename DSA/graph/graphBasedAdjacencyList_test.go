package graph

import "testing"

func Test_graphBasedList_lifecycle(t *testing.T) {
	graphInstance := newGraph(8)

	graphInstance.AddEdge(0, 1)
	graphInstance.AddEdge(1, 3)
	graphInstance.AddEdge(1, 6)
	graphInstance.AddEdge(6, 7)
	graphInstance.AddEdge(0, 2)
	graphInstance.AddEdge(2, 4)
	graphInstance.AddEdge(2, 5)

	graphInstance.Display()

	graphInstance.DFS()
}

func Test_graphBasedList_bfs(t *testing.T) {
	graphInstance := newGraph(8)

	graphInstance.AddEdge(0, 1)
	graphInstance.AddEdge(1, 2)
	graphInstance.AddEdge(1, 3)
	graphInstance.AddEdge(1, 4)
	graphInstance.AddEdge(2, 5)
	graphInstance.AddEdge(2, 6)
	graphInstance.AddEdge(3, 5)
	graphInstance.AddEdge(4, 5)
	graphInstance.AddEdge(6, 7)

	graphInstance.Display()

	graphInstance.BFS()
}
