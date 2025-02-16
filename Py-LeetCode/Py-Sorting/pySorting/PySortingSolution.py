from collections import defaultdict, Counter
from typing import List
import heapq
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        # if not intervals:
        #     return []
        # intervals.sort(key=lambda s: s[0])
        # start = intervals[0][0]
        # end = intervals[0][1]
        # res = []
        # for row in range(1, len(intervals)):
        #     curr_start = intervals[row][0]
        #     curr_end = intervals[row][1]
        #     if curr_start > end:
        #         res.append([start, end])
        #         start = curr_start
        #     else:
        #         start = min(start, curr_start)
        #     if (curr_end > curr_start) and (curr_end > end):
        #         end = curr_end
        # res.append([start, end])
        # return res
        intervals.sort(key=lambda x: x[0])

        merged = []
        for interval in intervals:
            # if the list of merged intervals is empty or if the current
            # interval does not overlap with the previous, simply append it.
            if not merged or merged[-1][1] < interval[0]:
                merged.append(interval)
            else:
                # otherwise, there is overlap, so we merge the current and previous
                # intervals.
                merged[-1][1] = max(merged[-1][1], interval[1])

        return merged

    def findKthLargest(self, nums: list[int], k: int) -> int:
        # result_lst = heapq.nlargest(k, nums)
        # return result_lst[k-1]
        # Quickselect- # Partitioning
        self.find_pivot(nums, 0, len(nums) - 1, k-1)
        return nums[k-1]

    def find_pivot(self, nums: list[int], left, right, k):
        if right - left <= 0:
            return
        pivot_index = self.partition(nums, left, right)
        if k == pivot_index:
            return nums[pivot_index]
        elif k < pivot_index:
            self.find_pivot(nums, left, pivot_index - 1, k)
        elif k > pivot_index:
            self.find_pivot(nums, pivot_index + 1, right, k)

    def partition(self, nums: list[int], left, right: int) -> int:
        pivot_index = left
        pivot = nums[left]
        left += 1
        while left <= right:
            if (nums[left] < pivot) and (nums[right] > pivot):
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
                right -= 1
            if nums[left] >= pivot:
                left += 1
            if nums[right] <= pivot:
                right -= 1
        nums[right], nums[pivot_index] = nums[pivot_index], nums[right]
        return right

    def kClosest(self, points: list[list[int]], k: int) -> list[list[int]]:
        # heapq.heapify(points)
        # distances = []
        # res = []
        # for i in range(len(points)):
        #     point = points[i]
        #     x, y = point[0], point[1]
        #     distance = x**2 + y**2
        #     heapq.heappush(distances, (distance, i))
        # while k > 0:
        #     item = heapq.heappop(distances)
        #     res.append(points[item[1]])
        #     k -= 1
        #
        # return res
        # def findSqrt(point: list[int]) -> int:
        #     return point[0]**2 + point[1]**2
        # heap = [(-findSqrt(points[i]), i) for i in range(k)]
        # heapq.heapify(heap)
        # for i in range(k, len(points)):
        #     item = -findSqrt(points[i])
        #     if item > heap[0][0]:
        #         heapq.heappushpop(heap, (item, i))
        # return [points[i] for (_, i) in heap]

        def findSqrt(point: list[int]) -> int:
            return point[0]**2 + point[1]**2

        def getSortedPivot(left, right: int) -> int:
            # p_index = left + (right - left)//2
            p_index = right
            pivot = points[p_index]
            p_dist = findSqrt(pivot)
            right -= 1

            while True:
                l_left = findSqrt(points[left])
                r_right = findSqrt(points[right])
                while l_left < p_dist:
                    left += 1
                    l_left = findSqrt(points[left])
                while r_right > p_dist:
                    right -= 1
                    r_right = findSqrt(points[right])
                if left >= right:
                    break
                else:
                    points[right], points[left] = points[left], points[right]
                    left += 1

            points[p_index], points[left] = points[left], points[p_index]
            return left

        pivot_index = len(points)
        l, r = 0, len(points) - 1
        while pivot_index != k:
            pivot_index = getSortedPivot(l, r)
            if pivot_index < k:
                l = pivot_index + 1
            elif pivot_index > k:
                r = pivot_index - 1
        return points[:k]

    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        # res = []
        # h = [(0, 0)]*k
        # mp = defaultdict(int)
        # for num in nums:
        #     mp[num] += 1
        # for item in mp.items():
        #     if item[1] > h[0][0]:
        #         heapq.heappop(h)
        #         heapq.heappush(h, (item[1], item[0]))
        # for item in reversed(h):
        #     res.append(item[1])
        # return res
        if len(nums) == k:
            return nums
        count = Counter(nums)*
        return heapq.nlargest(k, count.keys(), key=count.get)









Sol = Solution()
# print(Sol.findKthLargest(nums=[2,1], k=1))
# print(Sol.merge([[1,3],[2,6],[6,6],[6,18]]))
# print(Sol.kClosest([[3,3],[5,-1],[-2,4]], 2))
print(Sol.topKFrequent([1,1,1,2,2,3], 2))



