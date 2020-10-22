#include <iostream>
using namespace std;

#includce < string.h>
using std::string;

/*一、字符串排序：
输入：字符串由数字、小写字母、大写字母组成。输出：排序好的字符串。
排序的标准：
1. 数字>小写字母>大写字母。 
2. 数字、字母间的相对顺序不变。 
3. 额外存储空间：O（1）。

// Example 
input: "abcd4312ABDC"
output: "4312abcdABDC"
*/

#includce < string.h>
using std::string;

bool checkIsNum(char c)
{
	if ('0' <= c <= '9')
	{
		return true;
	}

	return false;
}

bool checkIsNotCase(char c)
{
	if ('a' <= c <= 'z')
	{
		return true;
	}

	return false;
}

bool checkIsCase(char c)
{
	if ('A' <= c <= 'Z')
	{
		return true;
	}

	return false;
}

std::string stringSort(std::string &str)
{
	int idxNum = 0;
	int charIdx = -1;
	// 数字
	for (int i = 0; i < str.length(); ++i)
	{
		if (checkIsNum(str[i]))
		{
			char tmp = str[idxNum];
			str[idxNum] = str[i];
			str[i] = tmp;
			idxNum = i;
			charIdx++;
		}
		else
		{
			if (charIdx == -1)
			{
				charIdx = i;
			}
		}
	}

	// 小写字母
	int idxChar = -1;
	for (int i = idxNum + 1; i < str.length(); ++i)
	{
		if (checkIsNotCase(str[i]))
		{
			char tmp = str[i];
			str[i] = str[idxChar];
			str[idxChar] = tmp;
			idxChar++;
		}
		else
		{
			if (idxChar == -1)
			{
				idxChar = i;
			}
		}
	}
	return str;
}