class Solution:
    def customSortString(self, order: str, s: str) -> str:
        order_map = {}
        for i, char in enumerate(order):
            order_map[char] = i

        letter_map = collections.defaultdict(int)
        for char in s:
            letter_map[char] += 1

        output = ""

        for key, value in order_map.items():
            if key in letter_map:
                output += key * letter_map[key]
                del letter_map[key]

        for key,value in letter_map.items():
            output += key * value
      
        return output
        