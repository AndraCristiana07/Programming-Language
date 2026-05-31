package ast

import (
	"fmt"
	"my_language/parser"
	"slices"
	"sort"
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
	currEnv        *Environment
	StructRegistry map[string][]string
	MethodRegistry map[string]map[string]*parser.FuncStmtContext // for member methods
}

type ReturnValueSignal struct {
	Value any
}

type RuntimeFunction struct {
	Parameters []string
	Body       *parser.BlockStmtContext
	Env        *Environment
}

type BreakSignal struct{}
type ContinueSignal struct{}

type NullValue struct{}

var LanguageNull = &NullValue{}

func NewVisitor() *Visitor {
	return &Visitor{
		currEnv:        NewGlobalEnvironment(),
		StructRegistry: make(map[string][]string),
		MethodRegistry: make(map[string]map[string]*parser.FuncStmtContext),
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
	} else if ctx.StructStmt() != nil {
		return ctx.StructStmt().Accept(v)
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
	} else if ctx.SwitchStmt() != nil {
		return ctx.SwitchStmt().Accept(v)
	} else if ctx.ForInStmt() != nil {
		return ctx.ForInStmt().Accept(v)
	} else if ctx.ForStmt() != nil {
		return ctx.ForStmt().Accept(v)
	} else if ctx.PostfixStmt() != nil {
		return ctx.PostfixStmt().Accept(v)
	} else if ctx.TryCatchStmt() != nil {
		return ctx.TryCatchStmt().Accept(v)
	} else if ctx.ThrowStmt() != nil {
		return ctx.ThrowStmt().Accept(v)
	} else if ctx.ReturnStmt() != nil {
		return ctx.ReturnStmt().Accept(v)
	} else if ctx.BreakStmt() != nil {
		return ctx.BreakStmt().Accept(v)
	} else if ctx.ContinueStmt() != nil {
		return ctx.ContinueStmt().Accept(v)
	}

	return nil
}

func (v *Visitor) VisitStructStmt(ctx *parser.StructStmtContext) any {
	structName := ctx.GetId().GetText()
	fields := make([]string, 0)

	for _, fieldCtx := range ctx.AllStructField() {
		fieldName := fieldCtx.GetText()
		fields = append(fields, fieldName)
	}

	if v.StructRegistry == nil {
		v.StructRegistry = make(map[string][]string)
	}
	v.StructRegistry[structName] = fields

	return nil
}

func (v *Visitor) VisitStructField(ctx *parser.StructFieldContext) any {
	return ctx.IDENTIFIER().GetText()
}

func (v *Visitor) VisitStruct(ctx *parser.StructContext) any {
	structLitCtx := ctx.StructLiteral()
	structName := structLitCtx.GetStructName().GetText()

	validFields, defined := v.StructRegistry[structName]
	if !defined {
		panic(RuntimeError("TypeError", fmt.Sprintf("Struct '%s' is not defined", structName), ctx))
	}

	instance := make(map[string]any)
	instance["__type__"] = structName
	// initialize instance memory map with nulls
	for _, field := range validFields {
		instance[field] = LanguageNull
	}

	// populate field values
	for _, entryCtx := range structLitCtx.AllMapEntry() {
		keyStr := entryCtx.Expr(0).GetText()
		keyStr = strings.Trim(keyStr, "\"'") // clean quotes

		val := entryCtx.Expr(1).Accept(v)

		// field validation check
		isFieldValid := slices.Contains(validFields, keyStr)

		if !isFieldValid {
			panic(RuntimeError("TypeError", fmt.Sprintf("Struct '%s' has no field named '%s'", structName, keyStr), entryCtx))
		}

		instance[keyStr] = val
	}

	return &instance
}

func (v *Visitor) VisitForInit(ctx *parser.ForInitContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	} else if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	}
	return nil
}
func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) any {
	return BreakSignal{}
}

func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) any {
	return ContinueSignal{}
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
	v.currEnv.Define(varName, value)
	return nil
}

func (v *Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) any {
	leftHandSide := ctx.Expr(0)
	assignedValue := ctx.Expr(1).Accept(v)

	// nested
	container, indexVal := v.resolveAssignTarget(leftHandSide)

	if container != nil {
		switch obj := container.(type) {
		case *[]any:
			idx, ok := indexVal.(int)
			if !ok {
				panic(RuntimeError("TypeError", "Array lookup index position must resolve to an integer", ctx))
			}
			(*obj)[idx] = assignedValue
		case *map[string]any:
			keyStr := fmt.Sprintf("%v", indexVal)
			(*obj)[keyStr] = assignedValue
		default:
			panic(RuntimeError("TypeError", fmt.Sprintf("TypeError: Cannot assign index to type %T", container), ctx))
		}
	} else {
		// plain assign
		if identCtx, ok := leftHandSide.(*parser.IdentifierContext); ok {
			name := identCtx.IDENTIFIER().GetText()
			if !v.currEnv.Assign(name, assignedValue) {
				panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", name), ctx))

			}
		} else {
			panic(RuntimeError("SyntaxError", "Invalid assignment target", ctx))

		}
	}

	return nil
}

