package main

import (
	"ds.com/ds/algorithms/sorting"
	"fmt"
)

// 3--------4--------5
// |        |        |
// |        |        |
// |        |        |
// 1--------2        6
//  \      /
//   \    /
//    \  /
//     0

func main() {
	//g := ds.NewGraph()
	//g.AddVertex("0")
	//g.AddVertex("1")
	//g.AddVertex("2")
	//g.AddVertex("3")
	//g.AddVertex("4")
	//g.AddVertex("5")
	//g.AddVertex("6")
	//g.AddEdge("3", "1")
	//g.AddEdge("3", "4")
	//g.AddEdge("4", "2")
	//g.AddEdge("4", "5")
	//g.AddEdge("1", "2")
	//g.AddEdge("1", "0")
	//g.AddEdge("0", "2")
	//g.AddEdge("6", "5")
	//g.ShowConnections()

	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	fmt.Println(sorting.Bubble(numbers))

}
