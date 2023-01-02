

func isValid(s string) bool {
    stack := make([]rune, 0)
    pairs := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
    }
    for _, char := range s {
        switch char {
            case '(', '{', '[':
            stack = append(stack, char)
            case ')', '}', ']':
            if len(stack) == 0 || stack[len(stack) - 1] != pairs[char] {
                return false
            }
            stack = stack[:len(stack)-1] // wtf why no pop go suck for algos frfr but i guess thats not its usecase :()
        }
    }
    return len(stack) == 0

}