func (v *Visitor) VisitPostfixStmt(ctx *parser.PostfixStmtContext) any {
	leftHandSide := ctx.Expr()
	op := ctx.GetOp().GetText()

	currentValue := leftHandSide.Accept(v)

	intValue, ok := currentValue.(int)
	if !ok {
		panic(RuntimeError("TypeError", "Postfix increment/decrement target must be an integer", ctx))
	}

	switch op {
	case "++":
		intValue++
	case "--":
		intValue--
	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown or unsupported operator: %s", op), ctx))
	}

	// write result to destination
	container, indexVal := v.resolveAssignTarget(leftHandSide)

	if container != nil {
		// nested
		switch obj := container.(type) {
		case *[]any:
			idx, ok := indexVal.(int)
			if !ok {
				panic(RuntimeError("TypeError", "Array index must be an integer", ctx))
			}
			(*obj)[idx] = intValue
		case *map[string]any:
			keyStr := fmt.Sprintf("%v", indexVal)
			(*obj)[keyStr] = intValue
		default:
			panic(RuntimeError("TypeError", fmt.Sprintf("TypeError: Cannot assign index to type %T", container), ctx))

		}
	} else {
		// plain assign
		if identCtx, ok := leftHandSide.(*parser.IdentifierContext); ok {
			varName := identCtx.IDENTIFIER().GetText()
			if !v.currEnv.Assign(varName, intValue) {
				panic(RuntimeError("TypeError", fmt.Sprintf("Failed to assign value to variable '%s'", varName), ctx))
			}
		} else {
			panic(RuntimeError("SyntaxError", "Invalid postfix statement target", ctx))
		}
	}

	return nil
}
func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	numStr := ctx.NUMBER().GetText()

	// float
	if strings.Contains(numStr, ".") {
		val, _ := strconv.ParseFloat(numStr, 64)
		return val
	}

	// fall back to standard integers
	val, _ := strconv.Atoi(numStr)
	return val
}

func (v *Visitor) VisitMapLiteral(ctx *parser.MapLiteralContext) any {
	// init empty map
	nativeMap := make(map[string]any)

	// loop through key-value entries
	for _, entryCtx := range ctx.AllMapEntry() {
		rawKey := entryCtx.Expr(0).Accept(v)
		val := entryCtx.Expr(1).Accept(v)

		// make sure key converts to a string
		keyStr := fmt.Sprintf("%v", rawKey)
		nativeMap[keyStr] = val
	}

	return &nativeMap
}

func (v *Visitor) VisitNull(ctx *parser.NullContext) any {
	return LanguageNull
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	varName := ctx.IDENTIFIER().GetText()
	val, exists := v.currEnv.Lookup(varName)
	if !exists {
		panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", varName), ctx))
	}
	return val
}

func (v *Visitor) VisitString(ctx *parser.StringContext) any {
	rawStr := ctx.STRING().GetText()

	if unquoted, err := strconv.Unquote(rawStr); err == nil {
		return unquoted
	}

	// remove the surrounding quotes

	return rawStr[1 : len(rawStr)-1]
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
			panic(RuntimeError("SyntaxError", fmt.Sprintf("Unsupported operator for strings: %s", op), ctx))
		}
	}

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", fmt.Sprintf("Both operands must be integers for operator: %s", op), ctx))
	}

	switch op {
	case "+":
		return lVal + rVal
	case "-":
		return lVal - rVal
	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown operator: %s", op), ctx))
	}
}

func (v *Visitor) VisitMulDivMod(ctx *parser.MulDivModContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	if left == nil {
		panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", left), ctx))
	}
	if right == nil {
		panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", right), ctx))
	}

	switch op {
	case "*":
		return left.(int) * right.(int)
	case "/":
		if right.(int) == 0 {
			panic(RuntimeError("ZeroDivisionError", "Cannot divide a number by zero", ctx))
		}
		return left.(int) / right.(int)
	case "%":
		if right.(int) == 0 {
			panic(RuntimeError("ZeroModuloError", "Cannot modulo by zero", ctx))
		}
		return left.(int) % right.(int)
	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown operator: %s", op), ctx))

	}
}

