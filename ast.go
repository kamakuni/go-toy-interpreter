package main

type Expr struct {
	Span Span
	Node Expr_
}

type Expr_ int

const (
	Block Expr_ = iota
)

type Constant int

const (
	String Constant = iota
	Number
	Bool
)
