import re
import sys
from typing import List
from collections import defaultdict


class Solution:
    def is_valid_parentheses(self, s):
        stack = []
        dt = {
            ')': '(',
            ']': '[',
            '}': '{'
        }
        for char in s:
            if char in ['(','[','{']:
                stack.append(char)
            if char in [')',']','}']:
                if len(stack) == 0:
                    return False
                last = stack.pop()
                if dt.get(char) != last:
                    return False
        if len(stack) > 0:
            return False
        return True

    def product_except_self(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0]*n
        #prefix
        prefix = 1
        for item in range(n):
            result[item] = prefix
            prefix *= nums[item]
        #suffix
        suffix = 1
        for item in range(n - 1, -1, -1):
            result[item] = suffix * result[item]
            suffix *= nums[item]

        return result

    def first_unique_char(self, s: str) -> int:
        mp = dict()
        for index in range(len(s)):
            if s[index] not in mp:
                mp[s[index]] = index
            else:
                mp[s[index]] = ""
        for v in mp.values():
            if v != "":
                return v
        return -1

    def most_common_word(self, paragraph: str, banned: List[str]) -> str:
        paragraph = re.split(r'\W+',paragraph.lower())
        mp = defaultdict(int)
        for word in paragraph:
            if word not in banned:
                if word != "":
                    mp[word] += 1
        max_count = max(mp.values())
        return [key for key, val in mp.items() if val == max_count][0]

    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        mp = {}
        result = []
        for index, number in enumerate(numbers):
            remainder = target - number
            if remainder in mp.keys():
                result = [mp.get(remainder) + 1, index + 1]
            else:
                mp[number] = index
        return result

    def reverseString(self, s: List[str]) -> None:
        def do_reverse(left, right: int) -> None:
            if left >= right:
                return
            s[left], s[right] = s[right], s[left]
            do_reverse(left + 1, right - 1)
        do_reverse(0, len(s)-1)
        print(s)


    def maxSubArray(self, nums: List[int]) -> int:
        # max_sum = nums[0]
        # array_sum = nums[0]
        # index = 1
        # while index < len(nums):
        #     array_sum = max(array_sum + nums[index], nums[index])
        #     max_sum = max(max_sum, array_sum)
        #     index += 1
        # return max_sum
        # Divide and Conquer
        #BaseCase
        if not nums:
            return 0

        def find_max(left, right: int) -> int:
            if left >= len(nums):
                return 0
            return nums[left] + find_max(left + 1, right)
        mid = len(nums) // 2
        left_max_subarray, right_max_subarray = 0, 0
        for index in range(0, mid - 1):
            left_subarray = find_max(index, mid)
            left_max_subarray = max(left_subarray, left_max_subarray)
        for index in range(mid + 1, len(nums)):
            right_subarray = find_max(index, len(nums))
            right_max_subarray = max(right_subarray, right_max_subarray)

        return max(left_max_subarray,
                   right_max_subarray,
                   left_max_subarray + nums[mid] + right_max_subarray,
                   left_max_subarray + nums[mid],
                   right_max_subarray + nums[mid])

    def romanToInt(self, s: str) -> int:

        mp = defaultdict(int, {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        })

        if s is None:
            return 0
        s = s.upper()
        result = 0
        n = len(s)
        i = 0
        while i < n:
            curr = mp[s[i]]
            if i < n-1 and mp[s[i+1]] > curr:
                result += mp[s[i+1]] - curr
                i += 1
            else:
                result += curr
            i += 1
        return result

    def trap(self, height: list[int]) -> int:
        # n = len(height)
        # ans = 0
        # mp_left = defaultdict(int)
        # mp_right = defaultdict(int)
        # left_max = 0
        # right_max = 0
        # for j in range(n-1, -1, -1):
        #     left_max = max(left_max, height[j])
        #     mp_left[j] = left_max
        # for j in range(0, n):
        #     right_max = max(right_max, height[j])
        #     mp_right[j] = right_max
        # for i in range(n):
        #     ans += min(mp_left[i], mp_right[i]) - height[i]
        # return ans
        # stack approach
        # ans = 0
        # current = 0
        # st = []
        # while current < len(height):
        #     while (len(st) != 0) and height[current] > height[st[-1]]:
        #         top = st.pop()
        #         if len(st) == 0:
        #             break
        #         distance = current - st[-1] - 1
        #         ht = min(height[st[-1]], height[current]) - height[top]
        #         ans += distance * ht
        #     st.append(current)
        #     current += 1
        # return ans
        # two pointers approach
        left = 0
        left_max = 0
        right = len(height) - 1
        right_max = 0
        ans = 0
        while left < right:
            if height[left] < height[right]:
                left_max = max(left_max, height[left])
                ans += left_max - height[left]
                left += 1
            else:
                right_max = max(right_max, height[right])
                ans += right_max - height[right]
                right -= 1
        return ans

    def threeSumClosest(self, nums: list[int], target: int) -> int:
        nums.sort()
        closest_sum = sys.maxsize
        closest_diff = sys.maxsize
        for i in range(len(nums)):
            left, right = i + 1, len(nums) - 1
            while left < right:
                nums_sum = nums[i] + nums[left] + nums[right]
                curr_diff = abs(target - nums_sum)
                if curr_diff <= closest_diff:
                    closest_sum = nums_sum
                    closest_diff = curr_diff
                if nums_sum < target:
                    left += 1
                else:
                    right -= 1
            if closest_diff == 0:
                break
        return closest_sum

    def reorderLogFiles(self, logs: list[str]) -> list[str]:
        dl_list = []
        ll_list = []
        for log in logs:
            if str.isdigit(log[-1]):
                dl_list.append(log)
            else:
                ll_list.append(log)
        ll_list.sort(key=lambda x: (x.split(' ')[1:], x.split(' ')[0]))
        return ll_list + dl_list


    def compareVersions(self, version1: str, version2: str)-> int:
        # i, j = 0, 0
        # while i < len(version1) or j < len(version2):
        #     v1, v2 = "", ""
        #     while i < len(version1) and version1[i] != '.':
        #         v1 += version1[i]
        #         i += 1
        #     while j < len(version2) and version2[j] != '.':
        #         v2 += version2[j]
        #         j += 1
        #
        #     if not v1:
        #         v1 = 0
        #     if not v2:
        #         v2 = 0
        #     v1 = int(v1)
        #     v2 = int(v2)
        #     if v1 < v2:
        #         return -1
        #     elif v1 > v2:
        #         return 1
        #     else:
        #         i += 1
        #         j += 1
        # return 0
        nums1 = version1.split(".")
        nums2 = version2.split(".")
        n1, n2 = len(nums1), len(nums2)

        # compare versions
        for i in range(max(n1, n2)):
            i1 = int(nums1[i]) if i < n1 else 0
            i2 = int(nums2[i]) if i < n2 else 0
            if i1 != i2:
                return 1 if i1 > i2 else -1

        # The versions are equal
        return 0

    def numberToWords(self, nums: int) -> str:
        if nums < 0:
            return ""
        if nums == 0:
            return "Zero"
        mp_digits = {
            "1": "One",
            "2": "Two",
            "3": "Three",
            "4": "Four",
            "5": "Five",
            "6": "Six",
            "7": "Seven",
            "8": "Eight",
            "9": "Nine",
            "10": "Ten",
            "11": "Eleven",
            "12": "Twelve",
            "13": "Thirteen",
            "14": "Fourteen",
            "15": "Fifteen",
            "16": "Sixteen",
            "17": "Seventeen",
            "18": "Eighteen",
            "19": "Nineteen",
            "20": "Twenty",
            "30": "Thirty",
            "40": "Forty",
            "50": "Fifty",
            "60": "Sixty",
            "70": "Seventy",
            "80": "Eighty",
            "90": "Ninety",
            "100": "Hundred"
        }
        mp_grp ={
            2: "Thousand",
            3: "Million",
            4: "Billion"
        }

        result = ""
        num_str = str(nums)
        grp_count = 0
        while len(str(num_str)) > 0:
            n = len(str(num_str))
            i = n - 3 if n - 3 > 0 else 0

            group = num_str[i:n]
            if len(group) > 0:
                grp_count += 1
            res = ""
            while len(group) > 0:
                grp_len = len(group)
                if int(group) == 0:
                    group = group[grp_len:]
                elif grp_len == 3:
                    if group[0] != '0':
                        if res != "":
                            res += " "
                        res += mp_digits[group[0]] + " " + mp_digits["100"]
                    group = group[1:]
                elif grp_len == 2 and group not in mp_digits:
                    if group[0] != '0':
                        if res != "":
                            res += " "
                        res += mp_digits[str(10*int(group[0]))]
                    group = group[1:]
                else:
                    if res != "":
                        res += " "
                    res += mp_digits[group]
                    group = group[grp_len:]

            if grp_count in mp_grp and res != "":
                res += " " + mp_grp[grp_count]
            if result != "" and res != "":
                res += " "
            result = res + result
            num_str = num_str[0:-3]

        return result





Sol = Solution()
#print(Sol.first_unique_char('aabb'))
#print(Sol.is_valid_parentheses("(]"))
#print(Sol.product_except_self([1, 2, 3, 4]))
#print(Sol.most_common_word("..Bob hit a ball, the hit BALL flew far after it was hit.", ["hit"]))
#print(Sol.twoSum([2,7,11,15], 9))
# print(Sol.maxSubArray([5,4,-1,7,8]))
# print(Sol.romanToInt("III"))
# print(Sol.trap([0,1,0,2,1,0,1,3,2,1,2,1]))
# print(Sol.threeSumClosest([-1,2,1,-4],1))
# Sol.reverseString(["H","a","n","n","a","h"])
# print(Sol.reorderLogFiles(["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]))
# print(Sol.compareVersions("1.01", "1.001"))
print(Sol.numberToWords(1000010))


