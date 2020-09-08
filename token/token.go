package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL 	= "ILLEGAL"
	EOF   		= "EOF"

	// Identifiers + literals
	IDENT 		= "IDENT"
	INT   		= "INT"

	// Operators
	ASSIGN		= "="
	PLUS		= "+"

	// Delimiters
	COMMA		= ","
	SEMICOLON	= ";"

	LPAREN		= "("
	RPAREN		= ")"

	LBRACE		= "{"
	RBRACE		= "}"

	// Keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
)

var keywords = map[string]Type{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}



