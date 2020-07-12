package datastructs

type CompareFun func(a interface{}, b interface{}) bool

type SortedElement struct {
	Key int
	
	Value interface{}
}