func (v *Visitor) VisitExponential(ctx *parser.ExponentialContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", "Both operands must be integers for operator: **", ctx))
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

func (v *Visitor) VisitMembership(ctx *parser.MembershipContext) any {
	left := ctx.GetLeftExpr().Accept(v)
	right := ctx.GetRightExpr().Accept(v)

	op := ctx.GetOp().GetText()

	found := false

	// substring (string in string)
	if leftStr, okL := left.(string); okL {
		if rightStr, okR := right.(string); okR {
			lClean := strings.Trim(leftStr, "\"")
			rClean := strings.Trim(rightStr, "\"")
			found = strings.Contains(rClean, lClean)
		}
	}

	// element in array
	if arrPtr, ok := right.(*[]any); ok && arrPtr != nil {
		if slices.Contains(*arrPtr, left) {
			found = true
		}
	}

	// key lookup inside a dict/map
	if mapPtr, ok := right.(*map[string]any); ok && mapPtr != nil {
		keyStr := fmt.Sprintf("%v", left)
		if s, ok := left.(string); ok {
			keyStr = strings.Trim(s, "\"")
		}
		_, exists := (*mapPtr)[keyStr]
		found = exists
	}

	// check if contains "not"
	if strings.Contains(op, "not") {
		return !found
	}

	return found
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
		panic(RuntimeError("TypeError", "Comparison operator '<' requires integer operands", ctx))

	case ">":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal > rVal
		}
		panic(RuntimeError("TypeError", "Comparison operator '>' requires integer operands", ctx))

	case "<=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal <= rVal
		}
		panic(RuntimeError("TypeError", "Comparison operator '<=' requires integer operands", ctx))

	case ">=":
		lVal, lOk := left.(int)
		rVal, rOk := right.(int)
		if lOk && rOk {
			return lVal >= rVal
		}
		panic(RuntimeError("TypeError", "Comparison operator '>=' requires integer operands", ctx))

	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown operator: %s", op), ctx))
	}
}

func (v *Visitor) VisitPrintStmt(ctx *parser.PrintStmtContext) any {
	val := ctx.Expr().Accept(v)

	fmt.Println(cleanStringRepr(val))

	return nil
}

func (v *Visitor) VisitBlockStmt(ctx *parser.BlockStmtContext) any {
	for _, child := range ctx.GetChildren() {
		var res any

		if stmt, ok := child.(*parser.StatementContext); ok {
			if stmt != nil {
				res = stmt.Accept(v)
			}
		} else if fn, ok := child.(*parser.FuncStmtContext); ok {
			if fn != nil {
				res = fn.Accept(v)
			}
		}

		if _, ok := res.(BreakSignal); ok {
			return BreakSignal{}
		}
		if _, ok := res.(ContinueSignal); ok {
			return ContinueSignal{}
		}
	}
	return nil
}

func (v *Visitor) VisitTernaryOp(ctx *parser.TernaryOpContext) any {
	condition := ctx.GetCondExpr().Accept(v)

	if condCheck(condition) {
		return ctx.GetTrueExpr().Accept(v)
	} else {
		return ctx.GetFalseExpr().Accept(v)
	}

}

func (v *Visitor) VisitIfInit(ctx *parser.IfInitContext) any {
	if ctx.VarDecl() != nil {
		return ctx.VarDecl().Accept(v)
	}
	if ctx.AssignStmt() != nil {
		return ctx.AssignStmt().Accept(v)
	}
	return nil
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) any {
	originalEnv := v.currEnv

	// if init exists, make new scope
	if ctx.GetInit() != nil {
		v.currEnv = NewScope(originalEnv)
	}

	// ensure scope is restored
	defer func() {
		v.currEnv = originalEnv
	}()

	// execute the init stmt
	if ctx.GetInit() != nil {
		ctx.GetInit().Accept(v)
	}

	condition := ctx.Expr(0).Accept(v)

	// check if cond is true
	if condCheck(condition) {
		return ctx.GetThenBranch().Accept(v)
	}

	elifConditions := ctx.GetElifCond()
	elifBranches := ctx.GetElifBranch()

	// elif branches
	for i := 0; i < len(elifConditions); i++ {
		elifCondVal := elifConditions[i].Accept(v)

		if condCheck(elifCondVal) {
			return elifBranches[i].Accept(v)
		}
	}

	// else branch
	if ctx.GetElseBranch() != nil {
		return ctx.GetElseBranch().Accept(v)
	}

	return nil
}

func condCheck(val any) bool {
	if val == nil {
		return false
	}
	if b, ok := val.(bool); ok {
		return b
	}
	if f, ok := val.(float64); ok {
		return f != 0
	}
	// check if empty string
	if s, ok := val.(string); ok {
		return s != "" && s != "\"\""
	}
	return true
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) any {
	for {
		condition := ctx.Expr().Accept(v)
		conditionBool, ok := condition.(bool)
		if !ok {
			panic(RuntimeError("TypeError", "Condition in while statement must be boolean", ctx))
		}

		if !conditionBool {
			break
		}

		res := ctx.GetBody().Accept(v)

		if _, ok := res.(BreakSignal); ok {
			break // exit while loop
		}
		if _, ok := res.(ContinueSignal); ok {
			continue // jump to the next condition evaluation
		}
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
				panic(RuntimeError("TypeError", "For loop condition must evaluate to a boolean expression", ctx))
			}
			if !conditionBool {
				break
			}
		}

		var res any
		if ctx.GetBody() != nil {
			res = ctx.GetBody().Accept(v)
		}

		if _, ok := res.(BreakSignal); ok {
			break // exit the loop entirely
		}

		// execute post statement
		if ctx.GetPost() != nil {
			ctx.GetPost().Accept(v)
		}

		// handle continue after the post statement has run
		if _, ok := res.(ContinueSignal); ok {
			continue
		}
	}

	return nil
}

