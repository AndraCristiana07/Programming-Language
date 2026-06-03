// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'null'", "'interface'", "'struct'", "'switch'", "'case'", "'default'",
		"'try'", "'catch'", "'throw'", "'in'", "'.'", "':'", "'var'", "'='",
		"'+'", "'-'", "'*'", "'/'", "'<'", "'>'", "'=='", "'!='", "'<='", "'>='",
		"'{'", "'('", "')'", "'}'", "'['", "']'", "'if'", "'else'", "'elif'",
		"'while'", "'for'", "'break'", "'continue'", "';'", "','", "'++'", "'--'",
		"'%'", "'**'", "'+='", "'-='", "'*='", "'/='", "'%='", "'**='", "'&'",
		"'|'", "'^'", "'<<'", "'>>'", "'~'", "'&='", "'|='", "'^='", "'<<='",
		"'>>='", "'lambda'", "'func'", "'return'", "'print'", "'true'", "'false'",
		"'and'", "'or'", "'not'",
	}
	staticData.SymbolicNames = []string{
		"", "", "INTERFACE", "STRUCT", "SWITCH", "CASE", "DEFAULT", "TRY", "CATCH",
		"THROW", "IN", "DOT", "COLON", "VAR", "EQUALS", "PLUS", "MINUS", "STAR",
		"SLASH", "LESS", "GREATER", "EQUALEQUAL", "BANGEQUAL", "LESSEQUAL",
		"GREATEREQUAL", "LBRACE", "LPAREN", "RPAREN", "RBRACE", "LBRACKET",
		"RBRACKET", "IF", "ELSE", "ELIF", "WHILE", "FOR", "BREAK", "CONTINUE",
		"SEMICOLON", "COMMA", "INC", "DEC", "MODULO", "EXPONENTIAL", "PLUSEQUAL",
		"MINUSEQUAL", "STAREQUAL", "SLASHEQUAL", "MODEQUAL", "EXPONENTIALEQUAL",
		"BITAND", "BITOR", "BITXOR", "BITLSHIFT", "BITRSHIFT", "BITNOT", "BITANDEQUAL",
		"BITOREQUAL", "BITXOREQAUL", "BITLSHIFTEQUAL", "BITRSHIFTEQUAL", "LAMBDA",
		"FUNC", "RETURN", "PRINT", "TRUE", "FALSE", "AND", "OR", "NOT", "IDENTIFIER",
		"NUMBER", "STRING", "LINE_COMMENT", "BLOCK_COMMENT", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "returnStmt", "terminator", "funcStmt", "receiver",
		"interfaceStmt", "method", "exprStmt", "tryCatchStmt", "throwStmt",
		"breakStmt", "continueStmt", "switchStmt", "caseBlock", "defaultBlock",
		"forInStmt", "loopTarget", "identifierList", "varDecl", "assignStmt",
		"arrayAssignStmt", "compoundAssignStmt", "printStmt", "ifStmt", "whileStmt",
		"forStmt", "blockStmt", "structStmt", "structField", "ifInit", "forInit",
		"forPost", "postfixStmt", "expr", "structLiteral", "mapEntry", "membershipOp",
		"arrayLit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 76, 621, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 5, 0, 82, 8, 0, 10, 0, 12,
		0, 85, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 109, 8, 1, 1, 1, 3, 1, 112, 8, 1, 1, 2, 1, 2, 3, 2, 116, 8, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 3, 4, 122, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5,
		4, 129, 8, 4, 10, 4, 12, 4, 132, 9, 4, 3, 4, 134, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 148, 8,
		6, 10, 6, 12, 6, 151, 9, 6, 1, 6, 1, 6, 4, 6, 155, 8, 6, 11, 6, 12, 6,
		156, 1, 6, 1, 6, 5, 6, 161, 8, 6, 10, 6, 12, 6, 164, 9, 6, 1, 6, 5, 6,
		167, 8, 6, 10, 6, 12, 6, 170, 9, 6, 3, 6, 172, 8, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 181, 8, 7, 10, 7, 12, 7, 184, 9, 7, 3, 7,
		186, 8, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 214, 8, 13, 10, 13, 12, 13, 217,
		9, 13, 1, 13, 5, 13, 220, 8, 13, 10, 13, 12, 13, 223, 9, 13, 1, 13, 3,
		13, 226, 8, 13, 1, 13, 5, 13, 229, 8, 13, 10, 13, 12, 13, 232, 9, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 241, 8, 14, 10, 14,
		12, 14, 244, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 250, 8, 15, 10,
		15, 12, 15, 253, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 3, 17, 264, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 271, 8, 17, 1, 18, 1, 18, 1, 18, 4, 18, 276, 8, 18, 11, 18, 12, 18,
		277, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 286, 8, 19, 10, 19,
		12, 19, 289, 9, 19, 1, 19, 3, 19, 292, 8, 19, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 3, 24, 320, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 5, 24, 331, 8, 24, 10, 24, 12, 24, 334, 9, 24, 1, 24,
		1, 24, 3, 24, 338, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 5, 27, 360, 8, 27, 10, 27, 12, 27, 363, 9, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 371, 8, 28, 10, 28, 12, 28,
		374, 9, 28, 1, 28, 1, 28, 4, 28, 378, 8, 28, 11, 28, 12, 28, 379, 1, 28,
		1, 28, 5, 28, 384, 8, 28, 10, 28, 12, 28, 387, 9, 28, 1, 28, 5, 28, 390,
		8, 28, 10, 28, 12, 28, 393, 9, 28, 3, 28, 395, 8, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 3, 30, 403, 8, 30, 1, 31, 1, 31, 3, 31, 407, 8,
		31, 1, 32, 1, 32, 3, 32, 411, 8, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 422, 8, 34, 10, 34, 12, 34, 425, 9,
		34, 3, 34, 427, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 5, 34, 438, 8, 34, 10, 34, 12, 34, 441, 9, 34, 3, 34, 443,
		8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 450, 8, 34, 10, 34, 12,
		34, 453, 9, 34, 3, 34, 455, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 5, 34, 465, 8, 34, 10, 34, 12, 34, 468, 9, 34, 1, 34,
		3, 34, 471, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 486, 8, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 536, 8, 34, 1, 34, 1, 34, 3, 34, 540,
		8, 34, 1, 34, 1, 34, 3, 34, 544, 8, 34, 3, 34, 546, 8, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 556, 8, 34, 10, 34, 12,
		34, 559, 9, 34, 3, 34, 561, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 567,
		8, 34, 10, 34, 12, 34, 570, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5,
		35, 577, 8, 35, 10, 35, 12, 35, 580, 9, 35, 3, 35, 582, 8, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 3, 37, 593, 8, 37,
		1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 599, 8, 38, 10, 38, 12, 38, 602, 9,
		38, 3, 38, 604, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 3, 38, 615, 8, 38, 1, 38, 1, 38, 3, 38, 619, 8, 38, 1, 38,
		0, 1, 68, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 0, 9, 2, 0, 38, 38, 75, 75, 2, 0, 44, 49, 56, 60, 1,
		0, 40, 41, 3, 0, 16, 16, 55, 55, 69, 69, 1, 0, 65, 66, 2, 0, 17, 18, 42,
		42, 1, 0, 15, 16, 1, 0, 53, 54, 1, 0, 19, 24, 692, 0, 83, 1, 0, 0, 0, 2,
		108, 1, 0, 0, 0, 4, 113, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 119, 1, 0,
		0, 0, 10, 138, 1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14, 175, 1, 0, 0, 0, 16,
		190, 1, 0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 200, 1, 0, 0, 0, 22, 203, 1,
		0, 0, 0, 24, 205, 1, 0, 0, 0, 26, 207, 1, 0, 0, 0, 28, 235, 1, 0, 0, 0,
		30, 245, 1, 0, 0, 0, 32, 254, 1, 0, 0, 0, 34, 270, 1, 0, 0, 0, 36, 272,
		1, 0, 0, 0, 38, 279, 1, 0, 0, 0, 40, 296, 1, 0, 0, 0, 42, 300, 1, 0, 0,
		0, 44, 307, 1, 0, 0, 0, 46, 311, 1, 0, 0, 0, 48, 314, 1, 0, 0, 0, 50, 339,
		1, 0, 0, 0, 52, 345, 1, 0, 0, 0, 54, 355, 1, 0, 0, 0, 56, 366, 1, 0, 0,
		0, 58, 398, 1, 0, 0, 0, 60, 402, 1, 0, 0, 0, 62, 406, 1, 0, 0, 0, 64, 410,
		1, 0, 0, 0, 66, 412, 1, 0, 0, 0, 68, 485, 1, 0, 0, 0, 70, 571, 1, 0, 0,
		0, 72, 585, 1, 0, 0, 0, 74, 592, 1, 0, 0, 0, 76, 618, 1, 0, 0, 0, 78, 82,
		3, 2, 1, 0, 79, 82, 3, 8, 4, 0, 80, 82, 3, 6, 3, 0, 81, 78, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86,
		87, 5, 0, 0, 1, 87, 1, 1, 0, 0, 0, 88, 109, 3, 38, 19, 0, 89, 109, 3, 40,
		20, 0, 90, 109, 3, 56, 28, 0, 91, 109, 3, 12, 6, 0, 92, 109, 3, 42, 21,
		0, 93, 109, 3, 44, 22, 0, 94, 109, 3, 16, 8, 0, 95, 109, 3, 46, 23, 0,
		96, 109, 3, 48, 24, 0, 97, 109, 3, 50, 25, 0, 98, 109, 3, 26, 13, 0, 99,
		109, 3, 32, 16, 0, 100, 109, 3, 52, 26, 0, 101, 109, 3, 54, 27, 0, 102,
		109, 3, 66, 33, 0, 103, 109, 3, 18, 9, 0, 104, 109, 3, 20, 10, 0, 105,
		109, 3, 4, 2, 0, 106, 109, 3, 22, 11, 0, 107, 109, 3, 24, 12, 0, 108, 88,
		1, 0, 0, 0, 108, 89, 1, 0, 0, 0, 108, 90, 1, 0, 0, 0, 108, 91, 1, 0, 0,
		0, 108, 92, 1, 0, 0, 0, 108, 93, 1, 0, 0, 0, 108, 94, 1, 0, 0, 0, 108,
		95, 1, 0, 0, 0, 108, 96, 1, 0, 0, 0, 108, 97, 1, 0, 0, 0, 108, 98, 1, 0,
		0, 0, 108, 99, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0, 108, 101, 1, 0, 0, 0,
		108, 102, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 108, 104, 1, 0, 0, 0, 108,
		105, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 111,
		1, 0, 0, 0, 110, 112, 3, 6, 3, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0,
		0, 0, 112, 3, 1, 0, 0, 0, 113, 115, 5, 63, 0, 0, 114, 116, 3, 68, 34, 0,
		115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 5, 1, 0, 0, 0, 117, 118,
		7, 0, 0, 0, 118, 7, 1, 0, 0, 0, 119, 121, 5, 62, 0, 0, 120, 122, 3, 10,
		5, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0,
		123, 124, 5, 70, 0, 0, 124, 133, 5, 26, 0, 0, 125, 130, 5, 70, 0, 0, 126,
		127, 5, 39, 0, 0, 127, 129, 5, 70, 0, 0, 128, 126, 1, 0, 0, 0, 129, 132,
		1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 134, 1, 0,
		0, 0, 132, 130, 1, 0, 0, 0, 133, 125, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 136, 5, 27, 0, 0, 136, 137, 3, 54, 27, 0, 137,
		9, 1, 0, 0, 0, 138, 139, 5, 26, 0, 0, 139, 140, 5, 70, 0, 0, 140, 141,
		5, 70, 0, 0, 141, 142, 5, 27, 0, 0, 142, 11, 1, 0, 0, 0, 143, 144, 5, 2,
		0, 0, 144, 145, 5, 70, 0, 0, 145, 149, 5, 25, 0, 0, 146, 148, 3, 6, 3,
		0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 171, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 162,
		3, 14, 7, 0, 153, 155, 3, 6, 3, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 159, 3, 14, 7, 0, 159, 161, 1, 0, 0, 0, 160, 154, 1, 0, 0, 0, 161,
		164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 168,
		1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 167, 3, 6, 3, 0, 166, 165, 1, 0,
		0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0,
		169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 152, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 5, 28, 0, 0, 174, 13,
		1, 0, 0, 0, 175, 176, 5, 70, 0, 0, 176, 185, 5, 26, 0, 0, 177, 182, 5,
		70, 0, 0, 178, 179, 5, 39, 0, 0, 179, 181, 5, 70, 0, 0, 180, 178, 1, 0,
		0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0,
		183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 177, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 5, 27, 0, 0, 188, 189,
		5, 38, 0, 0, 189, 15, 1, 0, 0, 0, 190, 191, 3, 68, 34, 0, 191, 17, 1, 0,
		0, 0, 192, 193, 5, 7, 0, 0, 193, 194, 3, 54, 27, 0, 194, 195, 5, 8, 0,
		0, 195, 196, 5, 26, 0, 0, 196, 197, 5, 70, 0, 0, 197, 198, 5, 27, 0, 0,
		198, 199, 3, 54, 27, 0, 199, 19, 1, 0, 0, 0, 200, 201, 5, 9, 0, 0, 201,
		202, 3, 68, 34, 0, 202, 21, 1, 0, 0, 0, 203, 204, 5, 36, 0, 0, 204, 23,
		1, 0, 0, 0, 205, 206, 5, 37, 0, 0, 206, 25, 1, 0, 0, 0, 207, 208, 5, 4,
		0, 0, 208, 209, 5, 26, 0, 0, 209, 210, 3, 68, 34, 0, 210, 211, 5, 27, 0,
		0, 211, 215, 5, 25, 0, 0, 212, 214, 3, 6, 3, 0, 213, 212, 1, 0, 0, 0, 214,
		217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 221,
		1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 220, 3, 28, 14, 0, 219, 218, 1,
		0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0,
		0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 226, 3, 30, 15, 0,
		225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 230, 1, 0, 0, 0, 227,
		229, 3, 6, 3, 0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228,
		1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 230, 1, 0,
		0, 0, 233, 234, 5, 28, 0, 0, 234, 27, 1, 0, 0, 0, 235, 236, 5, 5, 0, 0,
		236, 237, 3, 68, 34, 0, 237, 242, 5, 12, 0, 0, 238, 241, 3, 2, 1, 0, 239,
		241, 3, 6, 3, 0, 240, 238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241, 244,
		1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 29, 1, 0,
		0, 0, 244, 242, 1, 0, 0, 0, 245, 246, 5, 6, 0, 0, 246, 251, 5, 12, 0, 0,
		247, 250, 3, 2, 1, 0, 248, 250, 3, 6, 3, 0, 249, 247, 1, 0, 0, 0, 249,
		248, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 31, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 5, 35,
		0, 0, 255, 256, 5, 26, 0, 0, 256, 257, 3, 34, 17, 0, 257, 258, 5, 10, 0,
		0, 258, 259, 3, 68, 34, 0, 259, 260, 5, 27, 0, 0, 260, 261, 3, 54, 27,
		0, 261, 33, 1, 0, 0, 0, 262, 264, 5, 13, 0, 0, 263, 262, 1, 0, 0, 0, 263,
		264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 271, 5, 70, 0, 0, 266, 267,
		5, 26, 0, 0, 267, 268, 3, 36, 18, 0, 268, 269, 5, 27, 0, 0, 269, 271, 1,
		0, 0, 0, 270, 263, 1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 271, 35, 1, 0, 0,
		0, 272, 275, 5, 70, 0, 0, 273, 274, 5, 39, 0, 0, 274, 276, 5, 70, 0, 0,
		275, 273, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277,
		278, 1, 0, 0, 0, 278, 37, 1, 0, 0, 0, 279, 291, 5, 13, 0, 0, 280, 292,
		5, 70, 0, 0, 281, 282, 5, 26, 0, 0, 282, 287, 5, 70, 0, 0, 283, 284, 5,
		39, 0, 0, 284, 286, 5, 70, 0, 0, 285, 283, 1, 0, 0, 0, 286, 289, 1, 0,
		0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0,
		289, 287, 1, 0, 0, 0, 290, 292, 5, 27, 0, 0, 291, 280, 1, 0, 0, 0, 291,
		281, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 5, 14, 0, 0, 294, 295,
		3, 68, 34, 0, 295, 39, 1, 0, 0, 0, 296, 297, 3, 68, 34, 0, 297, 298, 5,
		14, 0, 0, 298, 299, 3, 68, 34, 0, 299, 41, 1, 0, 0, 0, 300, 301, 5, 70,
		0, 0, 301, 302, 5, 29, 0, 0, 302, 303, 3, 68, 34, 0, 303, 304, 5, 30, 0,
		0, 304, 305, 5, 14, 0, 0, 305, 306, 3, 68, 34, 0, 306, 43, 1, 0, 0, 0,
		307, 308, 3, 68, 34, 0, 308, 309, 7, 1, 0, 0, 309, 310, 3, 68, 34, 0, 310,
		45, 1, 0, 0, 0, 311, 312, 5, 64, 0, 0, 312, 313, 3, 68, 34, 0, 313, 47,
		1, 0, 0, 0, 314, 315, 5, 31, 0, 0, 315, 319, 5, 26, 0, 0, 316, 317, 3,
		60, 30, 0, 317, 318, 5, 38, 0, 0, 318, 320, 1, 0, 0, 0, 319, 316, 1, 0,
		0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322, 3, 68, 34,
		0, 322, 323, 5, 27, 0, 0, 323, 332, 3, 54, 27, 0, 324, 325, 5, 33, 0, 0,
		325, 326, 5, 26, 0, 0, 326, 327, 3, 68, 34, 0, 327, 328, 5, 27, 0, 0, 328,
		329, 3, 54, 27, 0, 329, 331, 1, 0, 0, 0, 330, 324, 1, 0, 0, 0, 331, 334,
		1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 337, 1, 0,
		0, 0, 334, 332, 1, 0, 0, 0, 335, 336, 5, 32, 0, 0, 336, 338, 3, 54, 27,
		0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 49, 1, 0, 0, 0, 339,
		340, 5, 34, 0, 0, 340, 341, 5, 26, 0, 0, 341, 342, 3, 68, 34, 0, 342, 343,
		5, 27, 0, 0, 343, 344, 3, 54, 27, 0, 344, 51, 1, 0, 0, 0, 345, 346, 5,
		35, 0, 0, 346, 347, 5, 26, 0, 0, 347, 348, 3, 62, 31, 0, 348, 349, 5, 38,
		0, 0, 349, 350, 3, 68, 34, 0, 350, 351, 5, 38, 0, 0, 351, 352, 3, 64, 32,
		0, 352, 353, 5, 27, 0, 0, 353, 354, 3, 54, 27, 0, 354, 53, 1, 0, 0, 0,
		355, 361, 5, 25, 0, 0, 356, 360, 3, 2, 1, 0, 357, 360, 3, 8, 4, 0, 358,
		360, 3, 6, 3, 0, 359, 356, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 358,
		1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0,
		0, 0, 362, 364, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 364, 365, 5, 28, 0, 0,
		365, 55, 1, 0, 0, 0, 366, 367, 5, 3, 0, 0, 367, 368, 5, 70, 0, 0, 368,
		372, 5, 25, 0, 0, 369, 371, 3, 6, 3, 0, 370, 369, 1, 0, 0, 0, 371, 374,
		1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 394, 1, 0,
		0, 0, 374, 372, 1, 0, 0, 0, 375, 385, 3, 58, 29, 0, 376, 378, 3, 6, 3,
		0, 377, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379,
		380, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 3, 58, 29, 0, 382, 384,
		1, 0, 0, 0, 383, 377, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385, 383, 1, 0,
		0, 0, 385, 386, 1, 0, 0, 0, 386, 391, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0,
		388, 390, 3, 6, 3, 0, 389, 388, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391,
		389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391,
		1, 0, 0, 0, 394, 375, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 396, 1, 0,
		0, 0, 396, 397, 5, 28, 0, 0, 397, 57, 1, 0, 0, 0, 398, 399, 5, 70, 0, 0,
		399, 59, 1, 0, 0, 0, 400, 403, 3, 38, 19, 0, 401, 403, 3, 40, 20, 0, 402,
		400, 1, 0, 0, 0, 402, 401, 1, 0, 0, 0, 403, 61, 1, 0, 0, 0, 404, 407, 3,
		38, 19, 0, 405, 407, 3, 40, 20, 0, 406, 404, 1, 0, 0, 0, 406, 405, 1, 0,
		0, 0, 407, 63, 1, 0, 0, 0, 408, 411, 3, 40, 20, 0, 409, 411, 3, 66, 33,
		0, 410, 408, 1, 0, 0, 0, 410, 409, 1, 0, 0, 0, 411, 65, 1, 0, 0, 0, 412,
		413, 3, 68, 34, 0, 413, 414, 7, 2, 0, 0, 414, 67, 1, 0, 0, 0, 415, 416,
		6, 34, -1, 0, 416, 417, 5, 70, 0, 0, 417, 426, 5, 26, 0, 0, 418, 423, 3,
		68, 34, 0, 419, 420, 5, 39, 0, 0, 420, 422, 3, 68, 34, 0, 421, 419, 1,
		0, 0, 0, 422, 425, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0,
		0, 424, 427, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 418, 1, 0, 0, 0, 426,
		427, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 486, 5, 27, 0, 0, 429, 430,
		7, 3, 0, 0, 430, 486, 3, 68, 34, 25, 431, 486, 3, 76, 38, 0, 432, 486,
		3, 70, 35, 0, 433, 442, 5, 25, 0, 0, 434, 439, 3, 72, 36, 0, 435, 436,
		5, 39, 0, 0, 436, 438, 3, 72, 36, 0, 437, 435, 1, 0, 0, 0, 438, 441, 1,
		0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 443, 1, 0, 0,
		0, 441, 439, 1, 0, 0, 0, 442, 434, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		444, 1, 0, 0, 0, 444, 486, 5, 28, 0, 0, 445, 454, 5, 61, 0, 0, 446, 451,
		5, 70, 0, 0, 447, 448, 5, 39, 0, 0, 448, 450, 5, 70, 0, 0, 449, 447, 1,
		0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0,
		0, 452, 455, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 446, 1, 0, 0, 0, 454,
		455, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457, 5, 12, 0, 0, 457, 486,
		3, 68, 34, 9, 458, 459, 5, 26, 0, 0, 459, 460, 3, 68, 34, 0, 460, 461,
		5, 39, 0, 0, 461, 466, 3, 68, 34, 0, 462, 463, 5, 39, 0, 0, 463, 465, 3,
		68, 34, 0, 464, 462, 1, 0, 0, 0, 465, 468, 1, 0, 0, 0, 466, 464, 1, 0,
		0, 0, 466, 467, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0,
		469, 471, 5, 39, 0, 0, 470, 469, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471,
		472, 1, 0, 0, 0, 472, 473, 5, 27, 0, 0, 473, 486, 1, 0, 0, 0, 474, 475,
		5, 26, 0, 0, 475, 486, 5, 27, 0, 0, 476, 486, 5, 70, 0, 0, 477, 486, 5,
		71, 0, 0, 478, 486, 7, 4, 0, 0, 479, 486, 5, 72, 0, 0, 480, 486, 5, 1,
		0, 0, 481, 482, 5, 26, 0, 0, 482, 483, 3, 68, 34, 0, 483, 484, 5, 27, 0,
		0, 484, 486, 1, 0, 0, 0, 485, 415, 1, 0, 0, 0, 485, 429, 1, 0, 0, 0, 485,
		431, 1, 0, 0, 0, 485, 432, 1, 0, 0, 0, 485, 433, 1, 0, 0, 0, 485, 445,
		1, 0, 0, 0, 485, 458, 1, 0, 0, 0, 485, 474, 1, 0, 0, 0, 485, 476, 1, 0,
		0, 0, 485, 477, 1, 0, 0, 0, 485, 478, 1, 0, 0, 0, 485, 479, 1, 0, 0, 0,
		485, 480, 1, 0, 0, 0, 485, 481, 1, 0, 0, 0, 486, 568, 1, 0, 0, 0, 487,
		488, 10, 24, 0, 0, 488, 489, 5, 43, 0, 0, 489, 567, 3, 68, 34, 25, 490,
		491, 10, 23, 0, 0, 491, 492, 7, 5, 0, 0, 492, 567, 3, 68, 34, 24, 493,
		494, 10, 22, 0, 0, 494, 495, 7, 6, 0, 0, 495, 567, 3, 68, 34, 23, 496,
		497, 10, 21, 0, 0, 497, 498, 7, 7, 0, 0, 498, 567, 3, 68, 34, 22, 499,
		500, 10, 20, 0, 0, 500, 501, 5, 50, 0, 0, 501, 567, 3, 68, 34, 21, 502,
		503, 10, 19, 0, 0, 503, 504, 5, 52, 0, 0, 504, 567, 3, 68, 34, 20, 505,
		506, 10, 18, 0, 0, 506, 507, 5, 51, 0, 0, 507, 567, 3, 68, 34, 19, 508,
		509, 10, 17, 0, 0, 509, 510, 7, 8, 0, 0, 510, 567, 3, 68, 34, 18, 511,
		512, 10, 16, 0, 0, 512, 513, 3, 74, 37, 0, 513, 514, 3, 68, 34, 17, 514,
		567, 1, 0, 0, 0, 515, 516, 10, 15, 0, 0, 516, 517, 5, 67, 0, 0, 517, 567,
		3, 68, 34, 16, 518, 519, 10, 14, 0, 0, 519, 520, 5, 68, 0, 0, 520, 567,
		3, 68, 34, 15, 521, 522, 10, 13, 0, 0, 522, 523, 5, 31, 0, 0, 523, 524,
		3, 68, 34, 0, 524, 525, 5, 32, 0, 0, 525, 526, 3, 68, 34, 14, 526, 567,
		1, 0, 0, 0, 527, 528, 10, 30, 0, 0, 528, 529, 5, 29, 0, 0, 529, 530, 3,
		68, 34, 0, 530, 531, 5, 30, 0, 0, 531, 567, 1, 0, 0, 0, 532, 533, 10, 29,
		0, 0, 533, 535, 5, 29, 0, 0, 534, 536, 3, 68, 34, 0, 535, 534, 1, 0, 0,
		0, 535, 536, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537, 539, 5, 12, 0, 0, 538,
		540, 3, 68, 34, 0, 539, 538, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540, 545,
		1, 0, 0, 0, 541, 543, 5, 12, 0, 0, 542, 544, 3, 68, 34, 0, 543, 542, 1,
		0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 546, 1, 0, 0, 0, 545, 541, 1, 0, 0,
		0, 545, 546, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0, 547, 567, 5, 30, 0, 0, 548,
		549, 10, 28, 0, 0, 549, 550, 5, 11, 0, 0, 550, 551, 5, 70, 0, 0, 551, 560,
		5, 26, 0, 0, 552, 557, 3, 68, 34, 0, 553, 554, 5, 39, 0, 0, 554, 556, 3,
		68, 34, 0, 555, 553, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0,
		0, 0, 557, 558, 1, 0, 0, 0, 558, 561, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0,
		560, 552, 1, 0, 0, 0, 560, 561, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562,
		567, 5, 27, 0, 0, 563, 564, 10, 27, 0, 0, 564, 565, 5, 11, 0, 0, 565, 567,
		5, 70, 0, 0, 566, 487, 1, 0, 0, 0, 566, 490, 1, 0, 0, 0, 566, 493, 1, 0,
		0, 0, 566, 496, 1, 0, 0, 0, 566, 499, 1, 0, 0, 0, 566, 502, 1, 0, 0, 0,
		566, 505, 1, 0, 0, 0, 566, 508, 1, 0, 0, 0, 566, 511, 1, 0, 0, 0, 566,
		515, 1, 0, 0, 0, 566, 518, 1, 0, 0, 0, 566, 521, 1, 0, 0, 0, 566, 527,
		1, 0, 0, 0, 566, 532, 1, 0, 0, 0, 566, 548, 1, 0, 0, 0, 566, 563, 1, 0,
		0, 0, 567, 570, 1, 0, 0, 0, 568, 566, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0,
		569, 69, 1, 0, 0, 0, 570, 568, 1, 0, 0, 0, 571, 572, 5, 70, 0, 0, 572,
		581, 5, 25, 0, 0, 573, 578, 3, 72, 36, 0, 574, 575, 5, 39, 0, 0, 575, 577,
		3, 72, 36, 0, 576, 574, 1, 0, 0, 0, 577, 580, 1, 0, 0, 0, 578, 576, 1,
		0, 0, 0, 578, 579, 1, 0, 0, 0, 579, 582, 1, 0, 0, 0, 580, 578, 1, 0, 0,
		0, 581, 573, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583,
		584, 5, 28, 0, 0, 584, 71, 1, 0, 0, 0, 585, 586, 3, 68, 34, 0, 586, 587,
		5, 12, 0, 0, 587, 588, 3, 68, 34, 0, 588, 73, 1, 0, 0, 0, 589, 593, 5,
		10, 0, 0, 590, 591, 5, 69, 0, 0, 591, 593, 5, 10, 0, 0, 592, 589, 1, 0,
		0, 0, 592, 590, 1, 0, 0, 0, 593, 75, 1, 0, 0, 0, 594, 603, 5, 29, 0, 0,
		595, 600, 3, 68, 34, 0, 596, 597, 5, 39, 0, 0, 597, 599, 3, 68, 34, 0,
		598, 596, 1, 0, 0, 0, 599, 602, 1, 0, 0, 0, 600, 598, 1, 0, 0, 0, 600,
		601, 1, 0, 0, 0, 601, 604, 1, 0, 0, 0, 602, 600, 1, 0, 0, 0, 603, 595,
		1, 0, 0, 0, 603, 604, 1, 0, 0, 0, 604, 605, 1, 0, 0, 0, 605, 619, 5, 30,
		0, 0, 606, 607, 5, 29, 0, 0, 607, 608, 3, 68, 34, 0, 608, 609, 5, 35, 0,
		0, 609, 610, 5, 70, 0, 0, 610, 611, 5, 10, 0, 0, 611, 614, 3, 68, 34, 0,
		612, 613, 5, 31, 0, 0, 613, 615, 3, 68, 34, 0, 614, 612, 1, 0, 0, 0, 614,
		615, 1, 0, 0, 0, 615, 616, 1, 0, 0, 0, 616, 617, 5, 30, 0, 0, 617, 619,
		1, 0, 0, 0, 618, 594, 1, 0, 0, 0, 618, 606, 1, 0, 0, 0, 619, 77, 1, 0,
		0, 0, 65, 81, 83, 108, 111, 115, 121, 130, 133, 149, 156, 162, 168, 171,
		182, 185, 215, 221, 225, 230, 240, 242, 249, 251, 263, 270, 277, 287, 291,
		319, 332, 337, 359, 361, 372, 379, 385, 391, 394, 402, 406, 410, 423, 426,
		439, 442, 451, 454, 466, 470, 485, 535, 539, 543, 545, 557, 560, 566, 568,
		578, 581, 592, 600, 603, 614, 618,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF              = antlr.TokenEOF
	GrammarParserT__0             = 1
	GrammarParserINTERFACE        = 2
	GrammarParserSTRUCT           = 3
	GrammarParserSWITCH           = 4
	GrammarParserCASE             = 5
	GrammarParserDEFAULT          = 6
	GrammarParserTRY              = 7
	GrammarParserCATCH            = 8
	GrammarParserTHROW            = 9
	GrammarParserIN               = 10
	GrammarParserDOT              = 11
	GrammarParserCOLON            = 12
	GrammarParserVAR              = 13
	GrammarParserEQUALS           = 14
	GrammarParserPLUS             = 15
	GrammarParserMINUS            = 16
	GrammarParserSTAR             = 17
	GrammarParserSLASH            = 18
	GrammarParserLESS             = 19
	GrammarParserGREATER          = 20
	GrammarParserEQUALEQUAL       = 21
	GrammarParserBANGEQUAL        = 22
	GrammarParserLESSEQUAL        = 23
	GrammarParserGREATEREQUAL     = 24
	GrammarParserLBRACE           = 25
	GrammarParserLPAREN           = 26
	GrammarParserRPAREN           = 27
	GrammarParserRBRACE           = 28
	GrammarParserLBRACKET         = 29
	GrammarParserRBRACKET         = 30
	GrammarParserIF               = 31
	GrammarParserELSE             = 32
	GrammarParserELIF             = 33
	GrammarParserWHILE            = 34
	GrammarParserFOR              = 35
	GrammarParserBREAK            = 36
	GrammarParserCONTINUE         = 37
	GrammarParserSEMICOLON        = 38
	GrammarParserCOMMA            = 39
	GrammarParserINC              = 40
	GrammarParserDEC              = 41
	GrammarParserMODULO           = 42
	GrammarParserEXPONENTIAL      = 43
	GrammarParserPLUSEQUAL        = 44
	GrammarParserMINUSEQUAL       = 45
	GrammarParserSTAREQUAL        = 46
	GrammarParserSLASHEQUAL       = 47
	GrammarParserMODEQUAL         = 48
	GrammarParserEXPONENTIALEQUAL = 49
	GrammarParserBITAND           = 50
	GrammarParserBITOR            = 51
	GrammarParserBITXOR           = 52
	GrammarParserBITLSHIFT        = 53
	GrammarParserBITRSHIFT        = 54
	GrammarParserBITNOT           = 55
	GrammarParserBITANDEQUAL      = 56
	GrammarParserBITOREQUAL       = 57
	GrammarParserBITXOREQAUL      = 58
	GrammarParserBITLSHIFTEQUAL   = 59
	GrammarParserBITRSHIFTEQUAL   = 60
	GrammarParserLAMBDA           = 61
	GrammarParserFUNC             = 62
	GrammarParserRETURN           = 63
	GrammarParserPRINT            = 64
	GrammarParserTRUE             = 65
	GrammarParserFALSE            = 66
	GrammarParserAND              = 67
	GrammarParserOR               = 68
	GrammarParserNOT              = 69
	GrammarParserIDENTIFIER       = 70
	GrammarParserNUMBER           = 71
	GrammarParserSTRING           = 72
	GrammarParserLINE_COMMENT     = 73
	GrammarParserBLOCK_COMMENT    = 74
	GrammarParserNL               = 75
	GrammarParserWS               = 76
)

