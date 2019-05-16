package hello

var Greeting = "Hello Go!" // exported package-level variable

// this is called on package load. it cannot be called from other functions
func init() {
	println(" > This is 'hello' package's init func. < ")
}

// we can have multiple init functions
func init() {
	println(" > This is 'hello' package's second init func. < ")
}

// Hello returns a hello string.
func Hello() string {
	return Greeting
}
