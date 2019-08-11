```C++
// An efficient program to randomly select 
// k items from a stream of items  
#include <bits/stdc++.h> 
#include <time.h>  
using namespace std;  
  
// A utility function to print an array  
void printArray(int stream[], int n)  
{  
    for (int i = 0; i < n; i++)  
        cout << stream[i] << " ";  
    cout << endl; 
}  
  
// A function to randomly select  
// k items from stream[0..n-1].  
void selectKItems(int stream[], int n, int k)  
{  
    int i; // index for elements in stream[]  
  
    // reservoir[] is the output array. Initialize  
    // it with first k elements from stream[]  
    int reservoir[k];  
    for (i = 0; i < k; i++)  
        reservoir[i] = stream[i];  
  
    // Use a different seed value so that we don't get  
    // same result each time we run this program  
    srand(time(NULL));  
  
    // Iterate from the (k+1)th element to nth element  
    for (; i < n; i++)  
    {  
        // Pick a random index from 0 to i.  
        int j = rand() % (i + 1);  
  
        // If the randomly picked index is smaller than k,  
        // then replace the element present at the index  
        // with new element from stream  
        if (j < k)  
        reservoir[j] = stream[i];  
    }  
  
    cout << "Following are k randomly selected items \n";  
    printArray(reservoir, k);  
}  
  
// Driver Code 
int main()  
{  
    int stream[] = {1, 2, 3, 4, 5, 6,  
                    7, 8, 9, 10, 11, 12};  
    int n = sizeof(stream)/sizeof(stream[0]);  
    int k = 5;  
    selectKItems(stream, n, k);  
    return 0;  
}  
  
// This is code is contributed by rathbhupendra 
```

```
有一串数字(在线)，不知道有多少，一个一个的出现。
随机选择10个(不够10个，全取)

对于1-10个，数据，直接存到最终结果里
对于第11个数据，我们生成一个随机数rd=rand(1, 11)
如果rd=11,很不幸，我们不选他
如果rd<=10, 我们把他地换掉第rd个数。

对于第11个数而言:
1/11的概率不取， 10/11的概率取

对于前10个数而言
取的概率是:
P(当第11个数不取)+P(当第11个数取)
=1/11 + 10/11 * 9/10
=1/11 + 9/11
=10/11

所有的数都是10/11，所以成立。

对于第12个数，
2/12的概率不取， 10/12的概率取

假设前n-1个数成立，即对于每个前n-1个数，留下的概率是 k/(n-1)

对于第n个数
取的概率是 k/n
对于前k个数，取的概率是
 P(k个数保留到现在)*(P(当不取第n个数)+P(当取第n个数))
=(k/(n-1)) * ( (n-k)/n + (k/n)*((k-1)/k) )
=k/n



```
