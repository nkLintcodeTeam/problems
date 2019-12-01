
```
function Dijksgtra(Graph, source):
    create vertex setQ
    for each vertex v in Graph:
        dist[v] <- INFINITY
        prev[v] <- UNDEFINED
        add v to Q

    dist[source] <- 0
    while Q is not empty:
        u <- vertex in Q with min dist[u]
        remove u from Q
        for each neighbor v of u:
            alt <- dist[u]+length(u, v)
            if alt < dist[v]:
                dist[v] <- alt
                prev[v] <- u
    
    return dist[], prev[]

```

```
function Dijkstra(Graph, source):
    dist[source]=0
    create vertex priority queue Q

    for each vertex v in Graph:
        if v!=source:
            dist[v] <- INFINITY
        prev[v] <- UNDEFINED
        Q.add_with_priority(v, dist[v])

    while Q is not empty:
        u <- Q.extract_min()
        for each neighbor v of u:
            alt <- dist[u]+length(u, v)
            if alt < dist[v]:
                dist[v] <- alt
                prev[v] <- u
                Q.decrease_priority(v, alt)

```