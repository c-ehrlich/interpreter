# interpreter

- ours tree-walks. this is the simplest and most portable, but also the slowest
- fast ones usually JIT some kind of bytecode with a VM
- similar to ours: LISP interpreter from "The Structure and Interpretation of Computer Programs"
- we need two things
  - 1. tree-walking evaluator
    - basically just an `eval` function that traverses left/right
    - possible pseudocode
    ```js
    function eval(astNode) {
        if (astNode is integerliteral) {
            return astNode.integerValue
        } else if (astNode is booleanLiteral) {
            return astNode.booleanValue
        } else if (astNode is infixExpression) {
            leftEvaluated = eval(astNode.Left)
            rightEvaluated = eval(astNode.Right)
            if astNode.Operator == "+" {
                return leftEvaluated + rightEvaluated
            } else if ast.Operator == "-" {
                return leftEvaluated - rightEvaluated
            }
        }
    }
    ```
  - 2. a way to represent monkey values in go (during treewalking, but also they need to stick around, ie)
  ```js
  let a = 1;
  a = 2;
  ```
  in Monkey `a` is an IntergerLiteral, but how do we store it in Go?
  some interpreters use native types, some use pointers, some mix
  ref: https://github.com/wren-lang/wren
  also how is stuff represented? primitives? pointers? wrapped in objects?
- we will wrap every value in an object, which is slow and expensive, but easier