// GrammarParser rules.
const (
	GrammarParserRULE_program            = 0
	GrammarParserRULE_statement          = 1
	GrammarParserRULE_returnStmt         = 2
	GrammarParserRULE_terminator         = 3
	GrammarParserRULE_funcStmt           = 4
	GrammarParserRULE_receiver           = 5
	GrammarParserRULE_interfaceStmt      = 6
	GrammarParserRULE_method             = 7
	GrammarParserRULE_exprStmt           = 8
	GrammarParserRULE_tryCatchStmt       = 9
	GrammarParserRULE_throwStmt          = 10
	GrammarParserRULE_breakStmt          = 11
	GrammarParserRULE_continueStmt       = 12
	GrammarParserRULE_switchStmt         = 13
	GrammarParserRULE_caseBlock          = 14
	GrammarParserRULE_defaultBlock       = 15
	GrammarParserRULE_forInStmt          = 16
	GrammarParserRULE_loopTarget         = 17
	GrammarParserRULE_identifierList     = 18
	GrammarParserRULE_varDecl            = 19
	GrammarParserRULE_assignStmt         = 20
	GrammarParserRULE_arrayAssignStmt    = 21
	GrammarParserRULE_compoundAssignStmt = 22
	GrammarParserRULE_printStmt          = 23
	GrammarParserRULE_ifStmt             = 24
	GrammarParserRULE_whileStmt          = 25
	GrammarParserRULE_forStmt            = 26
	GrammarParserRULE_blockStmt          = 27
	GrammarParserRULE_structStmt         = 28
	GrammarParserRULE_structField        = 29
	GrammarParserRULE_ifInit             = 30
	GrammarParserRULE_forInit            = 31
	GrammarParserRULE_forPost            = 32
	GrammarParserRULE_postfixStmt        = 33
	GrammarParserRULE_expr               = 34
	GrammarParserRULE_structLiteral      = 35
	GrammarParserRULE_mapEntry           = 36
	GrammarParserRULE_membershipOp       = 37
	GrammarParserRULE_arrayLit           = 38
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFuncStmt() []IFuncStmtContext
	FuncStmt(i int) IFuncStmtContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) AllFuncStmt() []IFuncStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncStmtContext); ok {
			len++
		}
	}

	tst := make([]IFuncStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncStmtContext); ok {
			tst[i] = t.(IFuncStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncStmt(i int) IFuncStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncStmtContext)
}

