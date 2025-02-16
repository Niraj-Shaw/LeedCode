from collections import defaultdict

from Trie import Trie
from TrieNode import TrieNode


class Solution:
    def longest_palindrome(self, s: str) -> str:
        result = [0,0]
        n = len(s)
        dp = [[False]*n for j in range(n)]
        for i in range(n):
            dp[i][i] = True
            result = [i, i+1]
        for i in range(n-1):
            if s[i] == s[i+1]:
                dp[i][i+1] = True
                result = [i, i+2]
        for i in range(n-3, -1, -1):
            for j in range(i+2, n):
                if s[i] == s[j] and dp[i+1][j-1]:
                    dp[i][j] = True
                    if j - i > result[1] - result[0] - 1:
                        result = [i, j + 1]

        return s[result[0]:result[1]]
        # n = len(s)
        # dp = [[False] * n for _ in range(n)]
        # ans = [0, 0]
        #
        # for i in range(n):
        #     dp[i][i] = True
        #
        # for i in range(n - 1):
        #     if s[i] == s[i + 1]:
        #         dp[i][i + 1] = True
        #         ans = [i, i + 1]
        #
        # for diff in range(2, n):
        #     for i in range(n - diff):
        #         j = i + diff
        #         if s[i] == s[j] and dp[i + 1][j - 1]:
        #             dp[i][j] = True
        #             ans = [i, j]
        #
        # i, j = ans
        # return s[i: j + 1]

    def maxProfit(self, prices: list[int]) -> int:
        max_profit = 0
        buying_price = prices[0]
        for index in range(1, len(prices)):
            selling_price = prices[index]
            curr_profit = selling_price - buying_price
            if curr_profit < 0:
                buying_price = selling_price
            elif curr_profit > max_profit:
                max_profit = curr_profit
        return max_profit

    def wordBreak(self, s: str, wordDict : list[str]) -> bool:
        # t = Trie()
        # memo = defaultdict(bool)
        # for word in wordDict:
        #     t.insert(word)

        # def backtrack(i: int) -> bool:
        #     if i in memo:
        #         return memo[i]
        #     if i == len(s):
        #         return True
        #
        #     node = t.root
        #     for j in range(i, len(s)):
        #         ch = s[j]
        #         if ch not in node.nodes:
        #             break
        #         node = node.nodes[ch]
        #         if node.isEnd:
        #             if backtrack(j + 1):
        #                 memo[i] = True
        #                 return True
        #
        #     memo[i] = False
        #     return False
        # return backtrack(0)

        # dp = [False] * len(s)
        # for i in range(len(s)):
        #     if i == 0 or dp[i - 1]:
        #         curr = t.root
        #         for j in range(i, len(s)):
        #             c = s[j]
        #             if c not in curr.nodes:
        #                 # No words exist
        #                 break
        #
        #             curr = curr.nodes[c]
        #             if curr.isEnd:
        #                 dp[j] = True
        #
        # return dp[-1]

        # bottom up approach

        # dp = [False] * len(s)
        # for i in range(len(s)):
        #     for word in wordDict:
        #         if i < len(word) - 1:
        #             continue
        #         if i == len(word) - 1 or dp[i - len(word)]:
        #             if s[i - len(word) + 1: i + 1] == word:
        #                 dp[i] = True
        #                 break
        # return dp[-1]

        # top down approach
        memo = defaultdict(bool)
        def doWork(i: int) -> bool:
            if i in memo:
                return memo[i]
            for word in wordDict:
                if i > len(s) - 1:
                    return True
                if s[i: i+len(word)] == word and doWork(i + len(word)):
                    memo[i] = True
                    return True
            memo[i] = False
            return False

        return doWork(0)




















Sol = Solution()
# print(Sol.longest_palindrome("aaaa"))
# print(Sol.maxProfit([7,6,4,3,1]))
print(Sol.wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]))
print(Sol.wordBreak("aaaaaaa", ["aaaa", "aaa"]))
print(Sol.wordBreak("abaaaab", ["a", "aa"]))
