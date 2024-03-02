package models;

type Node[T any] struct{
  Value T 
  Next *Node[T] 
}