class Solution:
	def isEvenOddTree(self, root: TreeNode) -> bool:
		queue, result, level = deque([root]), [], 1

		if root.val % 2 != 1 : return False

		while queue:
			cur_arr = []
			for _ in range(len(queue)):
				curNode = queue.popleft()
				cur_arr.append(curNode.val)

				if level % 2 == 0:
					if curNode.left:
						if curNode.left.val % 2 == 1:
							queue.append(curNode.left)
						else:
							return False
					if curNode.right:
						if curNode.right.val % 2 == 1:
							queue.append(curNode.right)
						else:
							return False
				else:
					if curNode.left:
						if curNode.left.val % 2 == 0:
							queue.append(curNode.left)
						else:
							return False
					if curNode.right:
						if curNode.right.val % 2 == 0:
							queue.append(curNode.right)
						else:
							return False
			level += 1

			if len(set(cur_arr)) == len(cur_arr): # check for duplicates
				if level % 2 == 0 and sorted(cur_arr) != cur_arr: return False

				if level % 2 == 1 and sorted(cur_arr, reverse=True) != cur_arr: return False

			else: return False # if there are duplicates, there cannot be an increasing or decreasing order

		return True