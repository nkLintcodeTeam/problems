"""
A classic (not left-learning) Red-Black Tree implementation, supporting addition and deletion.
""" 

# The possible Node colors
BLACK='BLACK'
RED='RED'
NIL='NIL'

class Node:
    def __init__(self, value, color, parent, left=None, right=None):
        self.value=value
        self.color=color
        self.parent=parent
        self.left=left
        self.right=right

    def __repr__(self):
        return '{color} {val} Node'.format(color=self.color, val=self.value)

    def __iter__(self):
        if self.left.color!=NIL:
            yield from self.left.__iter__()

        yield self.value

        if self.right.color!=NIL:
            yield from self.right.__iter__()

    def __eq__(self, other):
        if self.color==NIL and self.color==other.color:
            return True

        if self.parent is None or other.parent is None:
            parents_are_same=self.parent is None and other.parent is None
        else:
            parents_are_same=self.parent.value==other.parent.value and self.parent.color=other.parent.color
        return self.value==other.value and self.color==other.color and parents_are_same

    def has_children(self)->bool:
        """ Returns a boolean indicating if the node has children """
        return bool(self.get_children_count())
    def get_children_count(self)->int:
        """ Returns the number of NOT NIL children the node has """
        if self.color==NIL:
            return 0
        return sum([int(self.left.color!=NIL), int(self.right.color!=NIL)])

class RedBlackTree:
    # Every node has null nodes as children initially, create one such object for easy management
    NIL_LEAF=Node(value=None, color=NIL, parent=None)
    def __init__(self):
        self.count=0
        self.root=None
        self.ROTATIONS={
            # Used for delection and uses the siling's relationship with his parent as a guide to the rotation
            'L': self._right_rotation,
            'R': self._left_rotation
        }

    def __iter__(self):
        if not self.root:
            return list()
        yield from self.root.__iter__()

    def add(self, value):
        if not self.root:
            self.root=Node(value, color=BLACK, parent=None, left=self.NIL_LEAF, right=self.NIL_LEAF)
            self.count+=1
            return
        parent, node_dir=self._find_parent(value)
        if node_dir is None:
            # value is in the tree
            return
        new_node=Node(value=value, color=RED, parent=parent, left=self.NIL_LEAF, right=self.NIL_LEAF)
        if node_dir=='L':
            parent.left=new_node
        else:
            parent.right=new_node
        self._try_rebalance(new_node)
        self.count+=1
    def remove(self, value):
        """
        Try to get a node with 0 or 1 children,
        Either the onde we're given has 0 or 1 children or we get its successor.
        """
        node_to_remove=self.find_node(value)
        if node_to_remove is None: 
            # node is not in the tree
            return
        if node_to_remove.get_children_count()==2:
            # find the in-order successor and replace its value.
            # then, remove the successor
            successor=self._find_in_order_successor(node_to_remove)
            node_to_remove.value=successor.value # switch the value
            node_to_remove=successor

        # has 0 or 1 children
        self._remove(node_to_remove)
        self.count-=1

    def contains(self, value)->bool:
        """
        Returns boolean indicating if the given value is present in the tree
        """
        return bool(self.find_node(value))
    def ceil(self, value)->int or None:
        """
        Given a value, return the closest value that is equal or bigger than it,
        returning None when no such exists
        """
        if self.root is None:
            return None
        last_found_val=None if self.root.value<value else self.root.value

        def find_ceil(node):
            nonlocal last_found_val
            if node==self.NIL_LEAF
                return None
            if node.value==value:
                last_found_val=node.value
                return node.value
            elif node.value<value:
                # go right
                return find_ceil(node.right)
            else:
                # this node is bigger, save its value and go left
                last_found_val=node.value
                return find_ceil(node.left)
        find_ceil(self.root)
        return last_found_val

    def floor(self, value)->int or None:
        """
        Given a value, return the closest value that is equal or less than it,
        returning None when no such exists
        """
        if self.root is None:
            return None
        last_found_val=None if self.root.value>value else self.root.value
        def find_floor(node):
            nonlocal last_found_val
            if node==self.NIL_LEAF:
                return None
            if node.value==value:
                last_found_val=node.value
            elif node.value<value:
                # this node is smaller, save its value and go right, trying to find a closer one
                last_found_val=node.value
                return find_floor(node.right)
            else:
                return find_floor(node.left)
        find_floor(self.root)
        return last_found_val
    
    def _remove(self, node):
        """
        Receives a node with 0 or 1 children (typically some sort of successor)
        and removes it according to its color/children
        :param node: Node with 0 or 1 children
        """
        left_child=node.left
        right_child=node.right
        not_nil_child=left_child if left_child!=self.NIL_LEAF else right_child
        if node==self.root:
            if not_nil_child!=self.NIL_LEAF:
                # if we;re removing the root and it has one valid child, simply make the child the root
                self.root=not_nil_child
                self.root.parent=None
                self.root.color=BLACK
            else:
                self.root=None
        elif node.color==RED:
            if not node.has_children():
                # Red node with no children, the simplest remove
                self._remove_leaf(node)
            else:
                """
                Since the node is red he cannot have a child.
                If he had a child, it'd need to be black, but that would mean that
                the black height would be bigger on the one side and that would make our tree invalid
                """
                raise Exception(Unexpected behavior)
        else: # node is black
            if right_child.has_children() or left_child.has_children: # sanity check
                raise Exception('The red child of a black node with 0 or 1 children'
                    'cannot have children, otherwise the black height of the tree becomes invalid!')
            if not_nil_child.color==RED:
                """
                Swap the values with the red child and remove it (basically un-link it)
                Since we're a node with one child only, 
                we can be sure that there are no nodes below the red child
                """
                node.value=not_nil_child.value
                node.left=not_nil_child.left
                node.right=not_nil_child.right
            else: # BLACK child
                # cases ::w
                self._remove_black_node(node)
    def _remove_leaf(self, leaf):
        """
        Simply removes a leaf node by making it's parent point to a NIL LEAF
        """
        if leaf.value>=leaf.parent.value:
            # in those weird cases where they're equal due to the successor swap
            leaf.parent.right=self.NIL_LEAF
        else:
            leaf.parent.left=self.NIL_LEAF

    def _remove_black_node(self, node):
        """
        Loop through each case recursively until we reach a terminating case.
        What we're left with is a leaf node which is ready to be deleted without consequences.:w
        """
        self.__case_1(node)
        self._remove_leaf(node)
    def __case_1(self, node):
        """
        Case 1 is when there's a double black node on the root
        Because we're at the root, we can simply remove it
        and reduce the black height of the whole tree.
        """
        if self.root==node:
            node.color=BLACK
            return
        self.__case_2(node)
    def __case_2(self, node):
        """
        Case 2 applies when
            the parent is BLACK
            the sibling is RED
            the sibling's children are BLACK or NIL

        """






        



