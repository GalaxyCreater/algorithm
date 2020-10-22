package graph

import (
	"fmt"
)

/*
有向有权重的图
*/

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
func (self *Graph) AddRelation(me string, fri string, weight int) {
	var m *GraphNode = nil
	var f *GraphNode = nil
	ok := false
	if m, ok = self.members[me]; !ok {
		m = &GraphNode{
			name:      me,
			relations: map[string]*Relationship{},
		}
	}
	if f, ok = self.members[fri]; !ok {
		f = &GraphNode{
			name:      fri,
			relations: map[string]*Relationship{},
		}
	}

	m.relations[fri] = &Relationship{
		weight: weight,
		friend: f,
	}
	if len(self.members) == 0 {
		self.members = map[string]*GraphNode{}
	}
	self.members[me] = m
	self.members[fri] = f
}

func (self *Graph) Print() {
	for _, me := range self.members {
		for _, rel := range me.relations {
			fmt.Printf("%s -----%d----- %s\n", me.name, rel.weight, rel.friend.name)
		}
		fmt.Println("")
	}
	fmt.Println("--------------------------------")
}

/*
广度优先(找出的是最短路径)
*/
func (self *Graph) BreadthFirstSearch(beg string, end string) {
	visited := map[string]bool{}
	pre := map[string]string{}

	node, ok := self.members[beg]
	if !ok {
		return
	}

	que := []*GraphNode{}
	que = append(que, node) //保存需要访问的节点
	visited[beg] = true
	found := false
	for len(que) > 0 {
		node := que[0]
		for _, v := range node.relations {
			if visited[v.friend.name] { //访问过的节点，跳过
				continue
			}
			pre[v.friend.name] = node.name //记录访问路径
			visited[v.friend.name] = true
			if v.friend.name == end {
				found = true
				break
			}
			que = append(que, v.friend)
		}
		que = que[1:]
	}

	// 输出
	if found {
		self.PrintPre(pre, beg, end)
	}
}

/*
打印路径
*/
func (self *Graph) PrintPre(pre map[string]string, beg, end string) {
	fmt.Printf("%v\n", pre)
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
	fmt.Printf("path: ")
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%s    ", path[i])
	}
	fmt.Println()
}

/*
深度优先
*/
func (self *Graph) DepthFirstSearch(beg, end string) {
	self.found = false
	visited := map[string]bool{}
	pre := map[string]string{} // 路径
	visited[beg] = true
	if beg != end {
		self.depthFirstSearch(beg, end, pre, visited)
	}
	self.PrintPre(pre, beg, end)
}

func (self *Graph) depthFirstSearch(beg, end string, pre map[string]string, visited map[string]bool) {
	if self.found {
		return
	}

	m, ok := self.members[beg]
	if !ok {
		fmt.Println("can not find key:", beg)
		return
	}

	for _, v := range m.relations {
		if visited[v.friend.name] {
			continue
		}
		visited[v.friend.name] = true
		pre[v.friend.name] = beg

		if m.name == end {
			self.found = true
			return
		}
		self.depthFirstSearch(v.friend.name, end, pre, visited)
	}

}

/*
搜索N度好友
*/
func (self *Graph) SearchFriendDegree(beg string, degree int) {
	visited := map[string]bool{}
	pre := map[string]string{}

	node, ok := self.members[beg]
	if !ok {
		return
	}

	que := []*GraphNode{}
	que = append(que, node) //保存需要访问的节点
	visited[beg] = true

	degreeMap := map[string]int{}
	degreeMap[beg] = 0
	curDegree := 1
	borad := beg
	nextBorad := ""
	for len(que) > 0 {
		node := que[0]
		for _, v := range node.relations {
			if visited[v.friend.name] { //访问过的节点，跳过
				continue
			}
			pre[v.friend.name] = node.name //记录访问路径
			visited[v.friend.name] = true
			que = append(que, v.friend)
			degreeMap[v.friend.name] = curDegree

			nextBorad = v.friend.name
		}

		if borad == node.name {
			borad = nextBorad
			curDegree++
		}

		// 找到对应的度了，结束
		if curDegree > degree {
			break
		}

		que = que[1:]
	}

	fmt.Printf("%v\n", degreeMap)
	for k, v := range degreeMap {
		if v == degree {
			fmt.Printf("%s    ", k)
		}
	}
}
