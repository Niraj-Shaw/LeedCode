"""
class Solution(object):
    def missingNumber(self, nums):

        :type nums: List[int]
        :rtype: int
        """


def missing_number(nums):
    nums.sort()
    for num in range(len(nums)):
        curr = nums[num]
        if num != curr:
            return num
    return len(nums)


nums = [0,2,3]
result = missing_number(nums)
print(result)

