# -*- encoding: utf-8 -*-
from heapq import heappop, heappush, heapify

class Solution:
    def alienOrder(self, words):
        graph=self.buildGraph(words)
        return self.topological_sort(graph)
    
    def buildGraph(self, words):
        graph={}
        for word in words:
            for c in word:
                if c not in graph:
                    graph[c]=set()

        for i in range(len(words)-1):
            for j in range(min(len(words[i]), len(words[i+1]))):
                if words[i][j]!=words[i+1][j]:
                    graph[words[i][j]].add(words[i+1][j])
                    # must break, because the rest means nothing
                    break

        return graph

    def topological_sort(self, graph):
        # build indegree map
        indegree={
            ch: 0
            for ch in graph
        }

        for ch in graph:
            for neighbor in graph[ch]:
                indegree[neighbor]+=1

        q=[ch for ch in indegree if indegree[ch]==0]
        heapify(q)

        res=''
        while q:
            ch=heappop(q)
            res+=ch
            for neighbor in graph[ch]:
                indegree[neighbor]-=1
                if indegree[neighbor]==0:
                    heappush(q, neighbor)

        if len(res)==len(graph): 
            return res
        return ''


s=Solution()

words=["ze","yf","xd","wd","vd","ua","tt","sz","rd", "qd","pz","op","nw","mt","ln","ko","jm","il", "ho","gk","fa","ed","dg","ct","bb","ba"]
print s.alienOrder(words)
words=['ac', 'ab']
print s.alienOrder(words)




        
