/*
It is foolish to climb upwards when you can go down. Can you write a program that computes the number of paths through a staircase that can be climbed upwards n steps when taking up to k steps or less at a time?
*/

package main

/*

			I love the mention of the problem solving principles, https://math.berkeley.edu/~gmelvin/polya.pdf is where I originally learned how to solve these problems.



			Lets go into the actual compute science question here
			1 or 2 steps.


			     root

			1            2

		2    3      3     4

	3 4     4 5   5 6     6 7


		 We can solve this problem first by using a brute force.

		 At each stage, we simply ask, how many ways can we get to the next step.

		 We can do this by using a recursive function.
		 We either take one step, or we take two steps. until we reach zero
*/

func descendStairs(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return -1
	} else {
		return descendStairs(n-1) + descendStairs(n-2)
	}
}

/*

While this approach works well, we can optimize our solution further, observing that we are solving the same problem many times


f(0) = 1
f(1) = 1
f(2) = 2
f(3) = f(2) + f(1) = 3
f(4) = f(3) + f(2) = 5
f(5) = f(4) + f(3) = 8


We can use a dynamic programming approach to solve this problem.
*/

func descendStairsDP(n int) int {
	arr := []int{1, 1}

	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}
