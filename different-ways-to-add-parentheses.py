
class Solution:
    @cache
    def evaluate(self, x,y, operation):
        if operation == "+":
            return x + y
        if operation == "*":
            return x * y
        if operation == "-":
            return x - y
        
        return 0
    @cache
    def diffWaysToCompute(self, expression: str) -> List[int]:

        results = []
        num = True
        for i, val in enumerate(expression):
            if not val.isnumeric():
                # if we are not at a number then we need to evaluate our expression tree
                num = False
                left_operands = self.diffWaysToCompute(expression[0:i])
                right_operands = self.diffWaysToCompute(expression[i+1:])
                for dx in left_operands:
                    for dy in right_operands:
                        new_num = self.evaluate(dx,dy, val)
                        results.append(new_num)
        if num:
            results.append(int(expression))
        
        return results
