package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (blueprint [][]cell, checkpoints []checkpoint) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	blueprint = make([][]cell, len(lines))
	for i, line := range lines {
		row := make([]cell, len(line))
		for j, v := range line {
			switch v {
			case '#':
				row[j] = wallCell
			case '.':
				row[j] = emptyCell
			default:
				checkpoints = append(checkpoints, checkpoint{
					pos:    aoc.NewVector2(j, i),
					number: int(v - '0'),
				})
				row[j] = emptyCell
			}
		}
		blueprint[i] = row
	}

	return blueprint, checkpoints
}

func process(blueprint [][]cell, checkpoints []checkpoint) int {
	sort.Slice(checkpoints, func(i, j int) bool { return checkpoints[i].number < checkpoints[j].number })
	graph := buildGraph(blueprint, checkpoints)
	visited := make([]bool, len(checkpoints))

	return findMinDistance(graph, 0, visited, len(checkpoints))
}

func findMinDistance(graph [][]int, node int, visited []bool, nodesLeft int) int {
	if nodesLeft == 1 {
		return 0
	}

	visited[node] = true
	defer func() { visited[node] = false }()

	minLength := math.MaxInt64
	for i, v := range visited {
		if v {
			continue
		}
		minLength = aoc.Min(minLength, graph[node][i]+findMinDistance(graph, i, visited, nodesLeft-1))
	}

	return minLength
}

func buildGraph(blueprint [][]cell, checkpoints []checkpoint) [][]int {
	graph := make([][]int, len(checkpoints))
	for i := range graph {
		graph[i] = make([]int, len(checkpoints))
	}

	for node1, row := range graph {
		for node2, pathLength := range row {
			if node2 == node1 || pathLength != 0 {
				continue
			}
			dst := distance(blueprint, checkpoints[node1].pos, checkpoints[node2].pos)
			graph[node1][node2] = dst
			graph[node2][node1] = dst
		}
	}

	return graph
}

func distance(blueprint [][]cell, start, destination aoc.Vector2) int {
	pq := aoc.NewHeap(func(c1, c2 candidate) bool {
		score1 := c1.pos.Sub(destination).ManhattanDst() + c1.steps
		score2 := c2.pos.Sub(destination).ManhattanDst() + c2.steps
		return score1 < score2
	})

	pq.Push(candidate{pos: start})

	visited := map[aoc.Vector2]bool{}

	for pq.Len() > 0 {
		c := pq.Pop()
		for _, dir := range dirs {
			np := c.pos.Add(dir)
			if np == destination {
				return c.steps + 1
			}
			if blueprint[np.Y][np.X] == wallCell || visited[np] {
				continue
			}
			visited[np] = true
			pq.Push(candidate{pos: np, steps: c.steps + 1})
		}
	}

	panic("path not found")
}

type cell bool

const (
	emptyCell cell = true
	wallCell  cell = false
)

type checkpoint struct {
	pos    aoc.Vector2
	number int
}

type candidate struct {
	pos   aoc.Vector2
	steps int
}

var dirs = []aoc.Vector2{
	aoc.NewVector2(0, 1),
	aoc.NewVector2(0, -1),
	aoc.NewVector2(1, 0),
	aoc.NewVector2(-1, 0),
}