func (s *ProgramContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2269813676833693026) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2535) != 0) {
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserT__0, GrammarParserINTERFACE, GrammarParserSTRUCT, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
			{
				p.SetState(78)
				p.Statement()
			}

		case GrammarParserFUNC:
			{
				p.SetState(79)
				p.FuncStmt()
			}

		case GrammarParserSEMICOLON, GrammarParserNL:
			{
				p.SetState(80)
				p.Terminator()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext
	StructStmt() IStructStmtContext
	InterfaceStmt() IInterfaceStmtContext
	ArrayAssignStmt() IArrayAssignStmtContext
	CompoundAssignStmt() ICompoundAssignStmtContext
	ExprStmt() IExprStmtContext
	PrintStmt() IPrintStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	SwitchStmt() ISwitchStmtContext
	ForInStmt() IForInStmtContext
	ForStmt() IForStmtContext
	BlockStmt() IBlockStmtContext
	PostfixStmt() IPostfixStmtContext
	TryCatchStmt() ITryCatchStmtContext
	ThrowStmt() IThrowStmtContext
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	Terminator() ITerminatorContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StatementContext) StructStmt() IStructStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructStmtContext)
}

func (s *StatementContext) InterfaceStmt() IInterfaceStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceStmtContext)
}

