    /*

    Recursive solution
    Memoized solution
    [] bottom up solution
    two pointer solution


        f(0) = 0
        f(1) = 1
        f(2) = f(n-1), f(n-2) = 1
        f(3) = 1 + 1 = 2
        f(4) = 

        2^n
        5 
       / \ 
      4
      /\
     2  2
        /\
       1. 0
      

    */
func rec(n int, memo map[int]int) int {
    if val, ok := memo[n]; ok{ // 3
        return val
    } else {
        memo[n] = rec(n - 1, memo) + rec(n - 2, memo)
    }
    return memo[n]
}

// func fib(n int) int {
//     if n <= 1 {
//         return n
//     }
//     dp := make([]int, n+1)
//     dp[1] = 1 // 
//     dp[2] = 1 // 
//     for i := 2; i <= n; i++ {
//         // dp[5] = 3 + 2
//         dp[i] = dp[i-1] + dp[i-2]
//     }
//     return dp[n]
// }








func fib(n int) int { // n = 5
    if n <= 1 {
        return n
    }
    sumSoFar := 1
    prev := 1
    curr := 1 
    for i := 2; i < n; i++ { // 2..5
        sumSoFar = prev + curr // 1 + 1
        prev = curr 
        curr = sumSoFar
    }
    return sumSoFar
}