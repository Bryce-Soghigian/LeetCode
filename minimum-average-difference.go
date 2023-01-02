func minimumAverageDifference(nums []int) int {
    idx, ans := -1, 100001
    l, n, sum := 0, len(nums), 0
    for _, num := range nums { sum += num }
    for i := 0; i < n; i++ {
        l += nums[i]
        r := sum - l
        diff := 100001
        if i == n - 1 {
            diff = abs(l / (i + 1))
        } else {
            diff = abs((l / (i + 1)) - (r / (n - i - 1)))
        }
        if diff < ans {
            ans = diff
            idx = i
        }
    }
    return idx
}

func abs(n int) int {
    if n > 0 { return n }
    return -n
}