package main

import (
	"encoding/csv"
	"log"
	"os"
	"slices"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func main() {
	file, err := os.Open("23/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '-'

	connections, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	g := simple.NewUndirectedGraph()
	nodes := make(map[string]graph.Node)
	computers := make(map[graph.Node]string)
	for _, connection := range connections {
		for _, computer := range connection {
			if _, exists := nodes[computer]; !exists {
				node := g.NewNode()
				g.AddNode(node)

				nodes[computer] = node
				computers[node] = computer
			}
		}
		g.SetEdge(g.NewEdge(nodes[connection[0]], nodes[connection[1]]))
	}

	// Part 1: Sets of three computers
	triangles := make(map[[3]int64][3]graph.Node)
	for edges := g.Edges(); edges.Next(); {
		edge := edges.Edge()
		u, v := edge.From(), edge.To()

		for neighbours := g.From(u.ID()); neighbours.Next(); {
			w := neighbours.Node()
			if g.HasEdgeBetween(v.ID(), w.ID()) {
				triangle := [3]graph.Node{u, v, w}
				triangleId := [3]int64{u.ID(), v.ID(), w.ID()}
				slices.Sort(triangleId[:])

				for _, node := range triangle {
					computer := computers[node]
					if strings.HasPrefix(computer, "t") {
						triangles[triangleId] = [3]graph.Node{u, v, w}
						break
					}
				}
			}
		}
	}

	log.Printf("Sets of three computers: %d\n", len(triangles))

	// Part 2: Password
	networks := topo.BronKerbosch(g)

	var largestNetwork []graph.Node
	for _, network := range networks {
		if len(network) > len(largestNetwork) {
			largestNetwork = network
		}
	}

	var largestNetworkComputers []string
	for _, node := range largestNetwork {
		largestNetworkComputers = append(largestNetworkComputers, computers[node])
	}

	slices.Sort(largestNetworkComputers)
	password := strings.Join(largestNetworkComputers, ",")

	log.Printf("Password: %s", password)
}
