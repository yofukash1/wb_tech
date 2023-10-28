package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = struct{}{}
	// switch v := i.(type) {
	// case int:
	// 	// v is an int here, so e.g. v + 1 is possible.
	// 	fmt.Printf("Integer: %v", v)
	// case float64:
	// 	// v is a float64 here, so e.g. v + 1.0 is possible.
	// 	fmt.Printf("Float64: %v", v)
	// case string:
	// 	// v is a string here, so e.g. v + " Yeah!" is possible.
	// 	fmt.Printf("String: %v", v)
	// default:
	// 	// And here I'm feeling dumb. ;)
	// 	fmt.Printf("I don't know, ask stackoverflow.")
	// }
	fmt.Println(reflect.TypeOf(i))
}
