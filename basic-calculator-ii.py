class Solution(object):
    def calculate(self, s):
        def evaluate(stack):  # evaluate the last 3 elements of stack and modify in-place
            b = stack.pop()
            op = stack.pop()
            a = stack.pop()
            if op == '+': stack.append(a + b)
            elif op == '-': stack.append(a - b)
            elif op == '*': stack.append(a * b)
            elif op == '/': stack.append(a // b)

        stack = []
        for c in s:
            if c.isdigit():
                c = int(c)
                if stack and type(stack[-1]) == int:  # top of stack is digit, because the previous character is digit
                    stack[-1] = stack[-1] * 10 + c
                else:
                    stack.append(c)
            elif c in '+-*/':
                # The current stack top expression is '*' or '/', evaluate it immediately.
                # But if the the current stack top expression is '+' or '-', it needs to wait for the next '+' or '-'
                # or the end of s.
                if len(stack) > 1 and stack[-2] in '*/':
                    evaluate(stack)
                # Because '*' or '/' evaluated immediately, at most one expression is left on the top of the current
                # stack, and it can only be '+' or '-'. If the current character is '+' or '-', we should evaluate the
                # previous '+' or '-'.
                if len(stack) > 1 and c in '+-':
                    evaluate(stack)
                stack.append(c)

        while len(stack) > 1:  # At most 2 expressions are left, the first is '+' or '-' and the second is '*' or '/'
            evaluate(stack)
        return stack[-1]