import heapq
from collections import defaultdict

from LL import ListNode
from typing import Optional, List
from LN import Node
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # prev_node = None
        # temp_node = None
        # while head is not None:
        #     temp_node = ListNode(head.val, prev_node)
        #     # temp_node.val = head.val
        #     # temp_node.next = prev_node
        #     prev_node = ListNode(temp_node.val, temp_node.next)
        #     head = head.next
        # return temp_node
        # next implementation without using extra space
        # prev = None
        # while head is not None:
        #     temp_next = head.next
        #     curr = head
        #     curr.next = prev
        #     prev = curr
        #     head = temp_next
        # return prev
        # recursive
        if (not head) or (not head.next):
            return head
        temp_node = self.reverseList(head.next)
        head.next.next = head
        head.next = None
        return temp_node

    def copyRandom(self, head: 'Optional[Node]') -> 'Optional[Node]':
        visited = defaultdict(None)
        # if head is None:
        #     return None
        # node_addr = defaultdict(int)
        # new_node_addr = []
        # random_addr = defaultdict(None)
        # node_random = defaultdict(int)
        # new_head = Node(head.val)
        # curr = new_head
        # index = 0
        # while head:
        #     random_addr[index] = head.random
        #     node_addr[head] = index
        #     new_node = None
        #     if head.next:
        #         new_node = Node(head.next.val)
        #     curr.next = new_node
        #     new_node_addr.append(curr)
        #     curr = curr.next
        #     head = head.next
        #     index += 1
        #
        # for value, key in node_addr.items():
        #     if random_addr[key]:
        #         node_random[key] = node_addr[random_addr[key]]
        #     else:
        #         node_random[key] = -1
        #
        # head = new_head
        # index = 0
        # while head:
        #     if node_random[index] == -1:
        #         head.random = None
        #     else:
        #         head.random = new_node_addr[node_random[index]]
        #     index += 1
        #     head = head.next
        #
        # return new_head

        # def copy(head: 'Optional[Node]')-> -> 'Optional[Node]':
        #     if head is None:
        #         return None
        #     key_exists = visited.get(head, 'missing')
        #     if key_exists != 'missing':
        #         return visited[head]
        #     else:
        #         new_node = Node(head.val)
        #         visited[head] = new_node
        #         new_node.next = copy(head.next, visited)
        #         new_node.random = copy(head.random, visited)
        #         return new_node
        #
        # return copy(head, visited)

        # def copy(head: 'Optional[Node]'):
        #     if not head:
        #         return head
        #     key_exists = visited.get(head, 'missing')
        #     if key_exists == 'missing':
        #         new_node = Node(head.val)
        #         visited[head] = new_node
        #     return visited[head]
        # ptr = head
        # new = Node(ptr.val)
        # visited[head] = new
        # if not head:
        #     return head
        # while ptr:
        #     new.next = copy(ptr.next)
        #     new.random = copy(ptr.random)
        #     ptr = ptr.next
        #     new = new.next
        # return visited[head]
    # third approach
        if not head:
            return head
        ptr = head
        while ptr:
            new = Node(ptr.val)
            new.next = ptr.next
            ptr.next = new
            ptr = new.next
        ptr = head
        while ptr:
            if ptr.random:
                ptr.next.random = ptr.random.next
            else:
                ptr.next.random = None
            ptr = ptr.next.next

        old_list = head
        new_list = head.next
        new_head = head.next
        while old_list:
            old_list.next = old_list.next.next
            new_list.next = new_list.next.next if new_list.next else None
            old_list = old_list.next
            new_list = new_list.next
        return new_head

    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # h = []
        # dummy = ListNode(0)
        # head = dummy
        # for lst in lists:
        #     while lst:
        #         h.append(lst.val)
        #         lst = lst.next
        # heapq.heapify(h)
        # while h:
        #     curr = heapq.heappop(h)
        #     dummy.next = ListNode(curr)
        #     dummy = dummy.next
        # return head.next
        #   2nd approach
        # h = []
        # dummy = ListNode(0)
        # head = dummy
        # for lst in lists:
        #     if lst:
        #         heapq.heappush(h, lst)
        # while h:
        #     curr = heapq.heappop(h)
        #     dummy.next = curr
        #     dummy = dummy.next
        #     if curr.next:
        #         heapq.heappush(h, curr.next)
        # return head.next
        def mergeList(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
            head = ListNode(0)
            curr = head
            while l1 and l2:
                if l1.val <= l2.val:
                    curr.next = l1
                    l1 = l1.next
                else:
                    curr.next = l2
                    l2 = l2.next
                curr = curr.next
            if not l2:
                curr.next = l1
            else:
                curr.next = l2
            return head.next
        n = len(lists)
        interval = 1
        while interval < n:
            for i in range(0, n-interval, interval*2):
                lists[i] = mergeList(lists[i], lists[i+interval])
            interval *= 2
        return lists[0]













Sol = Solution()
# node5 = ListNode(5)
# node4 = ListNode(4, node5)
# node3 = ListNode(3, node4)
# node2 = ListNode(2, node3)
# node1 = ListNode(1, node2)
# n4 = Node(12, None, None)
# n3 = Node(10, n4, n4)
# n2 = Node(11, n3, n4)
# n1 = Node(13, n2, n3)
# n0 = Node(7, n1, None)
n8 = ListNode(6,None)
n7 = ListNode(2,n8)
n6 = ListNode(4,None)
n5 = ListNode(3,n6)
n4 = ListNode(1,n5)
n2 = ListNode(5,None)
n1 = ListNode(4,n2)
n0 = ListNode(1,n1)
res = Sol.mergeKLists([n0, n4, n7])
while res:
    print(res.val)
    res = res.next

# new_head = Sol.reverseList(node1)
# # while new_head is not None:
# #     print(new_head.val)
# #     new_head = new_head.next
# h = Sol.copyRandom(n0)
# while h:
#     print(h.val)
#     if h.random:
#         print(h.random.val)
#     else:
#         print(-1)
#     h = h.next
