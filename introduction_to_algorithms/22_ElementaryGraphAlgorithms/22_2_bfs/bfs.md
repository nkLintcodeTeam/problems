
广度优先遍历，也挺简单的。
主要使用一个FIFO队列。
一个原则：

没看见的元素是白色的。
放到队列里的元素是灰色的。
从队列出来的元素，遍历完children之后改为黑色。

```
BFS(G, s)
    for each u in (G.V-{s})
        u.color=WHITE
        u.d=inf
        u.parent=nil
    
    s.color=GRAY
    s.d=0
    s.parent=nil
    Q{}
    ENQUEUE(Q, s)
    while Q is not empty:
        u=DEQUEUE(Q)
        for each v in G.Adj(u)
            if v.color=WHITE
                v.color=GRAY
                v.d=u.d+1
                v.parent=u
                ENQUEUE(Q, v)
        u.color=BLACK


```