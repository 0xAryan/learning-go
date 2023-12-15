package main

import "fmt"


func Hello() string {
	return "Hello, World"
}

// i commented the old hello.go becuase i thing there can be 
// only one main function in a go project and this one was 
// throwing error

func main() {
	fmt.Println(Hello())
}