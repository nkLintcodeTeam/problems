#include<iostream>
#include<vector>
using namespace std;
vector<int> res; //组合结果 (栈存储) 
int N; //集合元素个数N
int M; //取出元素个数M
int *a; //集合A (数组存储)
void print() { //输出组合子集的结果 
	for (int i = 0; i < res.size(); ++i) {
		cout << res[i] << ' ';
	}
	cout << endl;
}
void combination(int k, int m) { //组合递归函数，其中k为标志元素的下标索引，m为待取的元素个数 
	if (res.size() == M) { //终止条件：取出元素个数达到m个
		print(); //输出一种组合结果 
	} else {
		if (k < N) {
			res.push_back(a[k]); //取出标志元素存入栈中
			combination(k + 1, m - 1); //取标志元素后，在剩余的n-1个元素中取m-1个元素
			res.pop_back(); //标志元素出栈
			combination(k + 1, m); //不取标志元素，在剩余的n-1个元素中取m个元素
		}
	}
}
int main() {
	cin >> N >> M; //输入N, M
	a = new int[N];
	for (int i = 0; i < N; a[i++] = i ) {} //初始化集合A为1-N的整型数值元素 
	combination(0, M); //初始标志元素为a[0]，待取元素个数为m个
	delete a; 
	return 0;
}
