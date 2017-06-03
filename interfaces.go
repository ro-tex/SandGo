package main

import "fmt"

type One int
type Two int
type Three int

type NumOne interface {
	IsANum()
}
type NumTwo interface {
	IsANumTwo()
}

type NumThree interface {
	IsANum()
	IsANumTwo()
}

func (n One) IsANum() {}

func (n Two) IsANumTwo() {}

func (n Three) IsANum()    {}
func (n Three) IsANumTwo() {}

func Inter() {
	x := interface{}(Three(42)) // this need to be cast to an interface
	_, ok1 := x.(NumOne)
	_, ok2 := x.(NumTwo)
	_, ok3 := x.(NumThree)
	fmt.Printf(" %d\n NumOne:	%v\n NumTwo:	%v\n NumThree:	%v\n", x, ok1, ok2, ok3)
}
