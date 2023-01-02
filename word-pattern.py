class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        iso_map = {}
        values = set()
        patterns = s.split(" ")
        if len(pattern) != len(patterns):
            return False
        pt = 0
        for letter in pattern:
            if letter not in iso_map:
                if patterns[pt] in values:
                    return False
                iso_map[letter] = patterns[pt]
                values.add(patterns[pt])
            else:
                if patterns[pt] != iso_map[letter]:
                    return False
            pt += 1
        return True