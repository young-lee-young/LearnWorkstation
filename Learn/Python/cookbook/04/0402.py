# 委托迭代
# 逐步调试代码,注意代码执行的步骤

class Node:
    def __init__(self, values):
        self._values = values
        self._children = []


    def __repr__(self):
        return 'Node({!r})'.format(self._values)


    def add_child(self, node):
        self._children.append(node)


    def __iter__(self):
        return iter(self._children)


if __name__ == '__main__':
    root = Node(0)
    child1 = Node(1)
    child2 = Node(2)
    root.add_child(child1)
    root.add_child(child2)
    for ch in root:
        print(ch)