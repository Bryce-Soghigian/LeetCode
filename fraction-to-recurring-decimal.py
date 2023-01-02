class Solution(object):
    def fractionToDecimal(self, numerator, denominator):
        """
        :type numerator: int
        :type denominator: int
        :rtype: str
        """
        # generate integer part of solution
        sign = 1
        if numerator < 0:
            sign = -1
            numerator = -numerator
        if denominator < 0:
            sign *= -1
            denominator = -denominator
        # use s to store result as list (rather than string) for efficienty
        s = []
        if sign < 0 and numerator:
            s.append('-')
        s.append(str(numerator // denominator))
        numerator = numerator % denominator
        if numerator:
            s.append('.')
        
        m = {}
        while numerator:
            i = m.get(numerator, -1)
            if i < 0:
                i = len(s)
                m[numerator] = i
                numerator *= 10
                s += str(numerator // denominator)
                numerator = numerator % denominator
            else:
                s.insert(i, '(')
                s.append(')')
                break
        return ''.join(s)