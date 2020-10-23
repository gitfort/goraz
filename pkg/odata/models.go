package odata

type OrderBy struct {
	Field string
	Sort  Sort
}

type Filter struct {
	Logical  Logic
	Field    string
	Operator Operate
	Value    interface{}
}
