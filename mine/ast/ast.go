package ast

import "monkey/token"

// every node in our AST has to implement the Node interface
// so it needs a TokenLiteral() method that returns the literal value of the token
// TokenLiteral() is only used for debugging and testing
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// The root node of every AST our parser produces
// Every valid Monkey program is a series of statements
// Statements is just a slice of AST nodes that implement the Statement interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let statement: identifier, expression, token
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // let _foo_ = bar;
	Value Expression  // let foo = _getBar(5) + 10_;
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the `return` token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
