/*************************************************************************
  > File Name: main.cpp
  > Author: fjp
  > Mail: fjp@xxx.com
  > Created Time: Wed 24 Feb 2021 05:26:24 PM CST
 ************************************************************************/

#include <iostream>
#include <vector>
using namespace std;
vector<int> ShuffleArray(vector<int> numbers)
{
	srand(time(0));
	for(int i=0;i<numbers.size();i++)
	{
		// 生成一个[i, 52-i]的随机数
		int r=rand()%(numbers.size()-i+1)+i;
		//交换元素
        swap(numbers[i],numbers[r]);
	}
	return numbers;
}

int main(){
	vector<int> numbers={0,0,1,1,2,2,3,3};
	numbers = ShuffleArray(numbers);
	for(int i=0;i<numbers.size();i++){
		cout<<numbers[i]<<" ";
	}
	cout<<endl;
}
