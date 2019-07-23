# -*- encoding: utf-8 -*-

import collections

class TrieNode:
    def __init__(self, value):
        self.isWord=False
        self.value=value
        self.children=collections.OrderedDict()
        self.word=''
        
    @classmethod
    def insert(cls, root, word):
        p=root
        for ch in word:
            child=p.children.get(ch)
            if child is None:
                child=TrieNode(ch)
                p.children[ch]=child
            p=p.children[ch]
        p.isWord=True
        p.word=word

# 层级遍历
words=['ae', 'abc', 'abd', 'xz']
root=TrieNode('')
for w in words:
    TrieNode.insert(root, w)

q=collections.deque()
q.append(root)
rst=''
while q:
    tmpq=collections.deque()
    while q:
        cur=q.popleft()
        for ch in cur.children:
            tmpq.append(cur.children[ch])
            rst+='-'+ch
    q=tmpq
    rst+="\n"

print rst
    




