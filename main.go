package main

import (
	"fmt"
	"math"
)

// Graph представляет собой ориентированный граф
type Graph struct {
	vertices int
	edges    [][]float64 // матрица смежности с весами
}

// NewGraph создает новый ориентированный граф с заданным количеством вершин
func NewGraph(vertices int) *Graph {
	// Инициализируем матрицу смежности бесконечностями
	edges := make([][]float64, vertices)
	for i := range edges {
		edges[i] = make([]float64, vertices)
		for j := range edges[i] {
			if i == j {
				edges[i][j] = 0 // Расстояние от вершины до себя равно 0
			} else {
				edges[i][j] = math.Inf(1) // Используем бесконечность для обозначения отсутствия ребра
			}
		}
	}

	return &Graph{
		vertices: vertices,
		edges:    edges,
	}
}

// AddEdge добавляет направленное ребро из вершины u в вершину v с весом weight
func (g *Graph) AddEdge(u, v int, weight float64) {
	if u >= 0 && u < g.vertices && v >= 0 && v < g.vertices {
		g.edges[u][v] = weight
	}
}

// ShortestPath находит кратчайший путь между двумя вершинами с использованием алгоритма Флойда-Уоршелла
func (g *Graph) ShortestPath() [][]float64 {
	// Создаем копию матрицы смежности
	dist := make([][]float64, g.vertices)
	for i := range dist {
		dist[i] = make([]float64, g.vertices)
		copy(dist[i], g.edges[i])
	}

	// Алгоритм Флойда-Уоршелла
	for k := 0; k < g.vertices; k++ {
		for i := 0; i < g.vertices; i++ {
			for j := 0; j < g.vertices; j++ {
				// Если путь через вершину k короче, чем текущий путь
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

// GetShortestDistance возвращает кратчайшее расстояние между двумя вершинами
func (g *Graph) GetShortestDistance(start, end int) float64 {
	if start < 0 || start >= g.vertices || end < 0 || end >= g.vertices {
		return math.Inf(1) // Возвращаем бесконечность, если индексы некорректны
	}

	allDistances := g.ShortestPath()
	return allDistances[start][end]
}

// PrintGraph выводит информацию о графе
func (g *Graph) PrintGraph() {
	fmt.Println("Матрица смежности графа:")
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if math.IsInf(g.edges[i][j], 1) {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%.1f ", g.edges[i][j])
			}
		}
		fmt.Println()
	}
}

// PrintShortestPaths выводит кратчайшие пути между всеми парами вершин
func (g *Graph) PrintShortestPaths() {
	dist := g.ShortestPath()
	fmt.Println("\nКратчайшие пути между всеми парами вершин:")
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if math.IsInf(dist[i][j], 1) {
				fmt.Printf("Кратчайший путь из %d в %d: недостижим\n", i, j)
			} else {
				fmt.Printf("Кратчайший путь из %d в %d: %.1f\n", i, j, dist[i][j])
			}
		}
	}
}

func main() {
	// Создаем граф с 4 вершинами
	graph := NewGraph(4)

	// Добавляем ребра (вершины нумеруются с 0)
	graph.AddEdge(0, 1, 5)  // Ребро из 0 в 1 с весом 5
	graph.AddEdge(0, 2, 3)  // Ребро из 0 в 2 с весом 3
	graph.AddEdge(1, 2, 2)  // Ребро из 1 в 2 с весом 2
	graph.AddEdge(1, 3, 6)  // Ребро из 1 в 3 с весом 6
	graph.AddEdge(2, 3, 7)  // Ребро из 2 в 3 с весом 7

	// Выводим исходный граф
	graph.PrintGraph()

	// Выводим кратчайшие пути между всеми парами вершин
	graph.PrintShortestPaths()

	// Пример получения кратчайшего расстояния между конкретными вершинами
	start, end := 0, 3
	distance := graph.GetShortestDistance(start, end)
	if math.IsInf(distance, 1) {
		fmt.Printf("\nНет пути из %d в %d\n", start, end)
	} else {
		fmt.Printf("\nКратчайшее расстояние из %d в %d: %.1f\n", start, end, distance)
	}
}