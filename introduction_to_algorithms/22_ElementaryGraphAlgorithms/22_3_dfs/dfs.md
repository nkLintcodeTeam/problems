

深度优先遍历，要么用递归，要么用栈。

递归也可以看成是栈，函数调用栈。

递归解法
```
DFS(G)
    for each u in G.V
        u.color=WHITE
        u.parent=nil

    time=0

    for each u in G.V
        if u.color==WHITE
            DFS-VISIT(G, u)

DFS-VISIT(G, u)
    time=time+1
    u.displayTime=time
    u.color=GRAY
    for each v in G.Adj[u]
        if v.color==WHITE
            v.parent=u
            DFS-VISIT(G, v)
    u.color=BLACK
    time=time+1
    u.finishTime=time

```

