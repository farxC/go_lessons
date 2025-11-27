package main

// import (
// 	"iter"
// )

// type Element[T any] struct {
// 	next *Element[T]
// 	val  T
// }

// type LinkedListType[T any] struct {
// 	head, tail *Element[T]
// }

// func (lst *LinkedListType[T]) push(v T) {
// 	if lst.tail == nil {
// 		lst.head = &Element[T]{val: v}
// 		lst.tail = lst.head
// 	} else {
// 		lst.tail.next = &Element[T]{val: v}
// 		lst.tail = lst.tail.next // Next element that will be pushed
// 	}
// }

// func (lst *LinkedListType[T]) All() iter.Seq[T] {
// 	return func(yield func(T) bool) {
// 		for e := lst.head; e != nil; e = e.next {
// 			if !yield(e.val) {
// 				return
// 			}
// 		}
// 	}
// }
