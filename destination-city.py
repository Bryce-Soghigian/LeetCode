class Solution:
	def destCity(self, paths: List[List[str]]) -> str:
		"""
		Runtime: 44 ms, faster than 98.02% of Python3 online submissions for Destination City.
		Memory Usage: 14.4 MB, less than 15.26% of Python3 online submissions for Destination City.
		"""
		myDict=defaultdict(str)

		for start,end in paths:
			myDict[start]=end

		while True:
			if myDict[end]:
				end=myDict[end]
			else:
				return end