func (s *StatementContext) ArrayAssignStmt() IArrayAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayAssignStmtContext)
}

func (s *StatementContext) CompoundAssignStmt() ICompoundAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundAssignStmtContext)
}

func (s *StatementContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StatementContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StatementContext) ForInStmt() IForInStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *StatementContext) PostfixStmt() IPostfixStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixStmtContext)
}

func (s *StatementContext) TryCatchStmt() ITryCatchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryCatchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryCatchStmtContext)
}

func (s *StatementContext) ThrowStmt() IThrowStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThrowStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThrowStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) Terminator() ITerminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(88)
			p.VarDecl()
		}

	case 2:
		{
			p.SetState(89)
			p.AssignStmt()
		}

	case 3:
		{
			p.SetState(90)
			p.StructStmt()
		}

	case 4:
		{
			p.SetState(91)
			p.InterfaceStmt()
		}

	case 5:
		{
			p.SetState(92)
			p.ArrayAssignStmt()
		}

	case 6:
		{
			p.SetState(93)
			p.CompoundAssignStmt()
		}

	case 7:
		{
			p.SetState(94)
			p.ExprStmt()
		}

	case 8:
		{
			p.SetState(95)
			p.PrintStmt()
		}

	case 9:
		{
			p.SetState(96)
			p.IfStmt()
		}

	case 10:
		{
			p.SetState(97)
			p.WhileStmt()
		}

	case 11:
		{
			p.SetState(98)
			p.SwitchStmt()
		}

	case 12:
		{
			p.SetState(99)
			p.ForInStmt()
		}

	case 13:
		{
			p.SetState(100)
			p.ForStmt()
		}

	case 14:
		{
			p.SetState(101)
			p.BlockStmt()
		}

	case 15:
		{
			p.SetState(102)
			p.PostfixStmt()
		}

	case 16:
		{
			p.SetState(103)
			p.TryCatchStmt()
		}

	case 17:
		{
			p.SetState(104)
			p.ThrowStmt()
		}

	case 18:
		{
			p.SetState(105)
			p.ReturnStmt()
		}

	case 19:
		{
			p.SetState(106)
			p.BreakStmt()
		}

	case 20:
		{
			p.SetState(107)
			p.ContinueStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Terminator()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(GrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	NL() antlr.TerminalNode

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_terminator
	return p
}

func InitEmptyTerminatorContext(p *TerminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_terminator
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *TerminatorContext) NL() antlr.TerminalNode {
	return s.GetToken(GrammarParserNL, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (s *TerminatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTerminator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Terminator() (localctx ITerminatorContext) {
	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_terminator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserSEMICOLON || _la == GrammarParserNL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncStmtContext is an interface to support dynamic dispatch.
type IFuncStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	BlockStmt() IBlockStmtContext
	Receiver() IReceiverContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncStmtContext differentiates from other interfaces.
	IsFuncStmtContext()
}

type FuncStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStmtContext() *FuncStmtContext {
	var p = new(FuncStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcStmt
	return p
}

func InitEmptyFuncStmtContext(p *FuncStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcStmt
}

func (*FuncStmtContext) IsFuncStmtContext() {}

func NewFuncStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStmtContext {
	var p = new(FuncStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_funcStmt

	return p
}

func (s *FuncStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GrammarParserFUNC, 0)
}

func (s *FuncStmtContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *FuncStmtContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *FuncStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FuncStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FuncStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *FuncStmtContext) Receiver() IReceiverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReceiverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReceiverContext)
}

func (s *FuncStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *FuncStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *FuncStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFuncStmt(s)
	}
}

func (s *FuncStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFuncStmt(s)
	}
}

func (s *FuncStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFuncStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) FuncStmt() (localctx IFuncStmtContext) {
	localctx = NewFuncStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_funcStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserLPAREN {
		{
			p.SetState(120)
			p.Receiver()
		}

	}
	{
		p.SetState(123)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserIDENTIFIER {
		{
			p.SetState(125)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserCOMMA {
			{
				p.SetState(126)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(127)
				p.Match(GrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(135)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.BlockStmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReceiverContext is an interface to support dynamic dispatch.
type IReceiverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetStructType returns the structType token.
	GetStructType() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetStructType sets the structType token.
	SetStructType(antlr.Token)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsReceiverContext differentiates from other interfaces.
	IsReceiverContext()
}

type ReceiverContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	id         antlr.Token
	structType antlr.Token
}

func NewEmptyReceiverContext() *ReceiverContext {
	var p = new(ReceiverContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_receiver
	return p
}

func InitEmptyReceiverContext(p *ReceiverContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_receiver
}

func (*ReceiverContext) IsReceiverContext() {}

func NewReceiverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiverContext {
	var p = new(ReceiverContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_receiver

	return p
}

func (s *ReceiverContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiverContext) GetId() antlr.Token { return s.id }

func (s *ReceiverContext) GetStructType() antlr.Token { return s.structType }

func (s *ReceiverContext) SetId(v antlr.Token) { s.id = v }

func (s *ReceiverContext) SetStructType(v antlr.Token) { s.structType = v }

func (s *ReceiverContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ReceiverContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ReceiverContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *ReceiverContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *ReceiverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterReceiver(s)
	}
}

func (s *ReceiverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitReceiver(s)
	}
}

func (s *ReceiverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitReceiver(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Receiver() (localctx IReceiverContext) {
	localctx = NewReceiverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_receiver)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*ReceiverContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*ReceiverContext).structType = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceStmtContext is an interface to support dynamic dispatch.
type IInterfaceStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext
	AllMethod() []IMethodContext
	Method(i int) IMethodContext

	// IsInterfaceStmtContext differentiates from other interfaces.
	IsInterfaceStmtContext()
}

type InterfaceStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceStmtContext() *InterfaceStmtContext {
	var p = new(InterfaceStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_interfaceStmt
	return p
}

func InitEmptyInterfaceStmtContext(p *InterfaceStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_interfaceStmt
}

func (*InterfaceStmtContext) IsInterfaceStmtContext() {}

func NewInterfaceStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceStmtContext {
	var p = new(InterfaceStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_interfaceStmt

	return p
}

func (s *InterfaceStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceStmtContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserINTERFACE, 0)
}

func (s *InterfaceStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *InterfaceStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *InterfaceStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *InterfaceStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *InterfaceStmtContext) AllMethod() []IMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodContext); ok {
			len++
		}
	}

	tst := make([]IMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodContext); ok {
			tst[i] = t.(IMethodContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceStmtContext) Method(i int) IMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *InterfaceStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterInterfaceStmt(s)
	}
}

func (s *InterfaceStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitInterfaceStmt(s)
	}
}

func (s *InterfaceStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitInterfaceStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) InterfaceStmt() (localctx IInterfaceStmtContext) {
	localctx = NewInterfaceStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_interfaceStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(GrammarParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
		{
			p.SetState(146)
			p.Terminator()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserIDENTIFIER {
		{
			p.SetState(152)
			p.Method()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
					{
						p.SetState(153)
						p.Terminator()
					}

					p.SetState(156)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(158)
					p.Method()
				}

			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
			{
				p.SetState(165)
				p.Terminator()
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(173)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *MethodContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *MethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *MethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *MethodContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *MethodContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *MethodContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (s *MethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_method)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserIDENTIFIER {
		{
			p.SetState(177)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserCOMMA {
			{
				p.SetState(178)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(179)
				p.Match(GrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(187)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITryCatchStmtContext is an interface to support dynamic dispatch.
type ITryCatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTryBody returns the tryBody rule contexts.
	GetTryBody() IBlockStmtContext

	// GetCatchBody returns the catchBody rule contexts.
	GetCatchBody() IBlockStmtContext

	// SetTryBody sets the tryBody rule contexts.
	SetTryBody(IBlockStmtContext)

	// SetCatchBody sets the catchBody rule contexts.
	SetCatchBody(IBlockStmtContext)

	// Getter signatures
	TRY() antlr.TerminalNode
	CATCH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllBlockStmt() []IBlockStmtContext
	BlockStmt(i int) IBlockStmtContext

	// IsTryCatchStmtContext differentiates from other interfaces.
	IsTryCatchStmtContext()
}

type TryCatchStmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tryBody   IBlockStmtContext
	catchBody IBlockStmtContext
}

func NewEmptyTryCatchStmtContext() *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tryCatchStmt
	return p
}

func InitEmptyTryCatchStmtContext(p *TryCatchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tryCatchStmt
}

func (*TryCatchStmtContext) IsTryCatchStmtContext() {}

func NewTryCatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_tryCatchStmt

	return p
}

func (s *TryCatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TryCatchStmtContext) GetTryBody() IBlockStmtContext { return s.tryBody }

func (s *TryCatchStmtContext) GetCatchBody() IBlockStmtContext { return s.catchBody }

func (s *TryCatchStmtContext) SetTryBody(v IBlockStmtContext) { s.tryBody = v }

func (s *TryCatchStmtContext) SetCatchBody(v IBlockStmtContext) { s.catchBody = v }

func (s *TryCatchStmtContext) TRY() antlr.TerminalNode {
	return s.GetToken(GrammarParserTRY, 0)
}

func (s *TryCatchStmtContext) CATCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserCATCH, 0)
}

func (s *TryCatchStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *TryCatchStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *TryCatchStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *TryCatchStmtContext) AllBlockStmt() []IBlockStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStmtContext); ok {
			len++
		}
	}

	tst := make([]IBlockStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStmtContext); ok {
			tst[i] = t.(IBlockStmtContext)
			i++
		}
	}

	return tst
}

func (s *TryCatchStmtContext) BlockStmt(i int) IBlockStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *TryCatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryCatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryCatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTryCatchStmt(s)
	}
}

func (s *TryCatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTryCatchStmt(s)
	}
}

