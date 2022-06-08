# -*- coding: utf-8 -*-
# @Time:    2020/3/15 6:48 PM
# @Author:  leeyoung
# @File:    avl_tree.py
# @Content: 平衡二叉查找树

from random import shuffle
from tree_node import AVLTreeNode


# 解决树退化为链表的问题
# 优点：左右子树差最大为1，查询效率优于红黑树
# 缺点：由于要维持树的高度平衡，插入和删除需要进行大量再平衡操作，会损耗一定性能，比红黑树要差
# 如果没有大量插入、删除操作，查询操作比较多，可以用AVL Tree，如果插入、删除、查询操作次数近似，使用红黑树
class AVLTree:
    def __init__(self):
        self.root = None
        self.tree_size = 0

    def size(self):
        return self.tree_size

    def is_empty(self):
        return True if self.size() == 0 else False

    # 获取节点的高度
    def height(self, node):
        if node is None:
            return 0
        return node.height

    # 计算节点node平衡因子，左右子树高度差
    def balance(self, node):
        if node is None:
            return 0
        return self.height(node.left) - self.height(node.right)

    # 判断是否是二分搜索树，利用二分搜索树的中序遍历有序特性
    def is_bst(self):
        data_list = list()
        # 对二叉树进行中序遍历，将所有值放入列表中
        self.inorder_traversal(self.root, data_list)

        # 循环列表，看是否有序
        last_num = data_list[0]
        for current_num in data_list[1:]:
            if current_num < last_num:
                return False
            last_num = current_num
        return True

    # 中序遍历递归
    def inorder_traversal(self, node, data_list):
        if node is None:
            return

        # 先左子树，再根节点，最后左子树
        self.inorder_traversal(node.left, data_list)
        data_list.append(node.data)
        self.inorder_traversal(node.right, data_list)

    # 判断是否是平衡二叉树
    def is_balanced(self):
        return self._is_balanced(self.root)

    def _is_balanced(self, node):
        # 空树是平衡二叉树
        if node is None:
            return True

        # 获取当前节点平衡因子
        balance = self.balance(node)
        # 平衡因子大于1，不是平衡二叉树
        if abs(balance) > 1:
            print(balance)
            return False
        # 递归左子树和右 子树
        return self._is_balanced(node.left) and self._is_balanced(node.right)

    # 右旋转，对节点y进行右旋转，返回旋转后的新根节点x
    #        y                         x
    #       / \                      /   \
    #      x   T4                   z     y
    #     / \       - - - - ->     / \   / \
    #    z   T3                   T1 T2 T3 T4
    #   / \
    # T1   T2
    def _right_rotate(self, node):
        # 缓存左子树x
        new_root = node.left
        # 缓存左子树x的右节点T3
        node_left_right = new_root.right

        # 将新根节点的右子树设置为原根节点，将x的右子树设置为y
        new_root.right = node
        # 将原根节点的左子树设置为原左子树的右节点，将y的左子树设置为T3
        node.left = node_left_right

        # 更新高度值，更新y和x的高度
        node.height = max(self.height(node.left), self.height(node.right)) + 1
        new_root.height = max(self.height(new_root.left), self.height(new_root.right)) + 1

        return new_root

    # 左旋转，对节点y进行左旋转，返回旋转后的新的根节点x
    #     y                        x
    #    / \                     /   \
    #   T1  x                   y     z
    #      / \     - - - - ->  / \   / \
    #     T2  z               T1 T2 T3 T4
    #        / \
    #       T3 T4
    def _left_rotate(self, node):
        # 缓存右子树x
        new_root = node.right
        # 缓存右子树x的左节点T2
        node_right_left = new_root.left

        # 将新根节点的左子树设置为原根节点，将x的左子树设置为y
        new_root.left = node
        # 将原根节点的右子树设置为原右子树的左节点，将y的右子树设置为T2
        node.right = node_right_left

        # 更新树高度值，更新y和x的高度
        node.height = max(self.height(node.left), self.height(node.right)) + 1
        new_root.height = max(self.height(new_root.left), self.height(new_root.right)) + 1

        return new_root

    # 添加一个节点，递归方法
    def add_node(self, num):
        self.root = self._add_node(self.root, num)

    def _add_node(self, node, num):
        # 当前节点为空，创建新节点，返回创建后的节点
        if not node:
            self.tree_size += 1
            node = AVLTreeNode(num)
            return node

        # 往左子树添加节点
        if num < node.data:
            node.left = self._add_node(node.left, num)
        # 往右子树添加节点
        if num > node.data:
            node.right = self._add_node(node.right, num)

        # ---------------- 以上为平衡树常规添加节点操作，以下为AVL树再平衡操作 ------------------

        # 更新当前节点高度，高度为左右孩子最大高度加1
        node.height = 1 + max(self.height(node.left), self.height(node.right))

        # 获取当前节点平衡因子
        balance = self.balance(node)

        # 不平横，进行旋转
        # LL，平衡因子大于1，当前以node为根树不平衡；左子树的平衡因子大于等于0，说明是在当前节点左子树左边插入了一个节点
        if balance > 1 and self.balance(node.left) >= 0:
            # 右旋
            return self._right_rotate(node)
        # RR，平衡因子小于-1，当前以node为根树不平衡；右子树的平衡因子小于等于0，说明是在当前节点右子树右边插入了一个节点
        if balance < -1 and self.balance(node.right) <= 0:
            # 左旋
            return self._left_rotate(node)
        # LR，平衡因子大于1，当前以node为根树不平衡；左子树的平衡因子小于0，说明是在当前节点左子树右边插入了一个节点
        if balance > 1 and self.balance(node.left) < 0:
            # 左子树左旋，变成LL
            node.left = self._left_rotate(node.left)
            return self._right_rotate(node)
        # RL，平衡因子小于-1，当前以node为根树不平衡，右子树的平衡因子大于0，说明是在当前节点右子树左边插入了一个节点
        if balance < -1 and self.balance(node.right) > 0:
            # 右子树右旋，变成RR
            node.right = self._right_rotate(node.right)
            return self._left_rotate(node)

        # 返回当前节点
        return node

    # 查询元素，时间复杂度O(logn)
    def contains(self, num):
        return self._contains(self.root, num)

    def _contains(self, node, num):
        if not node:
            return False
        elif num == node.data:
            return True
        elif num < node.data:
            return self._contains(node.left, num)
        else:
            return self._contains(node.right, num)

    # 获取最小元素
    def minimum(self):
        if self.tree_size == 0:
            return None
        return self._minimum(self.root).data

    # 最小值在最左边
    def _minimum(self, node):
        while node.left:
            node = node.left
        return node

    # 删除任意节点
    def remove_node(self, num):
        self.root = self._remove_node(self.root, num)

    def _remove_node(self, node, num):
        if not node:
            return
        # 不直接返回node，对node进行再平衡，将返回的node暂存为return_node

        # 待删除的节点在左子树
        if num < node.data:
            node.left = self._remove_node(node.left, num)
            return_node = node
        # 待删除的节点在右子树
        elif num > node.data:
            node.right = self._remove_node(node.right, num)
            return_node = node
        # 当前节点即为删除节点
        else:
            # 待删除的节点只有右子树
            if not node.left:
                right_node = node.right
                node.right = None
                self.tree_size -= 1
                return_node = right_node
            # 待删除的节点只有左子树
            elif not node.right:
                left_node = node.left
                node.left = None
                self.tree_size -= 1
                return_node = left_node
            # 待删除的节点既有左子树，又有右子树
            else:
                successor = self._minimum(node.right)
                # successor.right = self._remove_min(node.right)
                # 这句话和上面这句话效果一样，下面这句话维护了平衡
                successor.right = self._remove_node(node.right, successor.data)
                successor.left = node.left

                # 删除node节点
                node.left = node.right = None
                return_node = successor

        # ---------------- 以上为平衡树常规删除节点操作，以下为AVL树再平衡操作 ------------------

        # 删除节点return_node有可能为空
        if not return_node:
            return return_node

        # 对要返回的节点return_node进行平衡处理
        # 更新当前节点高度，高度为左右孩子最大高度加1
        return_node.height = 1 + max(self.height(return_node.left), self.height(return_node.right))

        # 获取当前节点平衡因子
        balance = self.balance(return_node)

        # 不平横，进行旋转
        # LL，平衡因子大于1，当前以node为根树不平衡；左子树的平衡因子大于等于0，说明是在当前节点左子树左边插入了一个节点
        if balance > 1 and self.balance(return_node.left) >= 0:
            # 右旋
            return self._right_rotate(return_node)
        # RR，平衡因子小于-1，当前以node为根树不平衡；右子树的平衡因子小于等于0，说明是在当前节点右子树右边插入了一个节点
        if balance < -1 and self.balance(return_node.right) <= 0:
            # 左旋
            return self._left_rotate(return_node)
        # LR，平衡因子大于1，当前以node为根树不平衡；左子树的平衡因子小于0，说明是在当前节点左子树右边插入了一个节点
        if balance > 1 and self.balance(return_node.left) < 0:
            # 左子树左旋，变成LL
            return_node.left = self._left_rotate(return_node.left)
            return self._right_rotate(return_node)
        # RL，平衡因子小于-1，当前以node为根树不平衡，右子树的平衡因子大于0，说明是在当前节点右子树左边插入了一个节点
        if balance < -1 and self.balance(return_node.right) > 0:
            # 右子树右旋，变成RR
            return_node.right = self._right_rotate(return_node.right)
            return self._left_rotate(return_node)

        return return_node


avl_tree = AVLTree()
num_list = list(range(10))
shuffle(num_list)

# 添加元素
for num in num_list:
    avl_tree.add_node(num)

# 是否是二分搜索树
print(avl_tree.is_bst())

# 是否是平衡二叉树
print(avl_tree.is_balanced())

# 包含元素
print(avl_tree.contains(1))
print(avl_tree.contains(10))

# 删除任意值
avl_tree.remove_node(7)
print(avl_tree.is_bst())
print(avl_tree.is_balanced())
