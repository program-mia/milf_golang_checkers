package models

type Node struct {
	board  *Board
	left   *Node
	right  *Node
	rating float64
}

type Tree struct {
	root *Node
}
