package main

func foo() int {
	return 42
}

func main() {
	a := foo()
	if a < 10 {
		// ...
	} else if a == 10 {
		// ...
	} else {
		// ...
	}

	if b := foo(); b < 10 { // Short statement
		// ...
	}
}
