
```
V(i, j) // 对于前i件商店,在背包承重为j的情况下,背包所能带走的最大价值
if wj>j
    V(i, j)=V(i-1, j)
else
    V(i, j)=Max(V(i-1, j), V(i-1, j-wj)+vj)

```


