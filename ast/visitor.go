package ast

import (
	"fmt"
	"my_language/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Callable interface {
	NrArgs() int
	Call(v *Visitor, args []any) any
}

type Visitor struct {
	parser.BaseGrammarVisitor
	currEnv *Environment
}

type ReturnValueSignal struct {
	Value any
}

type RuntimeFunction struct {
	Parameters []string
	Body       *parser.BlockStmtContext
	Env        *Environment
}

func NewVisitor() *Visitor {
	return &Visitor{
		currEnv: NewGlobalEnvironment(),
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *Visitor) GetEnvironment() *Environment { return v.currEnv }

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	//loop over statements and function declarations
	for _, child := range ctx.GetChildren() {
		if parseTree, ok := child.(antlr.ParseTree); ok {
			parseTree.Accept(v)
		}
	}
	return nil
}

func (v *Visitor) VisitExprStmt(ctx *parser.ExprStmtContext) any {
	return ctx.Expr().Accept(v)
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	} else if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	} else if ctx.ArrayAssignStmt() != nil {
		return ctx.ArrayAssignStmt().Accept(v)
	} else if ctx.CompoundAssignStmt() != nil {
		return ctx.CompoundAssignStmt().Accept(v)
	} else if ctx.ExprStmt() != nil {
		return ctx.ExprStmt().Accept(v)
	} else if ctx.PrintStmt() != nil {
		return ctx.PrintStmt().Accept(v)
	} else if ctx.BlockStmt() != nil {
		return ctx.BlockStmt().Accept(v)
	} else if ctx.IfStmt() != nil {
		return ctx.IfStmt().Accept(v)
	} else if ctx.WhileStmt() != nil {
		return ctx.WhileStmt().Accept(v)
	} else if ctx.ForStmt() != nil {
		return ctx.ForStmt().Accept(v)
	} else if ctx.PostfixStmt() != nil {
		return ctx.PostfixStmt().Accept(v)
	} else if ctx.ReturnStmt() != nil {
		return ctx.ReturnStmt().Accept(v)
	}

	return nil
}

func (v *Visitor) VisitForInit(ctx *parser.ForInitContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	} else if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitForPost(ctx *parser.ForPostContext) any {
	if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	} else if ctx.PostfixStmt() != nil {
		return ctx.PostfixStmt().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitVarDecl(ctx *parser.VarDeclContext) any {
	varName := ctx.IDENTIFIER().GetText()
	value := ctx.Expr().Accept(v)
	// v.vars[varName] = value
	v.currEnv.Define(varName, value)
	return nil
}

func (v *Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) any {
	varName := ctx.IDENTIFIER().GetText()
	value := ctx.Expr().Accept(v)
	// v.vars[varName] = value
	if !v.currEnv.Assign(varName, value) {
		panic("Undefined variable: " + varName)
	}
	return nil
}

func (v *Visitor) VisitPostfixStmt(ctx *parser.PostfixStmtContext) any {
	varName := ctx.IDENTIFIER().GetText()
	op := ctx.GetOp().GetText()

	value, exists := v.currEnv.Lookup(varName)
	if !exists {
		panic("Undefined variable: " + varName)
	}

	intValue, ok := value.(int)
	if !ok {
		panic("Variable is not an integer: " + varName)
	}

	switch op {
	case "++":
		intValue++
	case "--":
		intValue--
	default:
		panic("Unknown postfix operator: " + op)
	}

	if !v.currEnv.Assign(varName, intValue) {
		panic("Failed to assign value to variable: " + varName)
	}

	return nil
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	// convert into integer
	val, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return val
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	varName := ctx.IDENTIFIER().GetText()
	// return v.vars[varName]
	val, exists := v.currEnv.Lookup(varName)
	if !exists {
		panic("Undefined variable: " + varName)
	}
	return val
}

func (v *Visitor) VisitString(ctx *parser.StringContext) any {
	str := ctx.STRING().GetText()
	// remove the surrounding quotes
	return str[1 : len(str)-1]
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	_, leftIsString := left.(string)
	_, rightIsString := right.(string)

	if leftIsString || rightIsString {
		if op == "+" {
			return cleanStringRepr(left) + cleanStringRepr(right)
		} else {
			panic("Unsupported operator for strings: " + op)
		}
	}

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for operator: " + op)
	}

	switch op {
	case "+":
		return lVal + rVal
	case "-":
		return lVal - rVal
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitMulDivMod(ctx *parser.MulDivModContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	if left == nil || right == nil {
		panic("Undefined variable used in expression")
	}

	switch op {
	case "*":
		return left.(int) * right.(int)
	case "/":
		if right.(int) == 0 {
			panic("Division by zero")
		}
		return left.(int) / right.(int)
	case "%":
		if right.(int) == 0 {
			panic("Modulo by zero")
		}
		return left.(int) % right.(int)
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitExponential(ctx *parser.ExponentialContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for operator: **")
	}

	return power(lVal, rVal)
}

func (v *Visitor) VisitBoolean(ctx *parser.BooleanContext) any {
	val := ctx.GetVal().GetText()
	if val == "true" {
		return true
	}
	return false
}

func (v *Visitor) VisitComparison(ctx *parser.ComparisonContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case "<":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal < rVal
		}
		panic("Comparison operator '<' requires integer operands")
	case ">":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal > rVal
		}
		panic("Comparison operator '>' requires integer operands")
	case "<=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal <= rVal
		}
		panic("Comparison operator '<=' requires integer operands")
	case ">=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal >= rVal
		}
		panic("Comparison operator '>=' requires integer operands")
	default:
		panic("Unknown operator: " + op)
	}
}

