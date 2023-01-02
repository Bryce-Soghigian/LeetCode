class Solution:
    def findNumOfValidWords(self, words: List[str], puzzles: List[str]) -> List[int]:
        f = lambda x: ''.join(sorted(list(set(x))))
        c = defaultdict(int)
        for i in words:
            x = f(i)
            for y in x:
                c[(y, x)] += 1
        res = []
        for i in puzzles:
            x = f(i)
            res.append(sum(c[(i[0], ''.join(s))] for s in chain.from_iterable(combinations(x, i) for i in range(len(x)+1))))
        return res