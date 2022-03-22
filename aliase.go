//Aliase

package main

import "fmt"

func main() {
	//deklarasi variable dengan tipe data uint8

	var a uint8 = 10
	var b byte // byte ini adalah alias dari tipe data uint8

	b = a //no error
	_ = b

	//contoh ke 2
	type second = uint8
	var hour second = 3600
	fmt.Printf("hour type : %T\n", hour) // <==> hour type : uint
}
