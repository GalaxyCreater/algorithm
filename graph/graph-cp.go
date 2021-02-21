package graph

import (
	"fmt"
)

/*
图:
https://static001.geekbang.org/resource/image/00/ea/002e9e54fb0d4dbf5462226d946fa1ea.jpg
*/

/*有向有权重的图*/
type Relationship struct {
	weight int //关系的权重
	friend *GraphNode
}

type GraphNode struct {
	name      string // 节点名字
	relations map[string]*Relationship
}

type Graph struct {
	members map[string]*GraphNode

	// 用于深度优先遍历
	found bool
}

/*
me ——> fri 方向的关系，权重为weight
*/
func (g *Graph) AddRelation(me string, fri string, weight int) {
	var m *GraphNode = nil
	var f *GraphNode = nil
	ok := false
	if m, ok = g.members[me]; !ok {
		m = &GraphNode{
			name:      me,
			relations: map[string]*Relationship{},
		}
	}
	if f, ok = g.members[fri]; !ok {
		f = &GraphNode{
			name:      fri,
			relations: map[string]*Relationship{},
		}
	}

	m.relations[fri] = &Relationship{
		weight: weight,
		friend: f,
	}
	if len(g.members) == 0 {
		g.members = map[string]*GraphNode{}
	}
	g.members[me] = m
	g.members[fri] = f
}

func (g *Graph) reset() {
	g.found = false
}

func (g *Graph) Print() {
	for _, me := range g.members {
		for _, rel := range me.relations {
			fmt.Printf("%s -----%d----- %s\n", me.name, rel.weight, rel.friend.name)
		}
		fmt.Println("")
	}
}

/*
打印路径
*/
func (self *Graph) PrintPre(pre map[string]string, beg, end string) {
	if len(pre) == 0 {
		fmt.Println("not find road", beg, end)
		return
	}
	//fmt.Printf("%v\n", pre)
	path := []string{}
	// 从结束到开始倒推
	tmp := end
	path = append(path, tmp)
	for tmp != beg {
		tmp1, ok := pre[tmp]
		if !ok {
			fmt.Printf("invalid key:%s\n", tmp)
			return
		}
		tmp = tmp1
		path = append(path, tmp)
		if tmp == beg {
			break
		}
	}
	fmt.Printf("%s - %s, path: ", beg, end)
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%s    ", path[i])
	}
	fmt.Println()
}

/*
广度优先(找出的是最短路径)
*/
func (g *Graph) BreadthFirstSearch(beg string, end string) {
	node, ok := g.members[beg]
	if !ok {
		fmt.Println("no member:", beg)
		return
	}
	if beg == end {
		return
	}
	que := []*GraphNode{}
	que = append(que, node)
	visted := map[string]bool{} // 记录是否访问过
	road := map[string]string{} // 记录路径,key:节点名,value:上一个节点名;为什么不是记录下一个节点?

	for len(que) > 0 {
		n := que[0]
		for k, v := range n.relations {
			if visted[k] {
				continue
			}
			visted[k] = true
			road[k] = n.name

			// 新发现的加入队列
			que = append(que, v.friend)
		}
		// 移除队头
		que = que[1:]
	}

	// 打印路径
	g.PrintPre(road, beg, end)
}

func (g *Graph) BfTest() {
	fmt.Println("")
	fmt.Println("---------------广度优先测试:")
	g.BreadthFirstSearch("0", "0")
	g.BreadthFirstSearch("0", "6")
	g.BreadthFirstSearch("1", "6")
	g.BreadthFirstSearch("6", "2")
	g.BreadthFirstSearch("0", "7")
}

/*
深度优先
*/
func (g *Graph) DpFind(beg, end string) {
	visted := map[string]bool{}
	preRoad := map[string]string{}

	n, ok := g.members[beg]
	if !ok {
		fmt.Println("node not exists")
		return
	}

	// 从起始点的所有连接点开始
	visted[beg] = true
	for k := range n.relations {
		g.dpFind(beg, k, end, preRoad, visted)
	}

	g.PrintPre(preRoad, beg, end)
	g.reset() // 方便测试,重置成员变量
}

/*
@end:用于判断结束条件
*/
func (g *Graph) dpFind(beg, next, end string,
	preRoad map[string]string, visted map[string]bool) {

	if g.found { // 如果深度的某个分支找到了,退出所有分支
		return
	}
	if visted[next] {
		return
	}

	preRoad[next] = beg
	visted[next] = true

	if end == next {
		g.found = true
		return
	}
	n, ok := g.members[next] // next开始继续找
	if !ok {
		fmt.Println("not exists node:", n.name)
		return
	}
	for k := range n.relations {
		g.dpFind(next, k, end, preRoad, visted)
	}
}

func (g *Graph) DfTest() {
	fmt.Println("")
	fmt.Println("---------------深度度优先测试:")
	g.DpFind("0", "0")
	g.DpFind("0", "6")
	g.DpFind("1", "6")
	g.DpFind("6", "2")
	g.DpFind("0", "7")
}

/*
搜索N度好友(广度优先)
*/
func (g *Graph) SearchFriendDegree(name string, degree int) {
	visted := map[string]bool{}
	n, ok := g.members[name]
	if !ok {
		fmt.Println("not exist member:", name)
		return
	}

	i := 0
	dq := []*GraphNode{}
	dq = append(dq, n)
	for ; i < degree; i++ {
		nextDq := map[string]*GraphNode{} // 这里不能用[]*GraphNode,当relations有相同好友时,切片里会存在重复数据
		for i := 0; i < len(dq); i++ {
			n = dq[i]
			if visted[n.name] {
				continue
			}
			visted[n.name] = true
			for k, v := range n.relations {
				if visted[k] {
					continue
				}
				nextDq[k] = v.friend
			}
		}
		dq = []*GraphNode{}
		for _, v := range nextDq {
			dq = append(dq, v)
		}
	}

	if i == degree && len(dq) > 0 {
		for i := 0; i < len(dq); i++ {
			fmt.Printf("%s  ", dq[i].name)
		}
		fmt.Println("")
	} else {
		fmt.Println("not found")
	}

}

func (g *Graph) FriendDegreeTest() {
	fmt.Println("")
	fmt.Println("---------------n度好友测试:")
	g.SearchFriendDegree("1", 3)
	g.SearchFriendDegree("1", 2)
	g.SearchFriendDegree("1", 0)
	g.SearchFriendDegree("1", 5)
}
