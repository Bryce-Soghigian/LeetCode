func dailyTemperatures(temperatures []int) []int {
    // Create a stack to store the indices of the temperatures
    stack := make([]int, 0)

    // Create a slice to store the result
    result := make([]int, len(temperatures))

    // Iterate through the temperatures in reverse order
    for i := len(temperatures)-1; i >= 0; i-- {
        // Pop the top element from the stack while the top element is less than the current temperature
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
            stack = stack[:len(stack)-1]
        }

        // If the stack is non-empty, set the result for the current temperature to the difference between the current index and the top element of the stack
        if len(stack) > 0 {
            result[i] = stack[len(stack)-1] - i
        }

        // Push the current index onto the stack
        stack = append(stack, i)
    }

    return result
}
