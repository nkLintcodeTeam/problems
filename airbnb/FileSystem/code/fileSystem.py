import collections

class Node:
    def __init__(self, f, value):
        self.f=f
        self.value=value
        self.callback=None
        self.children=collections.defaultdict()
    # def next(self, value):
    #     if value in self.children:
    #         return self.children[value]
    #     return None
    # def addChild(self, value, child):
    #     self.children[value]=child

class FileSystem:
    def __init__(self):
        self.root=Node('', 0)
    
    def create(self, filePath, value):
        fileList=filePath.split('/')[1:]
        if len(fileList)==1:
            lastFile=fileList[0]
            fileList=[]
        else:
            lastFile=fileList[-1]
            fileList=fileList[:-1]
        print 'fileList', fileList
        curNode=self.root
        for f in fileList:
            if f not in curNode.children:
                return False
            curNode=curNode.children[f]
        curNode.children[lastFile]=Node(lastFile, value)
        return True
    
    def get(self, filePath):
        fileList=filePath.split('/')[1:]
        curNode=self.root
        for f in fileList:
            if f not in curNode.children:
                return -1
            curNode=curNode.children[f]
        
        return curNode.value
    
    def set(self, filePath, value):
        fileList=filePath.split('/')[1:]
        curNode=self.root
        callbackList=[]
        for f in fileList:
            if f not in curNode.children:
                return False
            curNode=curNode.children[f]
            if curNode.callback is not None:
                callbackList.append(curNode.callback)
        curNode.value=value
        while len(callbackList)>0:
            f=callbackList.pop()
            f()

        return True
    def watch(self, filePath, callback):
        fileList=filePath.split('/')[1:]
        curNode=self.root
        for f in fileList:
            if f not in curNode.children:
                return False
            curNode=curNode.children[f]
        curNode.callback=callback
        return True

s=FileSystem()
s.create('/a', 2)
s.create('/a/b', 1)
def abtmp():
    print 'this is /a/b tmp'
def atmp():
    print 'this is /a tmp'
s.watch('/a/b', abtmp)
s.watch('/a', atmp)

s.set('/a/b', 12)
print s.get('/a/b')