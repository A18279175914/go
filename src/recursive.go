package main

import (
	"fmt"
// 	"os"
// 	"golang.org/x/net/html"
)
func square(n int) int { return n * n}
func product(m string,n int) int {fmt.Println(m)
	return n}
func squares() func() int{
	var x int
	return func() int{
		x++
		return x*x
	}
}

func main() {
	var f func(int) int
	var p func(string,int)int
	// sss := "sssss"
	jj,kk:=Parse("ss")
	ll := &jj
	fmt.Println(ll)
	fmt.Println(*ll)
	fmt.Println(jj)
	fmt.Println(kk)
	// switch  sss {
	// 	case "Spades":                                // ...
	// 	case "Hearts":                                // ...
	// 	case "Diamonds":                              // ...
	// 	case "Clubs":                                 // ...
	// 	default:
	// 		panic("xxx") // Joker?
	// }
	f=square
	s :=squares()
	fmt.Println(s()) // "1"
    fmt.Println(s()) // "4"
    fmt.Println(s()) // "9"
    fmt.Println(s()) // "16"
	fmt.Println(f(3))
	p=product
	var a string
	a="a"
	b :="b"
	fmt.Println(p(b,3))
	fmt.Println(product(b,3))
	{
		fmt.Println(b)
		a :="ss"
		fmt.Println(a)
	}
	fmt.Println(a)
	fmt.Println(doit(0))
}
func Parse(input string) (s int, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
	}()
	s=0
	ff(3)
	return
    // ...parser...
}
func ff(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
    defer fmt.Printf("defer %d\n", x)
    ff(x - 1)
}
func doit(x int) int{
	 defer dothat(x)
	 defer dathis(x)
	 if x == 0{
		 return x
	 }
	 defer dathis(x)
	 return x
}
func dothat(x int)int{
	fmt.Println(x+1)
	return x+x
}
func dathis(x int)int{
	fmt.Println(x+2)
	return x*x
}


