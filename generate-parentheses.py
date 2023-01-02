class Solution:



    def generateParenthesis(self, n: int) -> List[str]:

            result = []

            def helper(s, n_o, n_c, n):

                if n_o == n and n_c == n:
                    result.append(s)

                if n_o < n:
                    helper(s + "(", n_o + 1, n_c, n)

                if n_c < n_o:
                    helper(s+")", n_o, n_c + 1, n)

            helper("", 0, 0, n)

            return result