func (v *Visitor) VisitForInStmt(ctx *parser.ForInStmtContext) any {
	collection := ctx.Expr().Accept(v)

	varName := ctx.IDENTIFIER().GetText()

	// check if 'var' exixsts in assign
	hasVarKeyword := ctx.VAR() != nil

	previousEnv := v.currEnv
	v.currEnv = NewScope(previousEnv)
	defer func() {
		v.currEnv = previousEnv
	}()

	assignLoopVar := func(val any) {
		if hasVarKeyword {
			// for (var x in arr)
			v.currEnv.Define(varName, val)
		} else {
			// for (x in arr)
			if success := v.currEnv.Assign(varName, val); !success {
				v.currEnv.Define(varName, val)
			}
		}
	}

	executeBody := func() bool {
		// execute block body
		res := ctx.BlockStmt().Accept(v)
		if _, ok := res.(BreakSignal); ok {
			return false // break out of the loop
		}
		if _, ok := res.(ContinueSignal); ok {
			return true // skip to next iteration
		}
		return true
	}

	if targetPtr, ok := collection.(*[]any); ok && targetPtr != nil {
		// dereference the pointer
		for _, value := range *targetPtr {
			assignLoopVar(value)

			if !executeBody() {
				return nil
			}
		}
		return nil
	}

	//string iteration
	if str, ok := collection.(string); ok {
		// clean up quotes
		cleanedStr := str
		if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
			cleanedStr = str[1 : len(str)-1]
		}

		// loop through the string by rune
		for _, runeVal := range cleanedStr {
			// convert rune back into single-character string
			charStr := string(runeVal)

			assignLoopVar(charStr)

			if !executeBody() {
				return nil
			}
		}
		return nil
	}

	//if non-pointer slice
	if targetSlice, ok := collection.([]any); ok {
		for _, value := range targetSlice {
			assignLoopVar(value)
			if !executeBody() {
				return nil
			}
		}
		return nil
	}

	// dictionary/map values
	if targetMap, ok := collection.(*map[string]any); ok && targetMap != nil {
		for _, value := range *targetMap {
			assignLoopVar(value)
			if !executeBody() {
				return nil
			}
		}
		return nil
	}

	panic(fmt.Sprintf("TypeError: Cannot iterate over type %T using for...in", collection))
}

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) any {
	// eval switch expt
	switchExpr := ctx.Expr().Accept(v)

	match := false

	// loop through cases
	for _, caseCtx := range ctx.AllCaseBlock() {
		if !match {
			caseValue := caseCtx.Expr().Accept(v)

			if switchExpr == caseValue {
				match = true
			}
		}
		// execute stmt if there;s a match
		if match {
			for _, stmt := range caseCtx.AllStatement() {
				res := stmt.Accept(v)
				// break hit
				if _, ok := res.(BreakSignal); ok {
					return nil
				}
			}
		}
	}
	// default block
	if !match && ctx.DefaultBlock() != nil {
		defaultCtx := ctx.DefaultBlock()
		for _, stmt := range defaultCtx.AllStatement() {
			res := stmt.Accept(v)
			// break hit
			if _, ok := res.(BreakSignal); ok {
				return nil
			}
		}
	}
	return nil
}

