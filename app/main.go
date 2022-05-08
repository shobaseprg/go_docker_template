package main

import "fmt"

var mm = map[string]string{"a": "a"}

var _, y = mm["a"]

func main() {
	fmt.Println(y)
}
