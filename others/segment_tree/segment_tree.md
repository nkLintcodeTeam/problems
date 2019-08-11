```
Operations:
  build(start, end, vals)  O(n)
  update(index, value) O(logn)
  rangeQuery(start, end) O(logn+k) 
    where k is the number of reported segments

struct SegmentTreeNode {
  int start;
  int end;
  int sum; // can be max/min
  SegmentTreeNode* left;
  SegmentTreeNode* right;
}

def buildTree(start, end, vals):
  if start==end:
    return SegmentTreeNode(start, end, vals[start])
  mid=(start+end)/2
  left=buildTree(start, mid, vals)
  right=buildTree(mid+1, end, vals)
  return SegmentTreeNode(start, end, left.sum+right.sum, left, right)

def updateTree(root, index, val):
  if root.start==root.end==index:
    root.sum=val
    return
  mid=(root.start+root.end)/2
  if index<=mid:
    updateTree(root.left, index, val)
  else:
    updateTree(root.rifht, index, val)
  root.sum=root.left.sum+root.right.sum
  // T(n)=T(n/2)+O(1)=log(n)

def querySum(root, i, j):
  if root.start==i and root.end==j:
    return root.sum
  mid=(start+end)/2
  if j<=mid:
    return querySum(root.left, i, j)
  else if i>mid;
    return querySum(root.right, i, j)
  else:
    return querySum(root.left, i, mid)+
      querySum(root.right, mid+1, j)

https://www.youtube.com/watch?v=rYBtViWXYeI

```

