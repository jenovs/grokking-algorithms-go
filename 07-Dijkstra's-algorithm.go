package main

import (
	"fmt"
	"math"
)

type graph struct {
	node  string
	edges map[string]int
}

func dijkstra() {
	graph := map[string]map[string]int{
		"book": {
			"LP":     5,
			"poster": 0,
		},
		"LP": {
			"drum_set":    20,
			"bass_guitar": 15,
		},
		"bass_guitar": {
			"piano": 20,
		},
		"poster": {
			"bass_guitar": 30,
			"drum_set":    35,
		},
		"drum_set": {
			"piano": 10,
		},
		"piano": {},
	}

	start := "book"

	costs := map[string]float64{
		"piano": math.Inf(1),
	}
	updateCosts(graph, start, costs, 0)

	parents := map[string]string{}
	updateParents(graph, start, parents)

	processed := []string{}

	node := lowestCost(costs, processed)
	var nodeCost float64

	for len(node) > 0 {
		nodeCost = costs[node]
		for k, v := range graph[node] {
			if nv := float64(v) + nodeCost; nv < costs[k] {
				costs[k] = nv
				parents[k] = node
			}
		}
		processed = append(processed, node)
		updateCosts(graph, node, costs, nodeCost)
		node = lowestCost(costs, processed)
	}

	fmt.Println("costs:  ", costs)
	fmt.Println("parents:", parents)
}

func updateCosts(g map[string]map[string]int, s string, c map[string]float64, cs float64) {
	nc := g[s]
	for k, v := range nc {
		c[k] = float64(v) + cs
	}
}

func updateParents(g map[string]map[string]int, s string, p map[string]string) {
	nc := g[s]
	for k := range nc {
		p[k] = s
	}
}

func lowestCost(c map[string]float64, p []string) string {
	var minK string
	minV := math.Inf(1)
	for k, v := range c {
		if v < minV && !includes(p, k) {
			minV, minK = v, k
		}
	}
	return minK
}

func includes(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
