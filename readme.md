# 20190423计划

[TOC]

## 背景
工作之余，练练脑子。

日积月累，水滴石穿。

## 题目列表

优先做高频题，重复题目，写好指向哪里就成，不需要再写一遍。(leetcode与lintcode题目大部分重复，各个公司的题目也有来自leetcode/lintcode的)

解题报告，可以附上比较好的讲解地址(博客，视频)，就不用自己写了

最好有能运行的正确代码，语言不限。(看代码可能比看博客更好)

代码不一定对哈

#### 算法导论
书中觉得重要的算法写一写，按照章节安排目录
``` 
introduction_to_algorithms
    4_DivideAndConquer
    6_HeapSort wncbb Doing
    7_QuickSort wncbb Done
    15_DynamicProgramming
    16_GreedyAlgorithms
    22_ElementaryGraphAlgorithms
        22_2_bfs wncbb Done
        22_3_dfs wncbb Doing
        22_4_topologicalSort
        22_5_stronglyConnectedComponents wncbb Doing
    24_SingleSourceShortestPaths
        24_1_bellmanFord wncbb doing
        24_3_dijkastra wncbb doing
```


#### leetcode题目
```
leetcode
```

#### lintcode题目
```
lintcode
    10_RegularExpressionMatching wncbb Doing 
    11_BinaryTreeMaximumPathSum wncbb Doing  
```

#### airbnb题目
```
Airbnb
    A
    B
    C
        Calculator wncbb doing
    D
        DisplayPage wncbb doing
    E
    F
        FileSystem wncbb doing
        FindMedianInLargeFileOfIntegers wncbb doing
    G
    H
    I
        IPRangeToCIDR wncbb doing
    J
    K
        KEditDistance wncbb doing
    L
    M
        MeetingTime wncbb doing
        MenuCombinationSum wncbb doing
        MinimumCostWithAtMostKStop wncbb doing
    N
    O
    P
    Q
    R
        RoundPrices wncbb doing
    S
        SlidingGame wncbb doing
    T
        TextJustification wncbb doing
    U
    V
    W
        WaterDrop wncbb doing
        Wizards wncbb doing
    X
    Y
    Z
```

#### others

Union Find Sets  wncbb done

## 理论

#### 一致性(强，弱，最终)

#### 操作系统常见概念
1. 锁实现(单核锁，多核锁)

#### 网络常见概念
1. http2.0
2. 流量控制与拥塞控制

#### 微服务环境下，数据一致性方案(各大公司方案)

#### 一致性hash算法

## 开源项目源码阅读

#### go

1. slice底层实现
2. map底层实现
3. 内存分配(类似tcmalloc, 顺带jemalloc)
4. gc机制
5. 协程机制，比如调度等
6. channel实现
7. interface实现
   
#### leveldb

顺带把常见的系统函数过一遍

#### redis

1. 集群方案, 扩容缩容(面试问的很多很细)
2. 顺带把select/poll/epoll的内容过一遍

#### mysql
1. 事务过程
2. 锁实现
3. 性能调优

#### raft
比如etcd, consul

raft算法go实现

[https://github.com/hashicorp/raft](https://github.com/hashicorp/raft)

#### kafka/tmq
1. 时间轮
2. at least once, at most once, exactly once

#### nginx
1. 多进程数据交互方式

#### mapreduce/spark/flink基础的概念

#### HDFS/HBase/Hive/TiDB基础的概念

## fork test
