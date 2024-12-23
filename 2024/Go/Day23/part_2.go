package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

// Graph structure to represent an undirected graph
type Graph struct {
	Edges map[string]map[string]bool
}

// NewGraph initializes an empty graph
func NewGraph() *Graph {
	return &Graph{Edges: make(map[string]map[string]bool)}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(a, b string) {
	if g.Edges[a] == nil {
		g.Edges[a] = make(map[string]bool)
	}
	if g.Edges[b] == nil {
		g.Edges[b] = make(map[string]bool)
	}
	g.Edges[a][b] = true
	g.Edges[b][a] = true
}

// FindCliquesRecursive finds all maximal cliques in the graph
func (g *Graph) FindCliquesRecursive() [][]string {
	var cliques [][]string

	var bronKerbosch func(r, p, x map[string]bool)
	bronKerbosch = func(r, p, x map[string]bool) {
		if len(p) == 0 && len(x) == 0 {
			// Found a maximal clique
			clique := make([]string, 0, len(r))
			for node := range r {
				clique = append(clique, node)
			}
			sort.Strings(clique)
			cliques = append(cliques, clique)
			return
		}

		for node := range p {
			nr := copySet(r)
			nr[node] = true

			np := intersectSets(p, g.Edges[node])
			nx := intersectSets(x, g.Edges[node])

			bronKerbosch(nr, np, nx)

			delete(p, node)
			x[node] = true
		}
	}

	r := make(map[string]bool)
	p := make(map[string]bool)
	x := make(map[string]bool)
	for node := range g.Edges {
		p[node] = true
	}
	bronKerbosch(r, p, x)

	return cliques
}

// Helper functions
func copySet(set map[string]bool) map[string]bool {
	copy := make(map[string]bool, len(set))
	for k := range set {
		copy[k] = true
	}
	return copy
}

func intersectSets(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range a {
		if b[k] {
			result[k] = true
		}
	}
	return result
}

// getIntoLANParty reads the input file, builds the graph, and finds the largest clique
func getIntoLANParty() string {
	g := NewGraph()

	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ps := strings.Split(line, "-")
		if len(ps) == 2 {
			g.AddEdge(ps[0], ps[1])
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find cliques
	cliques := g.FindCliquesRecursive()
	// Sort cliques by size (descending) and lexicographical order
	sort.Slice(cliques, func(i, j int) bool {
		if len(cliques[i]) != len(cliques[j]) {
			return len(cliques[i]) > len(cliques[j])
		}
		return strings.Join(cliques[i], ",") < strings.Join(cliques[j], ",")
	})

	// Get the largest clique and sort it lexicographically
	if len(cliques) == 0 {
		return ""
	}
	largestClique := cliques[0]
	return strings.Join(largestClique, ",")
}

// func main() {
// 	result := getIntoLANParty()
// 	fmt.Println(result)
// }
