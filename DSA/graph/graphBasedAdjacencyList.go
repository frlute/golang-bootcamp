package graph

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/frlute/golang-bootcamp/DSA/queue"
	"github.com/frlute/golang-bootcamp/DSA/stack"
)

// Graph 无向图
type graphBasedAdjacencyList struct {
	cap int //容量
	// TODO 只能数组？还是都可以，不太确定是否需要有序
	// TODO 需要排序，否则遍历路径存在随机性
	adj []*list.List // 邻接表, 是一个双链表
}

// TODO 感觉还是有点问题，把 capacity 作为了初始化的节点数
func newGraph(capacity int) Graph {
	gr := &graphBasedAdjacencyList{}
	gr.cap = capacity

	gr.adj = make([]*list.List, capacity)
	for index := 0; index < capacity; index++ {
		gr.adj[index] = list.New()
	}
	return gr
}

//insert as add edge，一条边存2次
func (g *graphBasedAdjacencyList) AddEdge(src, target int) {
	// if srcVertex, ok := g.adj[src]; ok {
	// 	srcVertex.PushBack(target)
	// } else {
	// 	panic(fmt.Sprintf("the vertex %+v not exits.", src))
	// }

	// if targetVertex, ok := g.adj[target]; ok {
	// 	targetVertex.PushBack(src)
	// } else {
	// 	panic(fmt.Sprintf("the vertex %+v not exits.", target))
	// }
	g.adj[src].PushBack(target)
	g.adj[target].PushBack(src)
}

func (g *graphBasedAdjacencyList) AddVertex(vertex int) {
	g.cap++
	g.adj[vertex] = list.New()
}

func (g *graphBasedAdjacencyList) Display() {
	for vertex, adjEdge := range g.adj {
		edge := g.getAdjacencyVertet(adjEdge)
		fmt.Printf("vertex: %+v, edge: %+v\n", vertex, edge)
	}
}

func (g *graphBasedAdjacencyList) getAdjacencyVertet(adjEdge *list.List) []int {
	adjVertex := make([]int, 0, g.cap)
	head := adjEdge.Front()
	for head != nil {
		adjVertex = append(adjVertex, head.Value.(int))
		head = head.Next()
	}
	return adjVertex
}

// 演示版，不做实际用途
func (g *graphBasedAdjacencyList) BFS() {
	visited := make(map[int]struct{})
	cache := queue.NewSequentialQueue(g.cap)

	startVertex := 0
	var path strings.Builder

	visited[startVertex] = struct{}{}
	path.WriteString(fmt.Sprintf("%d->", startVertex))
	cache.Enqueue(startVertex)

	for !cache.IsEmpty() {
		tmpVertex := cache.Dequeue().(int)

		adjVertex := g.getAdjacencyVertet(g.adj[tmpVertex])
		for _, vertex := range adjVertex {
			if _, ok := visited[vertex]; !ok {
				visited[vertex] = struct{}{}
				path.WriteString(fmt.Sprintf("%d->", vertex))
				cache.Enqueue(vertex)
			}
		}
	}

	fmt.Printf("bfs path: %+v\n", path.String())
}

// 演示版，不做实际用途
func (g *graphBasedAdjacencyList) DFS() {
	visited := make(map[int]struct{})
	stackCache := stack.NewStackBasedArray(int64(g.cap))

	var path strings.Builder
	firstVertex := 0

	visited[firstVertex] = struct{}{}
	path.WriteString(fmt.Sprintf("%d->", firstVertex))
	stackCache.Push(firstVertex)

	for !stackCache.IsEmpty() {
		unvisitedVertex := stackCache.Peek().(int)
		if _, ok := visited[unvisitedVertex]; ok {
			stackCache.Pop()
		} else {
			visited[unvisitedVertex] = struct{}{}
			path.WriteString(fmt.Sprintf("%d->", unvisitedVertex))
		}

		adjVertex := g.getAdjacencyVertet(g.adj[unvisitedVertex])
		for _, vertex := range adjVertex {
			if _, ok := visited[vertex]; !ok {
				stackCache.Push(vertex)
			}
		}
	}

	fmt.Printf("dfs path: %+v\n", path.String())
}