func (s *TryCatchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTryCatchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) TryCatchStmt() (localctx ITryCatchStmtContext) {
	localctx = NewTryCatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_tryCatchStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GrammarParserTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)

		var _x = p.BlockStmt()

		localctx.(*TryCatchStmtContext).tryBody = _x
	}
	{
		p.SetState(194)
		p.Match(GrammarParserCATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)

		var _x = p.BlockStmt()

		localctx.(*TryCatchStmtContext).catchBody = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IThrowStmtContext is an interface to support dynamic dispatch.
type IThrowStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THROW() antlr.TerminalNode
	Expr() IExprContext

	// IsThrowStmtContext differentiates from other interfaces.
	IsThrowStmtContext()
}

type ThrowStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowStmtContext() *ThrowStmtContext {
	var p = new(ThrowStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_throwStmt
	return p
}

func InitEmptyThrowStmtContext(p *ThrowStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_throwStmt
}

func (*ThrowStmtContext) IsThrowStmtContext() {}

func NewThrowStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStmtContext {
	var p = new(ThrowStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_throwStmt

	return p
}

func (s *ThrowStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowStmtContext) THROW() antlr.TerminalNode {
	return s.GetToken(GrammarParserTHROW, 0)
}

func (s *ThrowStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ThrowStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterThrowStmt(s)
	}
}

func (s *ThrowStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitThrowStmt(s)
	}
}

func (s *ThrowStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitThrowStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ThrowStmt() (localctx IThrowStmtContext) {
	localctx = NewThrowStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_throwStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(GrammarParserTHROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(GrammarParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(GrammarParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *SwitchStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *SwitchStmtContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SwitchStmtContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_switchStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.expr(0)
	}
	{
		p.SetState(210)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(212)
				p.Terminator()
			}

		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserCASE {
		{
			p.SetState(218)
			p.CaseBlock()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserDEFAULT {
		{
			p.SetState(224)
			p.DefaultBlock()
		}

	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
		{
			p.SetState(227)
			p.Terminator()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *CaseBlockContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *CaseBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseBlockContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlockContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_caseBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.expr(0)
	}
	{
		p.SetState(237)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case GrammarParserT__0, GrammarParserINTERFACE, GrammarParserSTRUCT, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
				{
					p.SetState(238)
					p.Statement()
				}

			case GrammarParserSEMICOLON, GrammarParserNL:
				{
					p.SetState(239)
					p.Terminator()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *DefaultBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *DefaultBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DefaultBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultBlockContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *DefaultBlockContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_defaultBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case GrammarParserT__0, GrammarParserINTERFACE, GrammarParserSTRUCT, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
				{
					p.SetState(247)
					p.Statement()
				}

			case GrammarParserSEMICOLON, GrammarParserNL:
				{
					p.SetState(248)
					p.Terminator()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInStmtContext is an interface to support dynamic dispatch.
type IForInStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	LoopTarget() ILoopTargetContext
	IN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	BlockStmt() IBlockStmtContext

	// IsForInStmtContext differentiates from other interfaces.
	IsForInStmtContext()
}

type ForInStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IBlockStmtContext
}

func NewEmptyForInStmtContext() *ForInStmtContext {
	var p = new(ForInStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInStmt
	return p
}

func InitEmptyForInStmtContext(p *ForInStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInStmt
}

func (*ForInStmtContext) IsForInStmtContext() {}

func NewForInStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInStmtContext {
	var p = new(ForInStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forInStmt

	return p
}

func (s *ForInStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *ForInStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *ForInStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForInStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ForInStmtContext) LoopTarget() ILoopTargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopTargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopTargetContext)
}

func (s *ForInStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ForInStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForInStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ForInStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *ForInStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForInStmt(s)
	}
}

func (s *ForInStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForInStmt(s)
	}
}

func (s *ForInStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForInStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForInStmt() (localctx IForInStmtContext) {
	localctx = NewForInStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_forInStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.LoopTarget()
	}
	{
		p.SetState(257)
		p.Match(GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.expr(0)
	}
	{
		p.SetState(259)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)

		var _x = p.BlockStmt()

		localctx.(*ForInStmtContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopTargetContext is an interface to support dynamic dispatch.
type ILoopTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoopTargetContext differentiates from other interfaces.
	IsLoopTargetContext()
}

type LoopTargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopTargetContext() *LoopTargetContext {
	var p = new(LoopTargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_loopTarget
	return p
}

func InitEmptyLoopTargetContext(p *LoopTargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_loopTarget
}

func (*LoopTargetContext) IsLoopTargetContext() {}

func NewLoopTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopTargetContext {
	var p = new(LoopTargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_loopTarget

	return p
}

func (s *LoopTargetContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopTargetContext) CopyAll(ctx *LoopTargetContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleLoopVarContext struct {
	LoopTargetContext
	id antlr.Token
}

func NewSingleLoopVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleLoopVarContext {
	var p = new(SingleLoopVarContext)

	InitEmptyLoopTargetContext(&p.LoopTargetContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopTargetContext))

	return p
}

func (s *SingleLoopVarContext) GetId() antlr.Token { return s.id }

func (s *SingleLoopVarContext) SetId(v antlr.Token) { s.id = v }

func (s *SingleLoopVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleLoopVarContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *SingleLoopVarContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *SingleLoopVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSingleLoopVar(s)
	}
}

func (s *SingleLoopVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSingleLoopVar(s)
	}
}

func (s *SingleLoopVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSingleLoopVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleUnpackLoopVarContext struct {
	LoopTargetContext
}

func NewTupleUnpackLoopVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleUnpackLoopVarContext {
	var p = new(TupleUnpackLoopVarContext)

	InitEmptyLoopTargetContext(&p.LoopTargetContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopTargetContext))

	return p
}

func (s *TupleUnpackLoopVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleUnpackLoopVarContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *TupleUnpackLoopVarContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *TupleUnpackLoopVarContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *TupleUnpackLoopVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTupleUnpackLoopVar(s)
	}
}

func (s *TupleUnpackLoopVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTupleUnpackLoopVar(s)
	}
}

func (s *TupleUnpackLoopVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTupleUnpackLoopVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) LoopTarget() (localctx ILoopTargetContext) {
	localctx = NewLoopTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_loopTarget)
	var _la int

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR, GrammarParserIDENTIFIER:
		localctx = NewSingleLoopVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserVAR {
			{
				p.SetState(262)
				p.Match(GrammarParserVAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(265)

			var _m = p.Match(GrammarParserIDENTIFIER)

			localctx.(*SingleLoopVarContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserLPAREN:
		localctx = NewTupleUnpackLoopVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.IdentifierList()
		}
		{
			p.SetState(268)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) GetId() antlr.Token { return s.id }

func (s *IdentifierListContext) SetId(v antlr.Token) { s.id = v }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*IdentifierListContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserCOMMA {
		{
			p.SetState(273)
			p.Match(GrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expr() IExprContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *VarDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *VarDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *VarDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *VarDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *VarDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *VarDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *VarDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_varDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(GrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserIDENTIFIER:
		{
			p.SetState(280)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserLPAREN:
		{
			p.SetState(281)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserCOMMA {
			{
				p.SetState(283)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(284)
				p.Match(GrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(289)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(290)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(293)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQUALS() antlr.TerminalNode

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignStmt
	return p
}

func InitEmptyAssignStmtContext(p *AssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assignStmt
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AssignStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.expr(0)
	}
	{
		p.SetState(297)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayAssignStmtContext is an interface to support dynamic dispatch.
type IArrayAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RBRACKET() antlr.TerminalNode
	EQUALS() antlr.TerminalNode

	// IsArrayAssignStmtContext differentiates from other interfaces.
	IsArrayAssignStmtContext()
}

type ArrayAssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayAssignStmtContext() *ArrayAssignStmtContext {
	var p = new(ArrayAssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt
	return p
}

func InitEmptyArrayAssignStmtContext(p *ArrayAssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt
}

func (*ArrayAssignStmtContext) IsArrayAssignStmtContext() {}

func NewArrayAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayAssignStmtContext {
	var p = new(ArrayAssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_arrayAssignStmt

	return p
}

func (s *ArrayAssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayAssignStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *ArrayAssignStmtContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ArrayAssignStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayAssignStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayAssignStmtContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ArrayAssignStmtContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALS, 0)
}

func (s *ArrayAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayAssignStmt(s)
	}
}

func (s *ArrayAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayAssignStmt(s)
	}
}

func (s *ArrayAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ArrayAssignStmt() (localctx IArrayAssignStmtContext) {
	localctx = NewArrayAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_arrayAssignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(GrammarParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.expr(0)
	}
	{
		p.SetState(303)
		p.Match(GrammarParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(GrammarParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompoundAssignStmtContext is an interface to support dynamic dispatch.
type ICompoundAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PLUSEQUAL() antlr.TerminalNode
	MINUSEQUAL() antlr.TerminalNode
	STAREQUAL() antlr.TerminalNode
	SLASHEQUAL() antlr.TerminalNode
	MODEQUAL() antlr.TerminalNode
	EXPONENTIALEQUAL() antlr.TerminalNode
	BITANDEQUAL() antlr.TerminalNode
	BITOREQUAL() antlr.TerminalNode
	BITXOREQAUL() antlr.TerminalNode
	BITLSHIFTEQUAL() antlr.TerminalNode
	BITRSHIFTEQUAL() antlr.TerminalNode

	// IsCompoundAssignStmtContext differentiates from other interfaces.
	IsCompoundAssignStmtContext()
}

type CompoundAssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCompoundAssignStmtContext() *CompoundAssignStmtContext {
	var p = new(CompoundAssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt
	return p
}

func InitEmptyCompoundAssignStmtContext(p *CompoundAssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt
}

func (*CompoundAssignStmtContext) IsCompoundAssignStmtContext() {}

func NewCompoundAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundAssignStmtContext {
	var p = new(CompoundAssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_compoundAssignStmt

	return p
}

func (s *CompoundAssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundAssignStmtContext) GetOp() antlr.Token { return s.op }

func (s *CompoundAssignStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompoundAssignStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompoundAssignStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompoundAssignStmtContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUSEQUAL, 0)
}

func (s *CompoundAssignStmtContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUSEQUAL, 0)
}

func (s *CompoundAssignStmtContext) STAREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTAREQUAL, 0)
}

func (s *CompoundAssignStmtContext) SLASHEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserSLASHEQUAL, 0)
}

func (s *CompoundAssignStmtContext) MODEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMODEQUAL, 0)
}

func (s *CompoundAssignStmtContext) EXPONENTIALEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEXPONENTIALEQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITANDEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITANDEQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITOREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITOREQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITXOREQAUL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITXOREQAUL, 0)
}

func (s *CompoundAssignStmtContext) BITLSHIFTEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITLSHIFTEQUAL, 0)
}

func (s *CompoundAssignStmtContext) BITRSHIFTEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITRSHIFTEQUAL, 0)
}

func (s *CompoundAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundAssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundAssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterCompoundAssignStmt(s)
	}
}

func (s *CompoundAssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitCompoundAssignStmt(s)
	}
}

func (s *CompoundAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCompoundAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) CompoundAssignStmt() (localctx ICompoundAssignStmtContext) {
	localctx = NewCompoundAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrammarParserRULE_compoundAssignStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.expr(0)
	}
	{
		p.SetState(308)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CompoundAssignStmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2234893722896564224) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CompoundAssignStmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(309)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	Expr() IExprContext

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPRINT, 0)
}

func (s *PrintStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrammarParserRULE_printStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(GrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IIfInitContext

	// GetThenBranch returns the thenBranch rule contexts.
	GetThenBranch() IBlockStmtContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_blockStmt returns the _blockStmt rule contexts.
	Get_blockStmt() IBlockStmtContext

	// GetElseBranch returns the elseBranch rule contexts.
	GetElseBranch() IBlockStmtContext

	// SetInit sets the init rule contexts.
	SetInit(IIfInitContext)

	// SetThenBranch sets the thenBranch rule contexts.
	SetThenBranch(IBlockStmtContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_blockStmt sets the _blockStmt rule contexts.
	Set_blockStmt(IBlockStmtContext)

	// SetElseBranch sets the elseBranch rule contexts.
	SetElseBranch(IBlockStmtContext)

	// GetElifCond returns the elifCond rule context list.
	GetElifCond() []IExprContext

	// GetElifBranch returns the elifBranch rule context list.
	GetElifBranch() []IBlockStmtContext

	// SetElifCond sets the elifCond rule context list.
	SetElifCond([]IExprContext)

	// SetElifBranch sets the elifBranch rule context list.
	SetElifBranch([]IBlockStmtContext)

	// Getter signatures
	IF() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllBlockStmt() []IBlockStmtContext
	BlockStmt(i int) IBlockStmtContext
	SEMICOLON() antlr.TerminalNode
	AllELIF() []antlr.TerminalNode
	ELIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	IfInit() IIfInitContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	init       IIfInitContext
	thenBranch IBlockStmtContext
	_expr      IExprContext
	elifCond   []IExprContext
	_blockStmt IBlockStmtContext
	elifBranch []IBlockStmtContext
	elseBranch IBlockStmtContext
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) GetInit() IIfInitContext { return s.init }

func (s *IfStmtContext) GetThenBranch() IBlockStmtContext { return s.thenBranch }

func (s *IfStmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfStmtContext) Get_blockStmt() IBlockStmtContext { return s._blockStmt }

func (s *IfStmtContext) GetElseBranch() IBlockStmtContext { return s.elseBranch }

func (s *IfStmtContext) SetInit(v IIfInitContext) { s.init = v }

func (s *IfStmtContext) SetThenBranch(v IBlockStmtContext) { s.thenBranch = v }

func (s *IfStmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfStmtContext) Set_blockStmt(v IBlockStmtContext) { s._blockStmt = v }

func (s *IfStmtContext) SetElseBranch(v IBlockStmtContext) { s.elseBranch = v }

func (s *IfStmtContext) GetElifCond() []IExprContext { return s.elifCond }

func (s *IfStmtContext) GetElifBranch() []IBlockStmtContext { return s.elifBranch }

func (s *IfStmtContext) SetElifCond(v []IExprContext) { s.elifCond = v }

func (s *IfStmtContext) SetElifBranch(v []IBlockStmtContext) { s.elifBranch = v }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfStmtContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserLPAREN)
}

func (s *IfStmtContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, i)
}

func (s *IfStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserRPAREN)
}

func (s *IfStmtContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, i)
}

func (s *IfStmtContext) AllBlockStmt() []IBlockStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStmtContext); ok {
			len++
		}
	}

	tst := make([]IBlockStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStmtContext); ok {
			tst[i] = t.(IBlockStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) BlockStmt(i int) IBlockStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *IfStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *IfStmtContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserELIF)
}

