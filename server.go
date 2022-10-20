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

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, name, fruits, numbers, orgs, cities, city)

	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
