package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	ASSSIGN   = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	DOUBLE_EQ = "=="
	NOT_EQ    = "!="

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNTION = "FUNCTION"
	LET     = "LET"
	TRUE    = "TRUE"
	FALSE   = "FALSE"
	IF      = "IF"
	ELSE    = "ELSE"
	RETURN  = "RETURN"
)

// creating a dictionary for keywords
var keywoards = map[string]TokenType{
	"fn":     FUNTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywoards[ident]; ok {
		return tok // is it a keyword
	}

	return IDENT // it just an identifier
}
