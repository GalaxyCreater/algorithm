package zong_example

import "test_code/algorithm/skiplist"

type HuntHead struct {
	id    int
	score int
	name  string
}

type HuntHeadCache struct {
	huntHeadMap map[int]*HuntHead
	scoreRank   *skiplist.SkipList
}

func ConstructHuntHeadCache() *HuntHeadCache {
	c := &HuntHeadCache{
		huntHeadMap: map[int]*HuntHead{},
		scoreRank:   skiplist.ConsturctSkipList(),
	}

	return c
}

func (self *HuntHeadCache) AddHuntHead(id int, score int, name string) {
	hunt := HuntHead{
		id:    id,
		score: score,
		name:  name,
	}
	self.huntHeadMap[id] = &hunt
	self.scoreRank.AddNode(hunt, score)
}
