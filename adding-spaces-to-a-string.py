class Solution:
    def addSpaces(self, s: str, spaces: List[int]) -> str:
        output = ""
        pt = 0
        for i, val in enumerate(s):
            if pt <= len(spaces)-1 and i == spaces[pt]:
                pt += 1
                output += " "
            output += val
        return output

        