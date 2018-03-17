
package main

import "fmt"

func main() {
	a := 111
	function(a)
}

func function(a... int) {
	fmt.Printf("a type is  %T", a)
}