
package main
import "fmt"
import "gopi/tempconv"
var kkk int = 2 
var s string = "sfs"
func main(){
	var ggg string
	var x,y=1,3
	// for _, arg:= range ss[1:] {这里好像加上_,就是忽略索引
	// 	s +=arg
	// }
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	p :=&y
	pwk := f()
	x ,y = y,x+y
	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(pwk)
	var f  = jjj("sfs",ggg)
	fmt.Println(f)
	var v  int =2
incr(&v)              // side effect: v is now 2
fmt.Println(incr(&v)) // "3" (and v is 3)
	fmt.Println(v)
}
func jjj(f string,k string) string{
	return f
}
func f() *int{
	return &kkk
}
func incr(p *int) int {
    *p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
    return *p
}