func (v *Visitor) VisitCompoundAssignStmt(ctx *parser.CompoundAssignStmtContext) any {
	leftHandSide := ctx.Expr(0)
	currentValue := leftHandSide.Accept(v)
	value := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	var result any

	// string concatenation
	if op == "+=" {
		if strVal, ok := currentValue.(string); ok {
			result = strVal + cleanStringRepr(value)
		}
	}

	// int math
	if result == nil {
		intCurrentValue, ok := currentValue.(int)
		if !ok {
			panic(RuntimeError("TypeError", "Left-hand side of compound assignment must evaluate to an integer or string", ctx))
		}

		intValue, ok := value.(int)
		if !ok {
			panic(RuntimeError("TypeError", "Right-hand side of compound assignment must be an integer", ctx))
		}

		switch op {
		case "+=":
			result = intCurrentValue + intValue
		case "-=":
			result = intCurrentValue - intValue
		case "*=":
			result = intCurrentValue * intValue
		case "/=":
			if intValue == 0 {
				panic(RuntimeError("ZeroDivisionError", "Cannot divide a number by zero", ctx))
			}
			result = intCurrentValue / intValue
		case "%=":
			if intValue == 0 {
				panic(RuntimeError("ZeroDivisionError", "Cannot modulo by zero", ctx))
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
			panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown compound assignment operator: %s", op), ctx))
		}
	}

	// write result to destination
	container, indexVal := v.resolveAssignTarget(leftHandSide)

	if container != nil {
		// nested
		switch obj := container.(type) {
		case *[]any:
			idx, ok := indexVal.(int)
			if !ok {
				panic(RuntimeError("TypeError", "Array index must be an integer", ctx))
			}

			(*obj)[idx] = result
		case *map[string]any:
			keyStr := fmt.Sprintf("%v", indexVal)
			(*obj)[keyStr] = result
		default:
			panic(RuntimeError("TypeError", fmt.Sprintf("Cannot assign index to type %T", container), ctx))
		}
	} else {
		// plain assign
		if identCtx, ok := leftHandSide.(*parser.IdentifierContext); ok {
			varName := identCtx.IDENTIFIER().GetText()
			if !v.currEnv.Assign(varName, result) {
				panic(RuntimeError("TypeError", fmt.Sprintf("Failed to assign value to variable '%s'", varName), ctx))
			}
		} else {
			panic(RuntimeError("SyntaxError", "Invalid compound assignment target", ctx))
		}
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
			panic(RuntimeError("TypeError", "Operand of 'not' must be boolean", ctx))
		}
		return !boolVal
	case "~":
		intVal, ok := value.(int)
		if !ok {
			panic(RuntimeError("TypeError", "Operand of '~' must be boolean", ctx))
		}
		return ^intVal
	case "-":
		intVal, ok := value.(int)
		if !ok {
			panic(RuntimeError("TypeError", "Unary minus operator can only be applied to an integer", ctx))

		}
		return -intVal
	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown unary operator: %s", op), ctx))
	}
}

func (v *Visitor) VisitBitShift(ctx *parser.BitShiftContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	op := ctx.GetOp().GetText()

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", "Both operands must be integers for bit shift operators", ctx))
	}

	switch op {
	case "<<":
		return lVal << rVal
	case ">>":
		return lVal >> rVal
	default:
		panic(RuntimeError("SyntaxError", fmt.Sprintf("Unknown bit shift operator: %s", op), ctx))

	}
}

func (v *Visitor) VisitBitAnd(ctx *parser.BitAndContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", "Both operands must be integers for bitwise AND operator", ctx))
	}

	return lVal & rVal
}

func (v *Visitor) VisitBitXor(ctx *parser.BitXorContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", "Both operands must be integers for bitwise XOR operator", ctx))
	}

	return lVal ^ rVal
}

func (v *Visitor) VisitBitOr(ctx *parser.BitOrContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)

	lVal, lOk := left.(int)
	rVal, rOk := right.(int)
	if !lOk || !rOk {
		panic(RuntimeError("TypeError", "Both operands must be integers for bitwise OR operator", ctx))
	}

	return lVal | rVal
}

func (v *Visitor) VisitAnd(ctx *parser.AndContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	leftBool, leftOk := left.(bool)
	rightBool, rightOk := right.(bool)
	if !leftOk || !rightOk {
		panic(RuntimeError("TypeError", "Operands of 'and' must be boolean", ctx))

	}
	return leftBool && rightBool
}

func (v *Visitor) VisitOr(ctx *parser.OrContext) any {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	leftBool, leftOk := left.(bool)
	rightBool, rightOk := right.(bool)
	if !leftOk || !rightOk {
		panic(RuntimeError("TypeError", "Operands of 'or' must be boolean", ctx))
	}
	return leftBool || rightBool
}

func (v *Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) any {
	return ctx.ArrayLit().Accept(v)
}

func (v *Visitor) VisitStandardArray(ctx *parser.StandardArrayContext) any {
	var elements []any
	for _, expr := range ctx.AllExpr() {
		elements = append(elements, expr.Accept(v))
	}
	// return a pointer to the slice
	return &elements
}

