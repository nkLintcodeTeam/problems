class UnionFindSet:
    def __init__(self, n):
        self.parents=[]
        self.ranks=[]
        for i in range(n):
            self.parents.append(i)
            self.ranks.append(0)

    def connected(self, a, b):
        return self.find(a)==self.find(b)

    def find(self, x):
        if x!=self.parents[x]:
            self.parents[x]=self.find(self.parents[x])
        return self.parents[x]

    def union(self, x, y):
        px=self.find(x)
        py=self.find(y)
        if self.ranks[px]>self.ranks[py]:
            self.parents[py]=px
        if self.ranks[px]<self.ranks[py]:
            self.parents[px]=py
        if self.ranks[px]==self.ranks[py]:
            self.parents[py]=px
            self.ranks[px]+=1

t=UnionFindSet(5)
print t.find(0)
print t.find(1)

t.union(0, 1)
print t.find(0)
print t.find(1)

t.union(2, 3)
t.union(0, 2)

print t.parents
print t.ranks

print t.connected(0, 3)


