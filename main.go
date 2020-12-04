package main

import "fmt"

// hotdog라는 타입
type hotdog int

var b hotdog

func main() {
	var a int
	b = 3
	// a = b  error! main.hotdog를 int에 할당할 수 없음
	a = int(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T", b) // main.hotdog
}
