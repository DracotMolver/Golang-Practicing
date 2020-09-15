package main

import "fmt"

func main() {
	// func (r receiver) identifier(parameter(s)) (return(s)) {...}
	foo()
	bar("diego")

	s1 := woo("MoneyPenny")
	fmt.Println(s1)

	fn, ln := mouse("Mickey", "Mouse")

	fmt.Println(fn)
	fmt.Println(ln)

	sum := variadic(2, 3, 4, 5, 6, 7)
	fmt.Println("Total: ", sum)

	// Anonymous function
	func() {
		fmt.Println("inside of an anonymous function")
	}()

	func(x int) {
		fmt.Println("inside of an anonymous function", x)
	}(32)

	// func expression
	f := func() {
		fmt.Println("my first fun expression")
	}

	f()

	n := name()
	fmt.Println(n())
}

// returning function
func name() func() string {
	return func() string {
		return "diego"
	}
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(greeting string) {
	fmt.Println("hello, ", greeting)
}

func woo(moviename string) string {
	return fmt.Sprint("Hello from woo, ", moviename)
}

func mouse(firstname, lastname string) (string, bool) {
	fn := fmt.Sprint(firstname, lastname, `, says "Hello"`)
	return fn, false
}

func variadic(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for items in position, ", i, " we are now adding, ", v)
	}

	return sum
}
