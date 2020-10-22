package graph_test

import (
	"test_code/algorithm/graph"
	"testing"
)

func TestGraph(t *testing.T) {
	g := graph.Graph{}
	g.AddRelation("0", "1", 0)
	g.AddRelation("0", "3", 0)
	g.AddRelation("1", "0", 0)
	g.AddRelation("1", "2", 0)
	g.AddRelation("1", "4", 0)
	g.AddRelation("2", "1", 0)
	g.AddRelation("2", "5", 0)
	g.AddRelation("3", "0", 0)
	g.AddRelation("3", "4", 0)
	g.AddRelation("4", "1", 0)
	g.AddRelation("4", "3", 0)
	g.AddRelation("4", "5", 0)
	g.AddRelation("4", "6", 0)
	g.AddRelation("5", "2", 0)
	g.AddRelation("5", "4", 0)
	g.AddRelation("5", "7", 0)
	g.AddRelation("6", "4", 0)
	g.AddRelation("6", "7", 0)
	g.AddRelation("7", "5", 0)
	g.AddRelation("7", "6", 0)

	g.Print()

	//g.BreadthFirstSearch("0", "6")

	g.DepthFirstSearch("1", "7")

	//g.SearchFriendDegree("1", 2)
}
