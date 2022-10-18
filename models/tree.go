package models

type Node struct {
	board *Board
	left  *Node
	right *Node
}

type Tree struct{}
