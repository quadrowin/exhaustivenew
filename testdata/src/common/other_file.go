package common

import "extra"

type C1 struct {
	extra.ExampleStruct
}

func NewC1() *C1 {
	return &C1{} // want `ExampleStruct is missing in new C1`
}

type C2 struct {
	extra.ExampleStruct
}

func NewC2(es extra.ExampleStruct) *C2 {
	return &C2{
		ExampleStruct: es,
	}
}

type C3 struct {
	C11 C1
	C12 C1 `exhaustivenew:"optional"`
	C13 C1
	C31 *C3
	C32 *C3 `exhaustivenew:"optional"`
	C33 *C3
}

func NewC3() C3 {
	return C3{ // want `C13, C33 are missing in new C3`
		C11: C1{}, // want `ExampleStruct is missing in new C1`
		C31: nil,
	}
}

func NewExtraStruct() *extra.ExampleStruct {
	return &extra.ExampleStruct{} // want `A, B are missing in new ExampleStruct`
}