func (v *Visitor) VisitListComprehension(ctx *parser.ListComprehensionContext) any {
	sourceCollection := ctx.GetSrcExpr().Accept(v)

	// loop iterator variable name
	varName := ctx.IDENTIFIER().GetText()

	// new slice that will hold results
	resultSlice := make([]any, 0)

	previousEnv := v.currEnv
	v.currEnv = NewScope(previousEnv)

	defer func() {
		v.currEnv = previousEnv
	}()

	processElement := func(element any) {
		// assigncurrent item to the loop variable
		v.currEnv.Define(varName, element)

		// check for if
		if ctx.GetFilterExpr() != nil {
			conditionValue := ctx.GetFilterExpr().Accept(v)
			if !condCheck(conditionValue) {
				return
			}
		}

		// eval expression
		transformedValue := ctx.GetTransformExpr().Accept(v)
		// append it to result
		resultSlice = append(resultSlice, transformedValue)
	}

	// loop over the source
	if arrPtr, ok := sourceCollection.(*[]any); ok && arrPtr != nil {
		for _, value := range *arrPtr {
			processElement(value)
		}
	} else if str, ok := sourceCollection.(string); ok {
		// if string -> clean and process
		cleanedStr := str
		if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
			cleanedStr = str[1 : len(str)-1]
		}
		for _, runeVal := range cleanedStr {
			processElement(string(runeVal))
		}
	} else {
		panic(RuntimeError("TypeError", fmt.Sprintf("Cannot perform list comprehension over type %T", sourceCollection), ctx))
	}

	return &resultSlice
}

func (v *Visitor) VisitArrayIndex(ctx *parser.ArrayIndexContext) any {
	collection := ctx.Expr(0).Accept(v)
	indexKey := ctx.Expr(1).Accept(v)

	// expect pointer to the slice (*[]any)
	if arrPtr, ok := collection.(*[]any); ok && arrPtr != nil {
		// dereference it locally to perform size checks and lookups safely
		arr := *arrPtr

		idx, ok := indexKey.(int)
		if !ok {
			// try converting to int
			if f, ok := indexKey.(float64); ok {
				idx = int(f)
			} else {
				panic(RuntimeError("TypeError", "Array index must be an integer", ctx))
			}
		}

		if idx < 0 || idx >= len(arr) {
			panic(RuntimeError("IndexError", fmt.Sprintf("Array index %d is out of bounds (length %d)", idx, len(arr)), ctx))
		}

		return arr[idx]
	}

	// dict/map
	if mapPtr, ok := collection.(*map[string]any); ok && mapPtr != nil {
		targetMap := *mapPtr

		// indexKey into string for map lookup
		var lookupKey string
		switch k := indexKey.(type) {
		case string:
			lookupKey = k
		case bool:
			// true -> "true", false -> "false"
			if k {
				lookupKey = "true"
			} else {
				lookupKey = "false"
			}
		case int:
			lookupKey = strconv.Itoa(k)
		case float64:
			lookupKey = strconv.FormatFloat(k, 'f', -1, 64)
		default:
			lookupKey = fmt.Sprintf("%v", indexKey)
		}

		// lookup inside map
		if val, exists := targetMap[lookupKey]; exists {
			return val
		}

		return nil
	}

	panic(RuntimeError("TypeError", fmt.Sprintf("Cannot read properties of type %T using brackets [...]", collection), ctx))
}

func (v *Visitor) VisitArrayAssignStmt(ctx *parser.ArrayAssignStmtContext) any {
	arrayName := ctx.IDENTIFIER().GetText()

	val, exists := v.currEnv.Lookup(arrayName)
	if !exists {
		panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", arrayName), ctx))

	}

	arrPtr, ok := val.(*[]any)
	if !ok {
		panic(RuntimeError("TypeError", fmt.Sprintf("Attempting to index a non-array variable: %s", arrayName), ctx))
	}

	indexVal := ctx.Expr(0).Accept(v)
	newVal := ctx.Expr(1).Accept(v)

	idx, ok := indexVal.(int)
	if !ok {
		panic(RuntimeError("TypeError", "Array index must be an integer", ctx))
	}

	arr := *arrPtr
	if idx < 0 || idx >= len(arr) {
		panic(RuntimeError("IndexError", fmt.Sprintf("Array index %d is out of bounds (length %d)", idx, len(arr)), ctx))
	}

	(*arrPtr)[idx] = newVal

	return newVal
}

