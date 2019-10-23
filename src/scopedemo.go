package main
import "fmt"
func main(){
	x := "hello!"
	z :=2
    for i := 0; i < len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Println( x) // "HELLO" (one letter per iteration)
		}
		// fmt.Println( x)
	}
	{var y string
		fmt.Println( y)}
	fmt.Println( x)
	fmt.Println(first(1,z))
}
func first(x int,_ int)int{
	return x
}