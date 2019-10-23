package main
import "fmt"
func main(){
	var x uint64
	x = 0
	x |= 1 <<63
	fmt.Println(x)
}
