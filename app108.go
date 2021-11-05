package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 1 // Integer Data Type
	var y int   // Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a, b ,c = 5.25, 25.25,14.45 //
	fmt.Print(reflect.TypeOf(a),a ,reflect.TypeOf(b),b,reflect.TypeOf(c),c)

}