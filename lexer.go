package main

type TokenType int

const (
	Keyword      TokenType = iota // like int string or let
	Identifier                    // like variable names
	Char                          // Char variables inside " ' "
	String                        // String variables inside quotes
	Number                        // Number variable
	True                          // Boolean true
	False                         // Boolean false
	Equals                        // =
	Plus                          // +
	Minus                         // -
	Multiple                      // *
	Divide                        // /
	Mod                           // %
	Greater                       // >
	Lesser                        // <
	GreaterEqual                  // >=
	LesserEqual                   // <=
	LParen                        // (
	RParen                        // )
	LBrace                        // {
	RBrace                        // }
	LBracket                      // [
	RBracket                      // ]
	Comma                         //
	Semicolon                     // ;
	Comment                       // '//'
	EOF                           // End of File
)
