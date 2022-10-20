/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hello())
}

func main() {

	// Integers
	var a int = 1
	var b int8 = 127
	var c int16 = 3545
	var d int32 = 89080880
	var e int64 = 3837473335

	// Unsigned integers
	var f uint8 = 254
	var g uint16 = 7979
	var h uint32 = 7988978
	var i uint64 = 48094454
	var j uint = 9889889900

	//Float

	var k float32 = 389343.343
	var l float64 = 8.89889899394

	//Strings
	var name string = "Arun" // only double quotes are allowed

	//Arrays.

	// Length of arrays is fixed.
	// We cannot add or remove elements from an array.

	var fruits = [3]string{"Apple", "Banana", "Grapes"} // array of length 3 with string elements
	var numbers = [5]int16{3, 5, 6, 7, 8}               // array of length 5 with integer elements

	var orgs = [...]string{"Apple", "Meta", "Google"} // here length is inferred based on the definition of the array

	//with := syntax
	cities := [5]string{"NewYork", "Washington", "London", "New Delhi", "Mumbai"}

	//Access an array.

	var city = cities[3]
	cities[3] = "Pune" // change the fourth element of cities array

	var array1 = [5]int{}                       // array not initialized
	var array2 = [5]int{1, 5}                   //array partially initialized
	var array3 = [3]int{3, 5, 6}                //array full initialized
	var cars = [10]string{2: "Volvo", 5: "BMW"} // array of length 10 but initialize only 3 and 6 element

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, name, fruits, numbers, orgs, cities, city, array1, array2, array3)
	fmt.Println(cars[1])
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
