package main

//  var x = 10 + 20
// [ varToken, identifierToken(x), equalsToken(=), numberToken(10), PlusToken(+), numberToken(20) ]

type TokenType string

const (
	VarToken        TokenType = "VAR"
	IdentifierToken TokenType = "IDENTIFIER"
	EqualsToken     TokenType = "EQUALS"
	NumberToken     TokenType = "NUMBER"
	PlusToken       TokenType = "PLUS"
)

// TODO: add keywords

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}

func Lex(input string) []Token {
	var tokens []Token
	var currentToken string

	for _, char := range input {
		switch char {
		case ' ', '\t', '\n':
			// for space, nl or tab, we finalize the current token and reset it
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			// skip
			continue
		case '+':
			// for operators, we finalize the current token and add the operator token
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(PlusToken, string(char)))
		case '=':
			if currentToken != "" {
				tokens = append(tokens, classifyToken(currentToken))
				currentToken = ""
			}
			tokens = append(tokens, NewToken(EqualsToken, string(char)))

		default:
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, classifyToken(currentToken))
	}

	return tokens
}

// check if the token is a keyword ("var"), a number, or an identifier
func classifyToken(token string) Token {
	switch token {
	case "var":
		return NewToken(VarToken, token)
	default:
		if isNumber(token) {
			return NewToken(NumberToken, token)
		}
		return NewToken(IdentifierToken, token)
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
