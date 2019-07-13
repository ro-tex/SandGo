package main

// this is the common interface we aim at using
type infoer interface {
	Info()
}

// parent class that implements the interface
type parent struct {
	Name string
}

func (p parent) Info() {
	println(p.Name)
}

// child class that does not implement the interface but gets it from the embedded parent object
type child struct {
	parent
}

func bar(i infoer) {
	i.Info()
}

func main() {

	p := parent{Name: "John"}
	c := child{parent: p}

	// bar is happy to accept both parent and child
	bar(p)
	bar(c)
}