func (v *Visitor) VisitPrintStmt(ctx *parser.PrintStmtContext) any {
	val := ctx.Expr().Accept(v)

	fmt.Println(cleanStringRepr(val))

	return nil
}

func (v *Visitor) VisitBlockStmt(ctx *parser.BlockStmtContext) any {
	for _, child := range ctx.GetChildren() {
		if stmt, ok := child.(*parser.StatementContext); ok {
			if stmt != nil {
				stmt.Accept(v)
			}
		} else if fn, ok := child.(*parser.FuncStmtContext); ok {
			if fn != nil {
				fn.Accept(v)
			}
		}
	}
	return nil
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) any {
	condition := ctx.Expr().Accept(v)
	conditionBool, ok := condition.(bool)
	if !ok {
		panic("Condition in if statement must be boolean")
	}

	if conditionBool {
		ctx.GetThenBranch().Accept(v)
	} else if ctx.GetElseBranch() != nil {
		ctx.GetElseBranch().Accept(v)
	}

	return nil
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) any {
	for {
		condition := ctx.Expr().Accept(v)
		conditionBool, ok := condition.(bool)
		if !ok {
			panic("Condition in while statement must be boolean")
		}

		if !conditionBool {
			break
		}
		ctx.GetBody().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitForStmt(ctx *parser.ForStmtContext) any {
	if ctx.GetInit() != nil {
		ctx.GetInit().Accept(v)
	}

	for {
		if ctx.GetCond() != nil {
			condition := ctx.GetCond().Accept(v)
			conditionBool, ok := condition.(bool)
			if !ok {
				panic("For loop condition must evaluate to a boolean expression")
			}
			if !conditionBool {
				break
			}
		}
		if ctx.GetBody() != nil {
			ctx.GetBody().Accept(v)
		}

		if ctx.GetPost() != nil {
			ctx.GetPost().Accept(v)
		}
	}

	return nil
}

func (v *Visitor) VisitCompoundAssignStmt(ctx *parser.CompoundAssignStmtContext) any {
	varName := ctx.IDENTIFIER().GetText()
	op := ctx.GetOp().GetText()
	value := ctx.Expr().Accept(v)

	currentValue, exists := v.currEnv.Lookup(varName)
	if !exists {
		panic("Undefined variable: " + varName)
	}

	// handle string concatenation for +=
	if op == "+=" {
		if strVal, ok := currentValue.(string); ok {
			if !v.currEnv.Assign(varName, strVal+cleanStringRepr(value)) {
				panic("Failed to assign value to variable: " + varName)
			}

			return nil
		}
	}

	intCurrentValue, ok := currentValue.(int)
	if !ok {
		panic("Variable is not an integer: " + varName)
	}

	intValue, ok := value.(int)
	if !ok {
		panic("Right-hand side of compound assignment must be an integer")
	}

	var result int
	switch op {
	case "+=":
		result = intCurrentValue + intValue
	case "-=":
		result = intCurrentValue - intValue
	case "*=":
		result = intCurrentValue * intValue
	case "/=":
		if intValue == 0 {
			panic("Division by zero in compound assignment")
		}
		result = intCurrentValue / intValue
	case "%=":
		if intValue == 0 {
			panic("Modulo by zero in compound assignment")
		}
		result = intCurrentValue % intValue
	case "**=":
		result = power(intCurrentValue, intValue)
	case "&=":
		result = intCurrentValue & intValue
	case "|=":
		result = intCurrentValue | intValue
	case "^=":
		result = intCurrentValue ^ intValue
	case "<<=":
		result = intCurrentValue << intValue
	case ">>=":
		result = intCurrentValue >> intValue
	default:
		panic("Unknown compound assignment operator: " + op)
	}
	if !v.currEnv.Assign(varName, result) {
		panic("Failed to assign value to variable: " + varName)
	}
	return nil
}

func (v *Visitor) VisitParentheses(ctx *parser.ParenthesesContext) any {
	return ctx.Expr().Accept(v)
}

func (v *Visitor) VisitUnary(ctx *parser.UnaryContext) any {
	op := ctx.GetOp().GetText()
	value := ctx.Expr().Accept(v)

	switch op {
	case "not":
		boolVal, ok := value.(bool)
		if !ok {
			panic("Operand of 'not' must be boolean")
		}
		return !boolVal
	case "~":
		intVal, ok := value.(int)
		if !ok {
			panic("Operand of '~' must be an integer")
		}
		return ^intVal
	case "-":
		intVal, ok := value.(int)
		if !ok {
			panic("TypeError: Unary minus operator can only be applied to an integer")
		}
		return -intVal
	default:
		panic("Unknown unary operator: " + op)
	}
}

func (v *Visitor) VisitBitShift(ctx *parser.BitShiftContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for bit shift operators")
	}

	switch op {
	case "<<":
		return lVal << rVal
	case ">>":
		return lVal >> rVal
	default:
		panic("Unknown bit shift operator: " + op)
	}
}

func (v *Visitor) VisitBitAnd(ctx *parser.BitAndContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for bitwise AND operator")
	}

	return lVal & rVal
}