func (s *IfStmtContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserELIF, i)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfStmtContext) IfInit() IIfInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfInitContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GrammarParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(GrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(316)

			var _x = p.IfInit()

			localctx.(*IfStmtContext).init = _x
		}
		{
			p.SetState(317)
			p.Match(GrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(321)
		p.expr(0)
	}
	{
		p.SetState(322)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)

		var _x = p.BlockStmt()

		localctx.(*IfStmtContext).thenBranch = _x
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserELIF {
		{
			p.SetState(324)
			p.Match(GrammarParserELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)

			var _x = p.expr(0)

			localctx.(*IfStmtContext)._expr = _x
		}
		localctx.(*IfStmtContext).elifCond = append(localctx.(*IfStmtContext).elifCond, localctx.(*IfStmtContext)._expr)
		{
			p.SetState(327)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)

			var _x = p.BlockStmt()

			localctx.(*IfStmtContext)._blockStmt = _x
		}
		localctx.(*IfStmtContext).elifBranch = append(localctx.(*IfStmtContext).elifBranch, localctx.(*IfStmtContext)._blockStmt)

		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserELSE {
		{
			p.SetState(335)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)

			var _x = p.BlockStmt()

			localctx.(*IfStmtContext).elseBranch = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	BlockStmt() IBlockStmtContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IBlockStmtContext
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *WhileStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *WhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *WhileStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.expr(0)
	}
	{
		p.SetState(342)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)

		var _x = p.BlockStmt()

		localctx.(*WhileStmtContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IForInitContext

	// GetCond returns the cond rule contexts.
	GetCond() IExprContext

	// GetPost returns the post rule contexts.
	GetPost() IForPostContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockStmtContext

	// SetInit sets the init rule contexts.
	SetInit(IForInitContext)

	// SetCond sets the cond rule contexts.
	SetCond(IExprContext)

	// SetPost sets the post rule contexts.
	SetPost(IForPostContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockStmtContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ForInit() IForInitContext
	Expr() IExprContext
	ForPost() IForPostContext
	BlockStmt() IBlockStmtContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   IForInitContext
	cond   IExprContext
	post   IForPostContext
	body   IBlockStmtContext
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) GetInit() IForInitContext { return s.init }

func (s *ForStmtContext) GetCond() IExprContext { return s.cond }

func (s *ForStmtContext) GetPost() IForPostContext { return s.post }

func (s *ForStmtContext) GetBody() IBlockStmtContext { return s.body }

func (s *ForStmtContext) SetInit(v IForInitContext) { s.init = v }

func (s *ForStmtContext) SetCond(v IExprContext) { s.cond = v }

func (s *ForStmtContext) SetPost(v IForPostContext) { s.post = v }

func (s *ForStmtContext) SetBody(v IBlockStmtContext) { s.body = v }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, i)
}

func (s *ForStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ForStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) ForPost() IForPostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForPostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForPostContext)
}

func (s *ForStmtContext) BlockStmt() IBlockStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStmtContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GrammarParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)

		var _x = p.ForInit()

		localctx.(*ForStmtContext).init = _x
	}
	{
		p.SetState(348)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)

		var _x = p.expr(0)

		localctx.(*ForStmtContext).cond = _x
	}
	{
		p.SetState(350)
		p.Match(GrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)

		var _x = p.ForPost()

		localctx.(*ForStmtContext).post = _x
	}
	{
		p.SetState(352)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)

		var _x = p.BlockStmt()

		localctx.(*ForStmtContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStmtContext is an interface to support dynamic dispatch.
type IBlockStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllFuncStmt() []IFuncStmtContext
	FuncStmt(i int) IFuncStmtContext
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext

	// IsBlockStmtContext differentiates from other interfaces.
	IsBlockStmtContext()
}

type BlockStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStmtContext() *BlockStmtContext {
	var p = new(BlockStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockStmt
	return p
}

func InitEmptyBlockStmtContext(p *BlockStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_blockStmt
}

func (*BlockStmtContext) IsBlockStmtContext() {}

func NewBlockStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStmtContext {
	var p = new(BlockStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_blockStmt

	return p
}

func (s *BlockStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *BlockStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *BlockStmtContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStmtContext) AllFuncStmt() []IFuncStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncStmtContext); ok {
			len++
		}
	}

	tst := make([]IFuncStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncStmtContext); ok {
			tst[i] = t.(IFuncStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) FuncStmt(i int) IFuncStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncStmtContext)
}

func (s *BlockStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *BlockStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) BlockStmt() (localctx IBlockStmtContext) {
	localctx = NewBlockStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GrammarParserRULE_blockStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2269813676833693026) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2535) != 0) {
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserT__0, GrammarParserINTERFACE, GrammarParserSTRUCT, GrammarParserSWITCH, GrammarParserTRY, GrammarParserTHROW, GrammarParserVAR, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserIF, GrammarParserWHILE, GrammarParserFOR, GrammarParserBREAK, GrammarParserCONTINUE, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserRETURN, GrammarParserPRINT, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
			{
				p.SetState(356)
				p.Statement()
			}

		case GrammarParserFUNC:
			{
				p.SetState(357)
				p.FuncStmt()
			}

		case GrammarParserSEMICOLON, GrammarParserNL:
			{
				p.SetState(358)
				p.Terminator()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(364)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructStmtContext is an interface to support dynamic dispatch.
type IStructStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllTerminator() []ITerminatorContext
	Terminator(i int) ITerminatorContext
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext

	// IsStructStmtContext differentiates from other interfaces.
	IsStructStmtContext()
}

type StructStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyStructStmtContext() *StructStmtContext {
	var p = new(StructStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structStmt
	return p
}

func InitEmptyStructStmtContext(p *StructStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structStmt
}

func (*StructStmtContext) IsStructStmtContext() {}

func NewStructStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStmtContext {
	var p = new(StructStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structStmt

	return p
}

func (s *StructStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStmtContext) GetId() antlr.Token { return s.id }

func (s *StructStmtContext) SetId(v antlr.Token) { s.id = v }

func (s *StructStmtContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRUCT, 0)
}

func (s *StructStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *StructStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *StructStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *StructStmtContext) AllTerminator() []ITerminatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITerminatorContext); ok {
			len++
		}
	}

	tst := make([]ITerminatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITerminatorContext); ok {
			tst[i] = t.(ITerminatorContext)
			i++
		}
	}

	return tst
}

func (s *StructStmtContext) Terminator(i int) ITerminatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *StructStmtContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructStmtContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldContext)
}

func (s *StructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStructStmt(s)
	}
}

func (s *StructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStructStmt(s)
	}
}

func (s *StructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructStmt() (localctx IStructStmtContext) {
	localctx = NewStructStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrammarParserRULE_structStmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(GrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*StructStmtContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
		{
			p.SetState(369)
			p.Terminator()
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserIDENTIFIER {
		{
			p.SetState(375)
			p.StructField()
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(377)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
					{
						p.SetState(376)
						p.Terminator()
					}

					p.SetState(379)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(381)
					p.StructField()
				}

			}
			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserSEMICOLON || _la == GrammarParserNL {
			{
				p.SetState(388)
				p.Terminator()
			}

			p.SetState(393)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(396)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrammarParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(GrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfInitContext is an interface to support dynamic dispatch.
type IIfInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext

	// IsIfInitContext differentiates from other interfaces.
	IsIfInitContext()
}

type IfInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfInitContext() *IfInitContext {
	var p = new(IfInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifInit
	return p
}

func InitEmptyIfInitContext(p *IfInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifInit
}

func (*IfInitContext) IsIfInitContext() {}

func NewIfInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfInitContext {
	var p = new(IfInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifInit

	return p
}

func (s *IfInitContext) GetParser() antlr.Parser { return s.parser }

func (s *IfInitContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *IfInitContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *IfInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIfInit(s)
	}
}

func (s *IfInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIfInit(s)
	}
}

func (s *IfInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IfInit() (localctx IIfInitContext) {
	localctx = NewIfInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GrammarParserRULE_ifInit)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.VarDecl()
		}

	case GrammarParserT__0, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.AssignStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	AssignStmt() IAssignStmtContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInit
	return p
}

func InitEmptyForInitContext(p *ForInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forInit
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ForInitContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GrammarParserRULE_forInit)
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.VarDecl()
		}

	case GrammarParserT__0, GrammarParserMINUS, GrammarParserLBRACE, GrammarParserLPAREN, GrammarParserLBRACKET, GrammarParserBITNOT, GrammarParserLAMBDA, GrammarParserTRUE, GrammarParserFALSE, GrammarParserNOT, GrammarParserIDENTIFIER, GrammarParserNUMBER, GrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.AssignStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForPostContext is an interface to support dynamic dispatch.
type IForPostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignStmt() IAssignStmtContext
	PostfixStmt() IPostfixStmtContext

	// IsForPostContext differentiates from other interfaces.
	IsForPostContext()
}

type ForPostContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForPostContext() *ForPostContext {
	var p = new(ForPostContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forPost
	return p
}

func InitEmptyForPostContext(p *ForPostContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forPost
}

func (*ForPostContext) IsForPostContext() {}

func NewForPostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForPostContext {
	var p = new(ForPostContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forPost

	return p
}

func (s *ForPostContext) GetParser() antlr.Parser { return s.parser }

func (s *ForPostContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ForPostContext) PostfixStmt() IPostfixStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixStmtContext)
}

func (s *ForPostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForPostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForPostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterForPost(s)
	}
}

func (s *ForPostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitForPost(s)
	}
}

