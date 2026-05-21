package main

//  var x = 10 + 20
// [ varToken, identifierToken(x), equalsToken(=), numberToken(10), PlusToken(+), numberToken(20) ]

type TokenType string

const (
	VarToken          TokenType = "VAR"
	IdentifierToken   TokenType = "IDENTIFIER"
	EqualsToken       TokenType = "EQUALS"
	NumberToken       TokenType = "NUMBER"
	PlusToken         TokenType = "PLUS"
	MinusToken        TokenType = "MINUS"
	StarToken         TokenType = "STAR"
	SlashToken        TokenType = "SLASH"
	PrintToken        TokenType = "PRINT"
	StringToken       TokenType = "STRING"
	LessToken         TokenType = "LESS"
	GreaterToken      TokenType = "GREATER"
	EqualEqualToken   TokenType = "EQUALEQUAL"
	BangEqualToken    TokenType = "BANGEQUAL"
	TrueToken         TokenType = "TRUE"
	FalseToken        TokenType = "FALSE"
	LessEqualToken    TokenType = "LESSEQUAL"
	GreaterEqualToken TokenType = "GREATEREQUAL"
	IfToken           TokenType = "IF"
	ElseToken         TokenType = "ELSE"
	LBraceToken       TokenType = "LBRACE"
	RBraceToken       TokenType = "RBRACE"
	WhileToken        TokenType = "WHILE"
	ForToken          TokenType = "FOR"
	SemiColonToken    TokenType = "SEMICOLON"
	FuncToken         TokenType = "FUNC"
	CommaToken        TokenType = "COMMA"
	LParenToken       TokenType = "LPAREN"
	RParenToken       TokenType = "RPAREN"
	ReturnToken       TokenType = "RETURN"
	ModuloToken       TokenType = "MODULO"
	ExponentialToken  TokenType = "EXPONENTIAL"
)

var keywords = map[string]TokenType{
	"var":    VarToken,
	"print":  PrintToken,
	"true":   TrueToken,
	"false":  FalseToken,
	"if":     IfToken,
	"else":   ElseToken,
	"while":  WhileToken,
	"for":    ForToken,
	"func":   FuncToken,
	"return": ReturnToken,
}

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

func NewToken(tokenType TokenType, value string, line int) Token {
	return Token{
		Type:  tokenType,
		Value: value,
		Line:  line,
	}
}

func Lex(input string) []Token {
	var tokens []Token
	var currentToken string
	currentLine := 1

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == '\n' {
			currentLine++
		}

		switch char {
		case '"':
			// for string literals, we read until the closing quote
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			var strValue string
			i++ // skip the opening quote
			// consume chars until we find the closing quote
			for i < len(input) && input[i] != '"' {
				strValue += string(input[i])
				i++
			}
			if i >= len(input) {
				panic("Unterminated string literal")
			}
			tokens = append(tokens, NewToken(StringToken, strValue, currentLine))

		case ' ', '\t', '\n':
			// for space, nl or tab, we finalize the current token and reset it
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// skip
			continue

		case '+':
			// for operators, we finalize the current token and add the operator token
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(PlusToken, string(char), currentLine))

		case '-':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(MinusToken, string(char), currentLine))

		case '*':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// look ahead to check if it's "**"
			if i+1 < len(input) && input[i+1] == '*' {
				tokens = append(tokens, NewToken(ExponentialToken, "**", currentLine))
				i++ // skip the next '*'
			} else {
				tokens = append(tokens, NewToken(StarToken, string(char), currentLine))
			}

		case '/':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(SlashToken, string(char), currentLine))

		case '=':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// check if "=="
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, NewToken(EqualEqualToken, "==", currentLine))
				i++ // skip the next '='
			} else {
				tokens = append(tokens, NewToken(EqualsToken, string(char), currentLine))
			}
		case '!':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// Look ahead to see if it's "!="
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, NewToken(BangEqualToken, "!=", currentLine))
				i++ // Skip the '='
			} else {
				panic("Unexpected character '!' without '='")
			}

		case '<':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// check if "<="
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, NewToken(LessEqualToken, "<=", currentLine))
				i++ // skip the next '='
			} else {
				tokens = append(tokens, NewToken(LessToken, string(char), currentLine))
			}

		case '>':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			// check if ">="
			if i+1 < len(input) && input[i+1] == '=' {
				tokens = append(tokens, NewToken(GreaterEqualToken, ">=", currentLine))
				i++ // skip the next '='
			} else {
				tokens = append(tokens, NewToken(GreaterToken, string(char), currentLine))
			}
		case '{':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(LBraceToken, string(char), currentLine))
		case '}':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(RBraceToken, string(char), currentLine))
		case ';':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(SemiColonToken, string(char), currentLine))
		case ',':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(CommaToken, string(char), currentLine))
		case '(':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(LParenToken, string(char), currentLine))

		case ')':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(RParenToken, string(char), currentLine))
		case '%':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken, currentLine))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(ModuloToken, string(char), currentLine))
		default:
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, classifyToken(currentToken, currentLine))
	}

	return tokens
}

// check if the token is a keyword, a number, or an identifier
func classifyToken(token string, line int) Token {
	if tokenType, ok := keywords[token]; ok {
		return NewToken(tokenType, token, line)
	} else if isNumber(token) {
		return NewToken(NumberToken, token, line)
	} else if isAlphabetic(token) {
		return NewToken(IdentifierToken, token, line)
	} else {
		panic("Unknown token: " + token)
	}
}

func isNumber(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isAlphabetic(s string) bool {
	for _, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return true
}