func (v *Visitor) VisitBitXor(ctx *parser.BitXorContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for bitwise XOR operator")
	}

	return lVal ^ rVal
}

func (v *Visitor) VisitBitOr(ctx *parser.BitOrContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic("Both operands must be integers for bitwise OR operator")
	}

	return lVal | rVal
}

func (v *Visitor) VisitAnd(ctx *parser.AndContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	leftBool, leftOk := left.(bool)
	rightBool, rightOk := right.(bool)
	if !leftOk || !rightOk {
		panic("Operands of 'and' must be boolean")
	}
	return leftBool && rightBool
}

func (v *Visitor) VisitOr(ctx *parser.OrContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	leftBool, leftOk := left.(bool)
	rightBool, rightOk := right.(bool)
	if !leftOk || !rightOk {
		panic("Operands of 'or' must be boolean")
	}
	return leftBool || rightBool
}

func (v *Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	var elements []any
	for _, expr := range ctx.AllExpr() {
		elements = append(elements, expr.Accept(v))
	}
	// return a pointer to the slice
	return &elements
}

func (v *Visitor) VisitArrayIndex(ctx *parser.ArrayIndexContext) any {
	array := ctx.Expr(0).Accept(v)
	index := ctx.Expr(1).Accept(v)

	// expect pointer to the slice (*[]any)
	arrPtr, ok := array.(*[]any)
	if !ok {
		panic("Attempting to index a non-array value")
	}

	// dereference it locally to perform size checks and lookups safely
	arr := *arrPtr

	idx, ok := index.(int)
	if !ok {
		panic("Array index must be an integer")
	}

	if idx < 0 || idx >= len(arr) {
		panic("Array index out of bounds")
	}

	return arr[idx]
}

func (v *Visitor) VisitArrayAssignStmt(ctx *parser.ArrayAssignStmtContext) any {
	arrayName := ctx.IDENTIFIER().GetText()

	val, exists := v.currEnv.Lookup(arrayName)
	if !exists {
		panic("Undefined variable: " + arrayName)
	}

	arrPtr, ok := val.(*[]any)
	if !ok {
		panic("Attempting to index a non-array variable: " + arrayName) // 👈 This is where your panic came from!
	}

	indexVal := ctx.Expr(0).Accept(v)
	newVal := ctx.Expr(1).Accept(v)

	idx, ok := indexVal.(int)
	if !ok {
		panic("Array index must be an integer")
	}

	arr := *arrPtr
	if idx < 0 || idx >= len(arr) {
		panic("Array index out of bounds")
	}

	(*arrPtr)[idx] = newVal

	return newVal
}

