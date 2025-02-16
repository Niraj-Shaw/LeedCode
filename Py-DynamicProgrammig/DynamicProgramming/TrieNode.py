from collections import defaultdict


class TrieNode:
    def __init__(self):
        self.nodes = defaultdict[str](None)
        self.isEnd = False
