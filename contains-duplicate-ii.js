var containsNearbyDuplicate = function(nums, k) {
    let length = nums.length;
    let cache = {};
    for (let i = 0; i < length; i++)  {
        if (i - cache[nums[i]] <= k) {
            return true;
        } else {
            cache[nums[i]] = i;
        };
    };
    return false;
};