package main

import (
	"fmt"
	"os"
)

type pointer struct {
	x, y int
}

func main() {
	p := pointer{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 8)
	fmt.Printf("%c\n", 97)

	fmt.Printf("%x\n", 8)
	fmt.Printf("%f\n", 8.2)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.451)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.451)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
