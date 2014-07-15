import sys


class Node:
    value = None
    left = None
    right = None

    def __init__(self, val):
        self.value = val


def insert(value, tree):
    if tree is None:
        return Node(value)

    if value <= tree.value:
        tree.left = insert(value, tree.left)
    else:
        tree.right = insert(value, tree.right)

    return tree


def lca(a, b, tree):
    # If both numbers are smaller, go left
    # If both numbers are greater, go right
    if a <= tree.value and b <= tree.value:
        if a == tree.value or b == tree.value:
            return tree
        else:
            return lca(a, b, tree.left)
    elif a > tree.value and b > tree.value:
        return lca(a, b, tree.right)
    else:
        return tree


if __name__ == "__main__":
    tree = None
    for x in [30, 8, 52, 3, 20, 10, 29]:
        tree = insert(x, tree)

    f = open(sys.argv[1])
    for l in f:
        a, b = tuple(int(x) for x in l.split())
        print lca(a, b, tree).value
