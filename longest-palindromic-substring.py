class Solution:
    def longestPalindrome(self, s):
        def UpdatedString(string):
            newString = ['#']
            for char in string:
                newString += [char, '#']
            return ''.join(newString)

        def Manachen(string):
            string = UpdatedString(string)
            LPS = [0 for _ in range(len(string))]
            C = 0
            R = 0

            for i in range(len(string)):
                iMirror = 2*C - i
                if R > i:
                    LPS[i] = min(R-i, LPS[iMirror])
                else:
                    LPS[i] = 0
                try:
                    while string[i + 1 + LPS[i]] == string[i - 1 - LPS[i]]:
                        LPS[i] += 1
                except:
                    pass

                if i + LPS[i] > R:
                    C = i
                    R = i + LPS[i]
            
            r, c = max(LPS), LPS.index(max(LPS))
            middle = int(c/2)
            if r % 2 == 0:
                chosen = s[middle-int(r/2):middle+int(r/2)]
            else:
                chosen = s[middle-int(r/2):middle+int(r/2)+1]
            return chosen
        return Manachen(s)