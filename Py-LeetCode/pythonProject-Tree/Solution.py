from typing import Optional, List
from Tree import TreeNode
from collections import defaultdict
from collections import deque

class Solution:
    def __init__(self):
        # Variable to store LCA node.
        self.ans = None

    def max_path_sum(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        self.max_sum = root.val
        self.dfs(root)
        return self.max_sum

    def dfs(self, node):
        if not node:
            return 0
        if not node.left and not node.right:
            self.max_sum = max(self.max_sum, node.val)
            return node.val
        left_gain = self.dfs(node.left)
        right_gain = self.dfs(node.right)
        self.max_sum = max(self.max_sum,
                           node.val,
                           node.val + left_gain + right_gain,
                           node.val + left_gain,
                           node.val + right_gain)
        return max(node.val,
                   node.val + right_gain,
                   node.val + left_gain,
                   0)

    def ladder_length(self, beginWord: str, endWord: str, wordList: list[str]) -> int:
        if endWord not in wordList or not beginWord or not endWord or not wordList:
            return 0
        all_combo = defaultdict(list)
        for word in wordList:
            for index in range(len(word)):
                combo_word = word[:index] + "*" + word[index + 1:]
                all_combo[combo_word].append(word)
        visited = defaultdict(bool)
        visited[beginWord] = True
        q = deque([(beginWord, 1)])
        while q:
            curr_node = q.popleft()
            curr_word = curr_node[0]
            curr_level = curr_node[1]
            for index in range(len(curr_word)):
                combo_word = curr_word[:index] + "*" + curr_word[index + 1:]
                if combo_word in all_combo.keys():
                    combo_values = all_combo[combo_word]
                    for word in combo_values:
                        if word == endWord:
                            return curr_level + 1
                        if word not in visited.keys():
                            q.append((word, curr_level + 1))
                            visited[word] = True
        return 0

    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # adj = [[] for _ in range(numCourses)]
        # for prerequisite in prerequisites:
        #     adj[prerequisite[1]].append(prerequisite[0])
        # visited = [False] * numCourses
        # stack = [False] * numCourses
        #
        # def dfs(course):
        #     if stack[course]:
        #         return True
        #     if visited[course]:
        #         return False
        #     visited[course] = True
        #     stack[course] = True
        #     for neighbour in adj[course]:
        #         if dfs(neighbour):
        #             return True
        #     stack[course] = False
        #     return False
        #
        # for crs in range(numCourses):
        #     if dfs(crs):
        #         return False
        # return True
        adj = [[]for _ in range(numCourses)]
        indegree = [0]*numCourses
        for prerequisite in prerequisites:
            adj[prerequisite[1]].append(prerequisite[0])
            indegree[prerequisite[0]] += 1

        queue = deque()
        for i in range(numCourses):
            if indegree[i] == 0:
                queue.append(i)
        node_visited = 0
        while queue:
            item = queue.popleft()
            node_visited += 1

            for neighbour in adj[item]:
                indegree[neighbour] -= 1
                if indegree[neighbour] == 0:
                    queue.append(neighbour)
        return node_visited == numCourses

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # if not root or root == p or root == q:
        #     return root
        # left = self.lowestCommonAncestor(root.left, p, q)
        # right = self.lowestCommonAncestor(root.right, p, q)
        # if left and not right:
        #     return root
        # elif left and right:
        #     return left
        # else:
        #     return right
        def recurse_tree(current_node: TreeNode) -> bool:

            # If reached the end of a branch, return False.
            if not current_node:
                return False

            # Left Recursion
            left = recurse_tree(current_node.left)

            # Right Recursion
            right = recurse_tree(current_node.right)

            # If the current node is one of p or q
            mid = current_node == p or current_node == q

            # If any two of the three flags left, right or mid become True.
            if mid + left + right >= 2:
                self.ans = current_node

            # Return True if either of the three bool values is True.
            return mid or left or right

        recurse_tree(root)
        return self.ans

    def floodFill(self, image: list[list[int]], sr: int, sc: int, color: int) -> list[list[int]]:
        if image[sr][sc] == color:
            return image
        prev_color = image[sr][sc]
        m, n = len(image), len(image[0])

        def doFloodFill(i: int, j: int):
            if image[i][j] == prev_color:
                image[i][j] = color
            else:
                return
            if i + 1 < m:
                doFloodFill(i+1, j)
            if i - 1 >= 0:
                doFloodFill(i-1, j)
            if j + 1 < n:
                doFloodFill(i, j + 1)
            if j - 1 >= 0:
                doFloodFill(i, j - 1)

        doFloodFill(sr, sc)

        return image



















T9 = TreeNode(4, None, None)
T8 = TreeNode(7, None, None)
T7 = TreeNode(8, None, None)
T6 = TreeNode(0, None, None)
T5 = TreeNode(2, T8, T9)
T4 = TreeNode(6, None, None)
T3 = TreeNode(1, T6, T7)
T2 = TreeNode(5, T4, T5)
T1 = TreeNode(3, T2, T3)
# T3 = TreeNode(3, None, None)
# T2 = TreeNode(2, None, None)
# T1 = TreeNode(1, T2, T3)

# T1 = TreeNode(-3, None, None)
# T2 = TreeNode(1, None, None)
# T1 = TreeNode(-2, T2, None)
Sol = Solution()
# print(Sol.max_path_sum(T1))
# print(Sol.ladder_length("hit", "cog", ["hot","dog","tog","cog"]))
# print(Sol.canFinish(5, [[1,4],[2,4],[3,1],[3,2]]))
# print(Sol.canFinish( 2,[[1,0],[0,1]]))
# print(Sol.lowestCommonAncestor(T1,T2,T3).val)
print(Sol.floodFill([[0, 0, 0],[0, 1, 0]], 1, 1, 2))
