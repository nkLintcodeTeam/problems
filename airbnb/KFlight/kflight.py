# -*- coding: utf-8 -*-

"""
-1表示无穷大
"""

import collections

class Solution:
    """
    n: the number of airport
    flights: distances between airports
    src: source airport
    dst: destination airport
    K: the number of transfer station
    """
    def findCheapestPrice(self, n, flights, src, dst, K):
        lookup=collections.defaultdict()
        for i in range(n):
            lookup[i]=-1
        # 从src出发，飞0次，显然只能到src，且没有花费
        lookup[src]=0

        for i in range(K+1):
            curLookup=lookup.copy()
            for s, d, c in flights:
                curLookup[d]=self.getMin(curLookup[d], self.add(lookup[s], c))
            lookup=curLookup

        return lookup[dst]

    def getMin(self, a, b):
        if a==-1:
            return b
        if b==-1:
            return a
        if a<b:
            return a
        return b
    def add(self, a, b):
        if a==-1 or b==-1:
            return -1
        return a+b
"""
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200
"""
n = 3
edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0
dst = 2
k = 1
s=Solution()
print s.findCheapestPrice(n, edges, src, dst, k)
