

算法导论里，节点序列从1开始计数。
但是编程，一般从0开始计数。
这里采用从0开始计数的方式。因此父节点与左右子节点的索引计算方式与算法导论书中的不太一样。
最简单的，考虑0，1，2节点，parent(1)=(1-1)/2。
下面的伪代码全是从0开始计数的
```
PARENT(i)
    return (i-1)/2

LEFT(i)
    return i*2+1

RIGHT(i)
    return i*2+2

MAX-HEAPIFY(A, i)
    l=LEFT(i)
    r=RIGHT(i)
    largest=i

    if l<A.heap-size and A[l]>A[largest]
        largest=l
    if r<A.heap-size and A[r]>A[largest]
        largest=r
    
    if i!=largest
        exchange A[i] with A[largest]
        MAX-HEAPIFY(A, largest)

BUILD-MAX-HEAP(A)
    A.heap-size=A.length
    for i=A.length/2 down to 0
        MAX-HEAPIFY(A, i)

HEAQPSORT(A)
    BUILD-MAX-HEAP(A)
    for i=A.length-1 downto 1
        exchange A[0] with A[i]
        A.heap-size=A.heap-size - 1
        MAX-HEAPIFY(A, 0)

# 以上便是堆排序所需要的代码
# 下面是其他的堆操作代码，比如peek, insert, pop

# peek
HEAP-MAXINUM(A)
    return A[0]

# pop
HEAP-EXTRACT-MAX(A)
    if A.heap-size<=0
        error "heap underflow"
    max=A[0]
    A[0]=A[A.heap-size - 1]
    A.heap-size=A.heap-size - 1
    MAX-HEAPIFY(A, 0)
    return max

# insert
# 这个没有用算法导论的方法
# 新插入元素，放到最后，然后把它升上去
HEAP-INSERT(A, value)
    A.heap-size=A.heap-size+1
    A[A.heap-size - 1]=value
    while i>0
        parent=PARENT(i)
        if A[i]>A[parent]
            exchange A[i] with A[parent]
            i=parent


```