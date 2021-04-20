package graph

import "container/list"

// Graph 无向图
type Graph struct {
	vertexCount int          // 顶点个数
	adj         []*list.List // 邻接表
}

func newGraph(capacity int) *Graph {
	gr := &Graph{}
	gr.vertexCount = capacity

	gr.adj = make([]*list.List, capacity)
	for index := range gr.adj {
		gr.adj[index] = list.New()
	}
	return gr
}

//insert as add edge，一条边存2次
func (g *Graph) addEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}
