package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Token representa un identificador reconocido
type Token struct {
	Type  string
	Value string
}

// Lexer analiza el texto de entrada y genera tokens
type Lexer struct {
	input string
	pos   int
}

// NewLexer crea un nuevo analizador léxico
func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

// NextToken analiza el siguiente token
func (l *Lexer) NextToken() *Token {
	if l.pos >= len(l.input) {
		return nil
	}

	char := rune(l.input[l.pos])

	// Verificar si es una letra
	if unicode.IsLetter(char) {
		start := l.pos
		for l.pos < len(l.input) && unicode.IsLetter(rune(l.input[l.pos])) {
			l.pos++
		}
		return &Token{Type: "IDENTIFICADOR", Value: l.input[start:l.pos]}
	}

	l.pos++
	return l.NextToken() // Ignorar caracteres no deseados y seguir buscando
}

func main() {
	// Solicitar entrada de texto al usuario
	fmt.Println("Ingrese el texto a analizar:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	lexer := NewLexer(texto)

	fmt.Println("Tokens  reconocidos:")
	for token := lexer.NextToken(); token != nil; token = lexer.NextToken() {
		fmt.Printf("Tipo: %s, Valor: %s\n", token.Type, token.Value)
	}
}