// func (v *Visitor) VisitArrayAssignStmt(ctx *parser.ArrayAssignStmtContext) any {
// 	varName := ctx.IDENTIFIER().GetText()
// 	index := ctx.Expr(0).Accept(v)
// 	value := ctx.Expr(1).Accept(v)

// 	target, ok := v.currEnv.Lookup(varName)
// 	if !ok {
// 		panic("Undefined variable: " + varName)
// 	}

// 	arr, ok := target.([]any)
// 	if !ok {
// 		panic("Attempting to index a non-array variable: " + varName)
// 	}

// 	idx, ok := index.(int)
// 	if !ok {
// 		panic("Array index must be an integer")
// 	}

// 	if idx < 0 || idx >= len(arr) {
// 		panic("Array index out of bounds")
// 	}

// 	arr[idx] = value
// 	return nil
// }

func (v *Visitor) VisitFuncStmt(ctx *parser.FuncStmtContext) any {
	funcName := ctx.IDENTIFIER(0).GetText()

	var params []string
	i := 1
	for {
		paramNode := ctx.IDENTIFIER(i)
		if paramNode == nil {
			break
		}
		params = append(params, paramNode.GetText())
		i++
	}

	v.currEnv.Define(funcName, &RuntimeFunction{
		Parameters: params,
		Body:       ctx.BlockStmt().(*parser.BlockStmtContext),
		Env:        v.currEnv,
	})

	return nil
}

func (f *RuntimeFunction) Call(v *Visitor, args []any) any {
	//link back to where the function was defined
	funcEnv := NewScope(f.Env)

	// bind evaluated arguments to parameter names
	for i, paramName := range f.Parameters {
		funcEnv.Define(paramName, args[i])
	}

	// if no visitor provided, create a temporary one
	if v == nil {
		v = &Visitor{currEnv: funcEnv}
	}

	// swap the active environment
	previousEnv := v.currEnv
	v.currEnv = funcEnv

	var returnedVal any = nil
	func() {
		// safe return value recovery handler
		defer func() {
			if r := recover(); r != nil {
				if returnSignal, ok := r.(ReturnValueSignal); ok {
					returnedVal = returnSignal.Value
				} else {
					panic(r) // panic if it's a bug/crash
				}
			}
		}()

		f.Body.Accept(v)
	}()

	// safely restore parent scope state
	v.currEnv = previousEnv
	return returnedVal
}

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	funcName := ctx.IDENTIFIER().GetText()

	// pull the entity out of the environment stack
	fn, exists := v.currEnv.Lookup(funcName)
	if !exists {
		panic("Undefined function: " + funcName)
	}

	callable, ok := fn.(Callable)
	if !ok {
		panic(funcName + " is not a callable function")
	}

	// evaluate the call arguments at the caller's execution environment line
	var argValues []any
	for _, argExpr := range ctx.AllExpr() {
		argValues = append(argValues, argExpr.Accept(v))
	}

	// validate input count parameters
	if callable.NrArgs() != -1 && len(argValues) != callable.NrArgs() {
		panic(fmt.Sprintf("Function %s expects %d arguments, got %d", funcName, callable.NrArgs(), len(argValues)))
	}

	return callable.Call(v, argValues)
}

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) any {
	var returnValue any = nil
	if ctx.Expr() != nil {
		returnValue = ctx.Expr().Accept(v)
	}
	panic(ReturnValueSignal{Value: returnValue})
}

func (f *RuntimeFunction) NrArgs() int {
	return len(f.Parameters)
}

func power(base, exp int) int {
	if exp < 0 {
		panic("Negative exponent not supported")
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// helper for standard print outputs
func cleanStringRepr(val any) string {
	if val == nil {
		return "nil"
	}

	switch v := val.(type) {
	case string:
		return v
	case *[]any:
		var sb strings.Builder
		sb.WriteString("[")
		for i, element := range *v {
			sb.WriteString(cleanStringRepr(element))
			if i < len(*v)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]")
		return sb.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