func (s *ForPostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForPost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForPost() (localctx IForPostContext) {
	localctx = NewForPostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GrammarParserRULE_forPost)
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.PostfixStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixStmtContext is an interface to support dynamic dispatch.
type IPostfixStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	INC() antlr.TerminalNode
	DEC() antlr.TerminalNode

	// IsPostfixStmtContext differentiates from other interfaces.
	IsPostfixStmtContext()
}

type PostfixStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPostfixStmtContext() *PostfixStmtContext {
	var p = new(PostfixStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_postfixStmt
	return p
}

func InitEmptyPostfixStmtContext(p *PostfixStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_postfixStmt
}

func (*PostfixStmtContext) IsPostfixStmtContext() {}

func NewPostfixStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixStmtContext {
	var p = new(PostfixStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_postfixStmt

	return p
}

func (s *PostfixStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixStmtContext) GetOp() antlr.Token { return s.op }

func (s *PostfixStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *PostfixStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PostfixStmtContext) INC() antlr.TerminalNode {
	return s.GetToken(GrammarParserINC, 0)
}

func (s *PostfixStmtContext) DEC() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEC, 0)
}

func (s *PostfixStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPostfixStmt(s)
	}
}

func (s *PostfixStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPostfixStmt(s)
	}
}

func (s *PostfixStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPostfixStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) PostfixStmt() (localctx IPostfixStmtContext) {
	localctx = NewPostfixStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GrammarParserRULE_postfixStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.expr(0)
	}
	{
		p.SetState(413)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PostfixStmtContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserINC || _la == GrammarParserDEC) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PostfixStmtContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullContext struct {
	ExprContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExprContext
	op antlr.Token
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(GrammarParserOR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentialContext struct {
	ExprContext
	op antlr.Token
}

func NewExponentialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentialContext {
	var p = new(ExponentialContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExponentialContext) GetOp() antlr.Token { return s.op }

func (s *ExponentialContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExponentialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentialContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExponentialContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExponentialContext) EXPONENTIAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEXPONENTIAL, 0)
}

func (s *ExponentialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExponential(s)
	}
}

func (s *ExponentialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExponential(s)
	}
}

func (s *ExponentialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExponential(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayIndexContext struct {
	ExprContext
}

func NewArrayIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayIndexContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayIndexContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ArrayIndexContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (s *ArrayIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type MethodCallContext struct {
	ExprContext
}

func NewMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallContext {
	var p = new(MethodCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MethodCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MethodCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOT, 0)
}

func (s *MethodCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *MethodCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *MethodCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *MethodCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *MethodCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructContext struct {
	ExprContext
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) StructLiteral() IStructLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructLiteralContext)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivModContext) STAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTAR, 0)
}

func (s *MulDivModContext) SLASH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSLASH, 0)
}

func (s *MulDivModContext) MODULO() antlr.TerminalNode {
	return s.GetToken(GrammarParserMODULO, 0)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

func (s *MulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	ExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GrammarParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonContext struct {
	ExprContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(GrammarParserLESS, 0)
}

func (s *ComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(GrammarParserGREATER, 0)
}

func (s *ComparisonContext) EQUALEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUALEQUAL, 0)
}

func (s *ComparisonContext) BANGEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBANGEQUAL, 0)
}

func (s *ComparisonContext) LESSEQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserLESSEQUAL, 0)
}

func (s *ComparisonContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserGREATEREQUAL, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleLiteralContext struct {
	ExprContext
}

func NewTupleLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleLiteralContext {
	var p = new(TupleLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TupleLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleLiteralContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *TupleLiteralContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TupleLiteralContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TupleLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *TupleLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *TupleLiteralContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *TupleLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTupleLiteral(s)
	}
}

func (s *TupleLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTupleLiteral(s)
	}
}

func (s *TupleLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTupleLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type MembershipContext struct {
	ExprContext
	leftExpr  IExprContext
	op        IMembershipOpContext
	rightExpr IExprContext
}

func NewMembershipContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembershipContext {
	var p = new(MembershipContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MembershipContext) GetLeftExpr() IExprContext { return s.leftExpr }

func (s *MembershipContext) GetOp() IMembershipOpContext { return s.op }

func (s *MembershipContext) GetRightExpr() IExprContext { return s.rightExpr }

func (s *MembershipContext) SetLeftExpr(v IExprContext) { s.leftExpr = v }

func (s *MembershipContext) SetOp(v IMembershipOpContext) { s.op = v }

func (s *MembershipContext) SetRightExpr(v IExprContext) { s.rightExpr = v }

func (s *MembershipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MembershipContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MembershipContext) MembershipOp() IMembershipOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMembershipOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMembershipOpContext)
}

func (s *MembershipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMembership(s)
	}
}

func (s *MembershipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMembership(s)
	}
}

func (s *MembershipContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMembership(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	ExprContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyTupleLiteralContext struct {
	ExprContext
}

func NewEmptyTupleLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyTupleLiteralContext {
	var p = new(EmptyTupleLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EmptyTupleLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyTupleLiteralContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *EmptyTupleLiteralContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *EmptyTupleLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterEmptyTupleLiteral(s)
	}
}

func (s *EmptyTupleLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitEmptyTupleLiteral(s)
	}
}

func (s *EmptyTupleLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitEmptyTupleLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ExprContext
	val antlr.Token
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanContext) GetVal() antlr.Token { return s.val }

func (s *BooleanContext) SetVal(v antlr.Token) { s.val = v }

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserFALSE, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	ExprContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *ParenthesesContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitOrContext struct {
	ExprContext
	op antlr.Token
}

func NewBitOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrContext {
	var p = new(BitOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitOrContext) GetOp() antlr.Token { return s.op }

func (s *BitOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BitOrContext) BITOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITOR, 0)
}

func (s *BitOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitOr(s)
	}
}

func (s *BitOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitOr(s)
	}
}

func (s *BitOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceIndexContext struct {
	ExprContext
	startOpt IExprContext
	endOpt   IExprContext
	stepOpt  IExprContext
}

func NewSliceIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceIndexContext {
	var p = new(SliceIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SliceIndexContext) GetStartOpt() IExprContext { return s.startOpt }

func (s *SliceIndexContext) GetEndOpt() IExprContext { return s.endOpt }

func (s *SliceIndexContext) GetStepOpt() IExprContext { return s.stepOpt }

func (s *SliceIndexContext) SetStartOpt(v IExprContext) { s.startOpt = v }

func (s *SliceIndexContext) SetEndOpt(v IExprContext) { s.endOpt = v }

func (s *SliceIndexContext) SetStepOpt(v IExprContext) { s.stepOpt = v }

func (s *SliceIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceIndexContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceIndexContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceIndexContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *SliceIndexContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOLON)
}

func (s *SliceIndexContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, i)
}

func (s *SliceIndexContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *SliceIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSliceIndex(s)
	}
}

func (s *SliceIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSliceIndex(s)
	}
}

func (s *SliceIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSliceIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUS, 0)
}

func (s *AddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitShiftContext struct {
	ExprContext
	op antlr.Token
}

func NewBitShiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftContext {
	var p = new(BitShiftContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitShiftContext) GetOp() antlr.Token { return s.op }

func (s *BitShiftContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitShiftContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BitShiftContext) BITLSHIFT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITLSHIFT, 0)
}

func (s *BitShiftContext) BITRSHIFT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITRSHIFT, 0)
}

func (s *BitShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitShift(s)
	}
}

func (s *BitShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitShift(s)
	}
}

func (s *BitShiftContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitShift(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExprContext struct {
	ExprContext
}

func NewLambdaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExprContext {
	var p = new(LambdaExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) LAMBDA() antlr.TerminalNode {
	return s.GetToken(GrammarParserLAMBDA, 0)
}

func (s *LambdaExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *LambdaExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LambdaExprContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserIDENTIFIER)
}

func (s *LambdaExprContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, i)
}

func (s *LambdaExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *LambdaExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitLambdaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLiteralContext struct {
	ExprContext
}

func NewArrayLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ArrayLit() IArrayLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLitContext)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	ExprContext
	op antlr.Token
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *UnaryContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITNOT, 0)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryOpContext struct {
	ExprContext
	trueExpr  IExprContext
	condExpr  IExprContext
	falseExpr IExprContext
}

func NewTernaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryOpContext {
	var p = new(TernaryOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TernaryOpContext) GetTrueExpr() IExprContext { return s.trueExpr }

func (s *TernaryOpContext) GetCondExpr() IExprContext { return s.condExpr }

func (s *TernaryOpContext) GetFalseExpr() IExprContext { return s.falseExpr }

func (s *TernaryOpContext) SetTrueExpr(v IExprContext) { s.trueExpr = v }

func (s *TernaryOpContext) SetCondExpr(v IExprContext) { s.condExpr = v }

func (s *TernaryOpContext) SetFalseExpr(v IExprContext) { s.falseExpr = v }

func (s *TernaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryOpContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *TernaryOpContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *TernaryOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TernaryOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TernaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTernaryOp(s)
	}
}

func (s *TernaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTernaryOp(s)
	}
}

func (s *TernaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTernaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapLiteralContext struct {
	ExprContext
}

func NewMapLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLiteralContext {
	var p = new(MapLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *MapLiteralContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntryContext)
}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitXorContext struct {
	ExprContext
	op antlr.Token
}

func NewBitXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorContext {
	var p = new(BitXorContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitXorContext) GetOp() antlr.Token { return s.op }

func (s *BitXorContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitXorContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BitXorContext) BITXOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITXOR, 0)
}

func (s *BitXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitXor(s)
	}
}

func (s *BitXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitXor(s)
	}
}

func (s *BitXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	ExprContext
	op antlr.Token
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(GrammarParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitAndContext struct {
	ExprContext
	op antlr.Token
}

func NewBitAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndContext {
	var p = new(BitAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitAndContext) GetOp() antlr.Token { return s.op }

func (s *BitAndContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BitAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BitAndContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GrammarParserBITAND, 0)
}

func (s *BitAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBitAnd(s)
	}
}

func (s *BitAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBitAnd(s)
	}
}

