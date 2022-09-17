package common

type A struct {
}

func NewA() *A {
	return &A{}
}

type B struct {
	FieldA A
}

func NewB() *B {
	return &B{} // want `FieldA is missing in new B`
}

func newB() *B {
	return &B{} // want `FieldA is missing in new B`
}

func NotAConstructor() *B {
	return &B{}
}

func NewUnnamed() *struct{} {
	return &struct{}{}
}
