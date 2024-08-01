package parser

import (
	"go/token"
	"monkey/lexer"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}