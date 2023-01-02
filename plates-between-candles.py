class Solution:
	def platesBetweenCandles(self, s: str, queries: List[List[int]]) -> List[int]:
		candles = []
		res = []
		for i, c in enumerate(s):
			if c == "|":
				candles.append((i, len(candles) + 1))
		for x, y in queries:
			l1 = 0
			r1 = len(candles) - 1
			while l1 <= r1:
				m1 = (l1 + r1) // 2
				if candles[m1][0] == x:
					l1 = m1
					break
				elif candles[m1][0] < x:
					l1 = m1 + 1
				else:
					r1 = m1 - 1
			l2 = 0
			r2 = len(candles) - 1
			while l2 <= r2:
				m2 = (l2 + r2) // 2
				if candles[m2][0] == y:
					r2 = m2
					break
				elif candles[m2][0] > y:
					r2 = m2 - 1              
				else:
					l2 = m2 + 1
			res.append(max(candles[r2][0] - candles[l1][0] - (candles[r2][1] - candles[l1][1]), 0))
		return res
        
        