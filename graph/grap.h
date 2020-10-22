#include <iostream>
#include <deque>

using namespace std;

struct node
{
	int val;
};

typedef deque<node *> DqType;
class grap
{
	node *nodeMap[10][10];

public:
	void bfs();
	void dfs();
};

void grap::dfs()
{
}

void grap::bfs()
{
}