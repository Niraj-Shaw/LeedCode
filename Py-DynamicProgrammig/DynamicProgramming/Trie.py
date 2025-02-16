from collections import defaultdict

from TrieNode import TrieNode


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str):
        cur_node = self.root
        for ch in word:
            if ch not in cur_node.nodes:
                cur_node.nodes[ch] = TrieNode()
            cur_node = cur_node.nodes[ch]
        cur_node.isEnd = True

    def search(self, word: str) -> bool:
        cur_node = self.root
        for ch in word:
            if not cur_node.nodes[ch]:
                return False
            cur_node = cur_node.nodes[ch]
        if "*" in cur_node.nodes:
            return True

        return False

