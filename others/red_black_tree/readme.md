1. A node is either red or black
2. The root and leaves(NIL) are black.
3. If a node is red, then its children are black.
4. All paths from a node to its NIL descendants contain the same number of black nodes.

1. Nodes require one storage bit to keep track of color.
2. The longest path(root to farthest NIL) is no more than twice the lenght of the shortest path(root to nearest NIL)
 - Shortest path: all black nodes
 - Longest path: alternating red and black

Search O(logn)
Insert O(logn)
Remove O(logn)

Space complexity O(n)

https://github.com/stanislavkozlovski/Red-Black-Tree/blob/master/rb_tree.py


https://www.youtube.com/watch?v=qvZGUFHWChY
