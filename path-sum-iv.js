/**
 * @param {number[]} nums
 * @return {number}
 */
var pathSum = function(nums) {
    let sum = 0;
  const tree = new Map();
  if (nums == null || nums.length == 0) return 0;

  for (let num of nums) {
    let key = Math.floor(num / 10);
    let value = num % 10;
    tree.set(key, value);
  }

  function traverse(root, preSum) {
    let level = Math.floor(root / 10);
    let pos = root % 10;
    let left = (level + 1) * 10 + pos * 2 - 1;
    let right = (level + 1) * 10 + pos * 2;

    let curSum = preSum + tree.get(root);

    if (!tree.has(left) && !tree.has(right)) {
      sum += curSum;
      return;
    }

    if (tree.has(left)) traverse(left, curSum);
    if (tree.has(right)) traverse(right, curSum);
  }

  traverse(Math.floor(nums[0] / 10), 0);

  return sum;  
};