func (s *BitAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldAccessContext struct {
	ExprContext
}

func NewFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessContext {
	var p = new(FieldAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOT, 0)
}

func (s *FieldAccessContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (s *FieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(416)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
			{
				p.SetState(418)
				p.expr(0)
			}
			p.SetState(423)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(419)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(420)
					p.expr(0)
				}

				p.SetState(425)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(428)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-16)) & ^0x3f) == 0 && ((int64(1)<<(_la-16))&9007749010554881) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(430)
			p.expr(25)
		}

	case 3:
		localctx = NewArrayLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(431)
			p.ArrayLit()
		}

	case 4:
		localctx = NewStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(432)
			p.StructLiteral()
		}

	case 5:
		localctx = NewMapLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(433)
			p.Match(GrammarParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
			{
				p.SetState(434)
				p.MapEntry()
			}
			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(435)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(436)
					p.MapEntry()
				}

				p.SetState(441)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(444)
			p.Match(GrammarParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(445)
			p.Match(GrammarParserLAMBDA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserIDENTIFIER {
			{
				p.SetState(446)
				p.Match(GrammarParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(451)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(447)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(448)
					p.Match(GrammarParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(453)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(456)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.expr(9)
		}

	case 7:
		localctx = NewTupleLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(458)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.expr(0)
		}
		{
			p.SetState(460)
			p.Match(GrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expr(0)
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(462)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(463)
					p.expr(0)
				}

			}
			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserCOMMA {
			{
				p.SetState(469)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(472)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewEmptyTupleLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(474)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(476)
			p.Match(GrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(477)
			p.Match(GrammarParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(478)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*BooleanContext).val = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserTRUE || _la == GrammarParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*BooleanContext).val = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 12:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(479)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(480)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(481)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.expr(0)
		}
		{
			p.SetState(483)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(566)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentialContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(487)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(488)

					var _m = p.Match(GrammarParserEXPONENTIAL)

					localctx.(*ExponentialContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(489)
					p.expr(25)
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(490)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(491)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046904320) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(492)
					p.expr(24)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(493)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(494)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserPLUS || _la == GrammarParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(495)
					p.expr(23)
				}

			case 4:
				localctx = NewBitShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(496)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(497)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BitShiftContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserBITLSHIFT || _la == GrammarParserBITRSHIFT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BitShiftContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(498)
					p.expr(22)
				}

			case 5:
				localctx = NewBitAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(499)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(500)

					var _m = p.Match(GrammarParserBITAND)

					localctx.(*BitAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(501)
					p.expr(21)
				}

			case 6:
				localctx = NewBitXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(502)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(503)

					var _m = p.Match(GrammarParserBITXOR)

					localctx.(*BitXorContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(504)
					p.expr(20)
				}

			case 7:
				localctx = NewBitOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(505)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(506)

					var _m = p.Match(GrammarParserBITOR)

					localctx.(*BitOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(507)
					p.expr(19)
				}

			case 8:
				localctx = NewComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(508)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(509)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33030144) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(510)
					p.expr(18)
				}

			case 9:
				localctx = NewMembershipContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MembershipContext).leftExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(512)

					var _x = p.MembershipOp()

					localctx.(*MembershipContext).op = _x
				}
				{
					p.SetState(513)

					var _x = p.expr(17)

					localctx.(*MembershipContext).rightExpr = _x
				}

			case 10:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(515)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(516)

					var _m = p.Match(GrammarParserAND)

					localctx.(*AndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(517)
					p.expr(16)
				}

			case 11:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(518)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(519)

					var _m = p.Match(GrammarParserOR)

					localctx.(*OrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(520)
					p.expr(15)
				}

			case 12:
				localctx = NewTernaryOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryOpContext).trueExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(522)
					p.Match(GrammarParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(523)

					var _x = p.expr(0)

					localctx.(*TernaryOpContext).condExpr = _x
				}
				{
					p.SetState(524)
					p.Match(GrammarParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(525)

					var _x = p.expr(14)

					localctx.(*TernaryOpContext).falseExpr = _x
				}

			case 13:
				localctx = NewArrayIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(527)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(528)
					p.Match(GrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(529)
					p.expr(0)
				}
				{
					p.SetState(530)
					p.Match(GrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 14:
				localctx = NewSliceIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(532)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(533)
					p.Match(GrammarParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(535)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
					{
						p.SetState(534)

						var _x = p.expr(0)

						localctx.(*SliceIndexContext).startOpt = _x
					}

				}
				{
					p.SetState(537)
					p.Match(GrammarParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(539)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
					{
						p.SetState(538)

						var _x = p.expr(0)

						localctx.(*SliceIndexContext).endOpt = _x
					}

				}
				p.SetState(545)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == GrammarParserCOLON {
					{
						p.SetState(541)
						p.Match(GrammarParserCOLON)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					p.SetState(543)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
						{
							p.SetState(542)

							var _x = p.expr(0)

							localctx.(*SliceIndexContext).stepOpt = _x
						}

					}

				}
				{
					p.SetState(547)
					p.Match(GrammarParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 15:
				localctx = NewMethodCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(548)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(549)
					p.Match(GrammarParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(550)
					p.Match(GrammarParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(551)
					p.Match(GrammarParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(560)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
					{
						p.SetState(552)
						p.expr(0)
					}
					p.SetState(557)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					for _la == GrammarParserCOMMA {
						{
							p.SetState(553)
							p.Match(GrammarParserCOMMA)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(554)
							p.expr(0)
						}

						p.SetState(559)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(562)
					p.Match(GrammarParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 16:
				localctx = NewFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(563)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(564)
					p.Match(GrammarParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(565)
					p.Match(GrammarParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructLiteralContext is an interface to support dynamic dispatch.
type IStructLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructLiteralContext differentiates from other interfaces.
	IsStructLiteralContext()
}

type StructLiteralContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	structName antlr.Token
}

func NewEmptyStructLiteralContext() *StructLiteralContext {
	var p = new(StructLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structLiteral
	return p
}

func InitEmptyStructLiteralContext(p *StructLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structLiteral
}

func (*StructLiteralContext) IsStructLiteralContext() {}

func NewStructLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructLiteralContext {
	var p = new(StructLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structLiteral

	return p
}

func (s *StructLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StructLiteralContext) GetStructName() antlr.Token { return s.structName }

func (s *StructLiteralContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *StructLiteralContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *StructLiteralContext) AllMapEntry() []IMapEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapEntryContext); ok {
			len++
		}
	}

	tst := make([]IMapEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapEntryContext); ok {
			tst[i] = t.(IMapEntryContext)
			i++
		}
	}

	return tst
}

func (s *StructLiteralContext) MapEntry(i int) IMapEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntryContext)
}

func (s *StructLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *StructLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStructLiteral(s)
	}
}

func (s *StructLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStructLiteral(s)
	}
}

func (s *StructLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructLiteral() (localctx IStructLiteralContext) {
	localctx = NewStructLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GrammarParserRULE_structLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)

		var _m = p.Match(GrammarParserIDENTIFIER)

		localctx.(*StructLiteralContext).structName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(572)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
		{
			p.SetState(573)
			p.MapEntry()
		}
		p.SetState(578)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserCOMMA {
			{
				p.SetState(574)
				p.Match(GrammarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(575)
				p.MapEntry()
			}

			p.SetState(580)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(583)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_mapEntry
	return p
}

func InitEmptyMapEntryContext(p *MapEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_mapEntry
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MapEntryContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMapEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) MapEntry() (localctx IMapEntryContext) {
	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GrammarParserRULE_mapEntry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.expr(0)
	}
	{
		p.SetState(586)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(587)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMembershipOpContext is an interface to support dynamic dispatch.
type IMembershipOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsMembershipOpContext differentiates from other interfaces.
	IsMembershipOpContext()
}

type MembershipOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembershipOpContext() *MembershipOpContext {
	var p = new(MembershipOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_membershipOp
	return p
}

func InitEmptyMembershipOpContext(p *MembershipOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_membershipOp
}

func (*MembershipOpContext) IsMembershipOpContext() {}

func NewMembershipOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembershipOpContext {
	var p = new(MembershipOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_membershipOp

	return p
}

func (s *MembershipOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MembershipOpContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *MembershipOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNOT, 0)
}

func (s *MembershipOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembershipOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMembershipOp(s)
	}
}

func (s *MembershipOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMembershipOp(s)
	}
}

func (s *MembershipOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMembershipOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) MembershipOp() (localctx IMembershipOpContext) {
	localctx = NewMembershipOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GrammarParserRULE_membershipOp)
	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(589)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(590)
			p.Match(GrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLitContext is an interface to support dynamic dispatch.
type IArrayLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayLitContext differentiates from other interfaces.
	IsArrayLitContext()
}

type ArrayLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLitContext() *ArrayLitContext {
	var p = new(ArrayLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayLit
	return p
}

func InitEmptyArrayLitContext(p *ArrayLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_arrayLit
}

func (*ArrayLitContext) IsArrayLitContext() {}

func NewArrayLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLitContext {
	var p = new(ArrayLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_arrayLit

	return p
}

func (s *ArrayLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLitContext) CopyAll(ctx *ArrayLitContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListComprehensionContext struct {
	ArrayLitContext
	transformExpr IExprContext
	id            antlr.Token
	srcExpr       IExprContext
	filterExpr    IExprContext
}

func NewListComprehensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListComprehensionContext {
	var p = new(ListComprehensionContext)

	InitEmptyArrayLitContext(&p.ArrayLitContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLitContext))

	return p
}

func (s *ListComprehensionContext) GetId() antlr.Token { return s.id }

func (s *ListComprehensionContext) SetId(v antlr.Token) { s.id = v }

func (s *ListComprehensionContext) GetTransformExpr() IExprContext { return s.transformExpr }

func (s *ListComprehensionContext) GetSrcExpr() IExprContext { return s.srcExpr }

func (s *ListComprehensionContext) GetFilterExpr() IExprContext { return s.filterExpr }

func (s *ListComprehensionContext) SetTransformExpr(v IExprContext) { s.transformExpr = v }

func (s *ListComprehensionContext) SetSrcExpr(v IExprContext) { s.srcExpr = v }

func (s *ListComprehensionContext) SetFilterExpr(v IExprContext) { s.filterExpr = v }

func (s *ListComprehensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListComprehensionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *ListComprehensionContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ListComprehensionContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ListComprehensionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *ListComprehensionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListComprehensionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListComprehensionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFIER, 0)
}

func (s *ListComprehensionContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *ListComprehensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListComprehension(s)
	}
}

func (s *ListComprehensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListComprehension(s)
	}
}

func (s *ListComprehensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitListComprehension(s)

	default:
		return t.VisitChildren(s)
	}
}

type StandardArrayContext struct {
	ArrayLitContext
}

func NewStandardArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StandardArrayContext {
	var p = new(StandardArrayContext)

	InitEmptyArrayLitContext(&p.ArrayLitContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLitContext))

	return p
}

func (s *StandardArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandardArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACKET, 0)
}

func (s *StandardArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACKET, 0)
}

func (s *StandardArrayContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StandardArrayContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StandardArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserCOMMA)
}

func (s *StandardArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMMA, i)
}

func (s *StandardArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStandardArray(s)
	}
}

func (s *StandardArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStandardArray(s)
	}
}

func (s *StandardArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStandardArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ArrayLit() (localctx IArrayLitContext) {
	localctx = NewArrayLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GrammarParserRULE_arrayLit)
	var _la int

	p.SetState(618)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStandardArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(594)
			p.Match(GrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2341871806870257666) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&243) != 0) {
			{
				p.SetState(595)
				p.expr(0)
			}
			p.SetState(600)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GrammarParserCOMMA {
				{
					p.SetState(596)
					p.Match(GrammarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(597)
					p.expr(0)
				}

				p.SetState(602)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(605)
			p.Match(GrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewListComprehensionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)
			p.Match(GrammarParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)

			var _x = p.expr(0)

			localctx.(*ListComprehensionContext).transformExpr = _x
		}
		{
			p.SetState(608)
			p.Match(GrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)

			var _m = p.Match(GrammarParserIDENTIFIER)

			localctx.(*ListComprehensionContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.Match(GrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(611)

			var _x = p.expr(0)

			localctx.(*ListComprehensionContext).srcExpr = _x
		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserIF {
			{
				p.SetState(612)
				p.Match(GrammarParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(613)

				var _x = p.expr(0)

				localctx.(*ListComprehensionContext).filterExpr = _x
			}

		}
		{
			p.SetState(616)
			p.Match(GrammarParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 34:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 27)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