func (v *Visitor) VisitFuncStmt(ctx *parser.FuncStmtContext) any {
	funcName := ctx.IDENTIFIER(0).GetText()

	// check if it's from a struct receiver
	if ctx.Receiver() != nil {
		receiverCtx := ctx.Receiver().(*parser.ReceiverContext)
		structTypeName := receiverCtx.GetStructType().GetText()

		// init the inner struct method map
		if v.MethodRegistry[structTypeName] == nil {
			v.MethodRegistry[structTypeName] = make(map[string]*parser.FuncStmtContext)
		}

		// register the raw execution context
		v.MethodRegistry[structTypeName][funcName] = ctx
		return nil
	}

	// standard global function
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

func (v *Visitor) VisitMethodCall(ctx *parser.MethodCallContext) any {
	obj := ctx.Expr(0).Accept(v)

	//  verify it's a modifiable struct instance
	mapPtr, ok := obj.(*map[string]any)
	if !ok || mapPtr == nil {
		panic(RuntimeError("TypeError", "Cannot call method of a non-object type", ctx))
	}

	// extract the instance type tag metadata
	structType := (*mapPtr)["__type__"].(string)
	methodName := ctx.IDENTIFIER().GetText()

	// look up the method in  Method Registry
	structMethods, foundStruct := v.MethodRegistry[structType]
	if !foundStruct {
		panic(RuntimeError("TypeError", fmt.Sprintf("Struct '%s' has no methods defined", structType), ctx))
	}

	methodCtx, foundMethod := structMethods[methodName]
	if !foundMethod {
		panic(RuntimeError("TypeError", fmt.Sprintf("Struct '%s' has no method named '%s'", structType, methodName), ctx))
	}

	// eval the argument values
	args := make([]any, 0)
	for _, exprCtx := range ctx.AllExpr() {
		args = append(args, exprCtx.Accept(v))
	}

	previousEnv := v.currEnv
	methodEnv := NewEnvironment()

	// bind the receiver object instance to the variable name
	receiverCtx := methodCtx.Receiver().(*parser.ReceiverContext)
	receiverVarName := receiverCtx.GetId().GetText()
	methodEnv.Define(receiverVarName, mapPtr)

	// bind the rest of the standard arguments to parameters
	paramIds := methodCtx.AllIDENTIFIER()
	// skip index 0 (method name)
	for i, paramCtx := range paramIds[1:] {
		if i < len(args) {
			methodEnv.Define(paramCtx.GetText(), args[i])
		} else {
			methodEnv.Define(paramCtx.GetText(), LanguageNull)
		}
	}

	// execute the body block scope statement loops safely
	v.currEnv = methodEnv
	var methodResult any = LanguageNull

	func() {
		defer func() {
			if r := recover(); r != nil {
				if ret, isReturn := r.(ReturnValueSignal); isReturn {
					methodResult = ret.Value
				} else {
					panic(r) // bubble up normal runtime panics
				}
			}
		}()
		methodCtx.BlockStmt().Accept(v)
	}()

	//go back to previous env
	v.currEnv = previousEnv
	return methodResult
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
		panic(RuntimeError("ReferenceError", fmt.Sprintf("Function '%s' is not defined", funcName), ctx))

	}

	callable, ok := fn.(Callable)
	if !ok {
		panic(RuntimeError("TypeError", fmt.Sprintf("'%s' is not a callable function", funcName), ctx))
	}

	// evaluate the call arguments at the caller's execution environment line
	var argValues []any
	for _, argExpr := range ctx.AllExpr() {
		argValues = append(argValues, argExpr.Accept(v))
	}

	// validate input count parameters
	if callable.NrArgs() != -1 && len(argValues) != callable.NrArgs() {
		panic(RuntimeError("TypeError", fmt.Sprintf("Function %s expects %d arguments, got %d", funcName, callable.NrArgs(), len(argValues)), ctx))
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

func (v *Visitor) VisitFieldAccess(ctx *parser.FieldAccessContext) any {
	obj := ctx.Expr().Accept(v)

	propName := ctx.IDENTIFIER().GetText()

	if obj == nil || obj == LanguageNull {
		panic(RuntimeError("TypeError", fmt.Sprintf("Cannot read property '%s' of null", propName), ctx))
	}

	switch target := obj.(type) {

	case *map[string]any:
		m := *target
		if val, exists := m[propName]; exists {
			return val
		}
		return nil

	case map[string]any:
		if val, exists := target[propName]; exists {
			return val
		}
		return nil

	default:
		panic(RuntimeError("TypeError", fmt.Sprintf("Cannot read property '%s' of type %T", propName, obj), ctx))
	}
}

func (f *RuntimeFunction) NrArgs() int {
	return len(f.Parameters)
}

func (v *Visitor) VisitTryCatchStmt(ctx *parser.TryCatchStmtContext) any {
	// return type from string to any to keep the raw map pointer
	panicked, caughtErrorObj := func() (isPanic bool, rawErr any) {
		defer func() {
			if r := recover(); r != nil {
				isPanic = true
				rawErr = r
			}
		}()

		// run try block
		ctx.GetTryBody().Accept(v)
		// no panics
		return false, nil
	}()

	// error caught -> route flow directly into catch scope
	if panicked {
		catchVarName := ctx.IDENTIFIER().GetText()

		// env scope for catch block
		previousEnv := v.currEnv
		v.currEnv = NewScope(previousEnv)

		v.currEnv.Define(catchVarName, caughtErrorObj)

		ctx.GetCatchBody().Accept(v)

		// restore original parent environment context
		v.currEnv = previousEnv
	}

	return nil
}
func (v *Visitor) VisitThrowStmt(ctx *parser.ThrowStmtContext) any {
	// eval the expression thrown
	thrownVal := ctx.Expr().Accept(v)

	panic(cleanStringRepr(thrownVal))
}

func RuntimeError(errorType string, message string, ctx antlr.ParserRuleContext) *map[string]any {
	line := 0
	if ctx != nil && ctx.GetStart() != nil {
		line = ctx.GetStart().GetLine()
	}

	// map for the Error Object
	errObj := map[string]any{
		"type":    errorType,
		"message": message,
		"line":    line,
		"text":    fmt.Sprintf("[%s Layer Rule] Line %d: %s", errorType, line, message),
	}

	return &errObj
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
	if val == nil || val == LanguageNull {
		return "null"
	}

	switch v := val.(type) {
	case string:
		return v
	case float64:
		// check if whole nr
		if v == float64(int64(v)) {
			return fmt.Sprintf("%.1f", v) // 5.0
		}
		return fmt.Sprintf("%g", v) //5.3
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
	case *map[string]any:
		m := *v

		// struct
		if typeName, isStruct := m["__type__"].(string); isStruct {
			var sb strings.Builder
			sb.WriteString(typeName + "{")

			// get field names
			fields := make([]string, 0)
			for k := range m {
				if k != "__type__" {
					fields = append(fields, k)
				}
			}
			sort.Strings(fields)

			for i, key := range fields {
				fmt.Fprintf(&sb, "%s: %s", key, cleanStringRepr(m[key]))
				if i < len(fields)-1 {
					sb.WriteString(", ")
				}
			}
			sb.WriteString("}")
			return sb.String()
		}

		// dict/map
		var sb strings.Builder
		sb.WriteString("{")
		i := 0
		for key, element := range m {
			fmt.Fprintf(&sb, "%s: %s", key, cleanStringRepr(element))
			if i < len(m)-1 {
				sb.WriteString(", ")
			}
			i++
		}
		sb.WriteString("}")
		return sb.String()

	default:
		return fmt.Sprintf("%v", v)

	}
}

func (v *Visitor) resolveAssignTarget(leftCtx parser.IExprContext) (any, any) {
	// braket lookup
	if idxCtx, ok := leftCtx.(*parser.ArrayIndexContext); ok {
		innerTargetCtx := idxCtx.Expr(0)
		indexVal := idxCtx.Expr(1).Accept(v)

		// nested (matrix[x][y])
		if _, isNestedIdx := innerTargetCtx.(*parser.ArrayIndexContext); isNestedIdx {
			parentContainer, parentIndex := v.resolveAssignTarget(innerTargetCtx)
			return v.unwrapContainer(parentContainer, parentIndex), indexVal
		}

		// nested field access (user.items[idx])
		if _, isNestedField := innerTargetCtx.(*parser.FieldAccessContext); isNestedField {
			parentContainer, parentIndex := v.resolveAssignTarget(innerTargetCtx)
			return v.unwrapContainer(parentContainer, parentIndex), indexVal
		}

		// plain struct (items[idx])
		if identCtx, ok := innerTargetCtx.(*parser.IdentifierContext); ok {
			name := identCtx.GetText()
			container, exists := v.currEnv.Lookup(name)
			if !exists {
				panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", name), leftCtx))
			}
			return container, indexVal
		}
	}

	// field acces by dot
	if fieldCtx, ok := leftCtx.(*parser.FieldAccessContext); ok {
		innerTargetCtx := fieldCtx.Expr()
		fieldName := fieldCtx.IDENTIFIER().GetText()

		// prevent runtime overwrite of internal system struct metadata
		if fieldName == "__type__" {
			panic(RuntimeError("TypeError", "Cannot overwrite internal system metadata '__type__'", leftCtx))
		}

		// nested index field lookup (party.members[0].active)
		if _, isNestedIdx := innerTargetCtx.(*parser.ArrayIndexContext); isNestedIdx {
			parentContainer, parentIndex := v.resolveAssignTarget(innerTargetCtx)
			return v.unwrapContainer(parentContainer, parentIndex), fieldName
		}

		// nested
		if _, isNestedField := innerTargetCtx.(*parser.FieldAccessContext); isNestedField {
			parentContainer, parentIndex := v.resolveAssignTarget(innerTargetCtx)
			return v.unwrapContainer(parentContainer, parentIndex), fieldName
		}

		// plain case
		if identCtx, ok := innerTargetCtx.(*parser.IdentifierContext); ok {
			name := identCtx.GetText()
			container, exists := v.currEnv.Lookup(name)
			if !exists {
				panic(RuntimeError("ReferenceError", fmt.Sprintf("Variable '%s' is not defined", name), leftCtx))
			}
			return container, fieldName
		}
	}

	return nil, nil
}

func (v *Visitor) unwrapContainer(parentContainer any, parentIndex any) any {
	switch obj := parentContainer.(type) {
	case *[]any:
		return (*obj)[parentIndex.(int)]
	case *map[string]any:
		return (*obj)[fmt.Sprintf("%v", parentIndex)]
	default:
		panic(RuntimeError("TypeError", fmt.Sprintf("TypeError: %T is not an indexable container", parentContainer), nil))
	}
}
