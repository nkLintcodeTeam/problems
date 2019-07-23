# -*- encoding: utf-8 -*-

import collections

class TrieNode:
    def __init__(self, value):
        self.value=value
        self.children=collections.OrderedDict()
        self.word=''
        self.isWord=False

    @classmethod
    def insert(cls, root, word):
        p=root
        for ch in word:
            child=p.children.get(ch)
            if child is None:
                child=TrieNode(ch)
                p.children[ch]=child
            p=p.children[ch]
        p.word=word
        p.isWord=True

class Solution:
    def wordSearchII(self, board, words):
        self.board=board
        trieRoot=TrieNode('')
        for word in words:
            TrieNode.insert(trieRoot, word)
        self.result=set()

        # len(board)*len(board[0])*4directions*max(one of words)
        for i in range(len(board)):
            for j in range(len(board[0])):
                self.dfs(i, j, trieRoot, '')

        return list(self.result)
               
    def dfs(self, i, j, trieNode, preStr):
        # must i>=len(self.board)
        # using or , not ||
        if i<0 or i>=len(self.board) or j<0 or j>=len(self.board[0]):
            return

        curCh=self.board[i][j]

        # already visited
        if curCh=='#': 
            return

        # if this is not a word, just return
        curNode=trieNode.children.get(curCh)
        if curNode is None:
            return

        curStr=preStr+self.board[i][j]

        if curNode.isWord:
            self.result.add(curNode.word)

        dirs=[
            [0, 1],
            [1, 0],
            [0, -1],
            [-1, 0],
        ]
        oldCh=self.board[i][j]
        # using 'for dx, dy in dirs:', not 'for dx, dy=dirs:'
        for dx, dy in dirs:
            x=i+dx
            y=j+dy
            self.board[i][j]='#'
            self.dfs(x, y, curNode, curStr)
            self.board[i][j]=oldCh

board=[
    ['d', 'o', 'a', 'f'],
    ['a', 'g', 'a', 'i'],
    ['d', 'c', 'a', 'n'],
]
words=["doa", "dog", "dad", "dgdg", "can", "again"]
s=Solution()
print s.wordSearchII(board, words)

        






         



