package structs

type TokenType int

const (
	INT TokenType = iota
	OP
)

type Token struct {
	Kind  TokenType
	Value int
	Op    string
}
