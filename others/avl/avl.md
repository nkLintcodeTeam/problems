```
AVL insert
1. Simpole BST insert
2. Fix AVL property
```


```
height of a node:
  length of the longest path from it down to a leaf
  =max{height(left child), height(right child)}+1

require heights of left & right children of every node
to differ by at most +-1

N(h) means the number of nodes for tree with height h
fibonacci  (phi^h)/sqrt(5), phi=1.618...
N(h)=1+N(h-1)+N(h-2)
N(h)>F(h)
N(h)=1+N(h-1)+N(h-2)
    > 1+2N(h-2)
    > 2N(h-2)
    =theta(2^(h/2))
=> h<2lg(n)

```

```
    x
  /   \
 A     y
      / \
     B   C

left-rotate(x): up->down
right-rotte(y): down->up
     y
   /   \
  x     C
 / \   
A   B


- support x is the lowest node violating AVL
- assume right(x) higher
- if x's right child is right-hearvy or balanced
-- Right Rotate(x)
     x
   /   \
  A     y
       / \
      B   C
Right Rotate(x):
     y
    / \
   x   C
  / \
 A   B

priority queue => heap, AVL
we use heap, because heap doesn't need extra space

```


