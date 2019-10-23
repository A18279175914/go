package main
import(
	"fmt"
	"url"
)
type Point struct{ X, Y float64 }
type Valuess []string
func (v Valuess) Ad(value string){
	v= append(v,value)
	fmt.Println(v)
}
func (v Point) Adds(value float64){
	v.X=value
	fmt.Println(v)
}
func main(){
	p := Point{1, 2}
	m := url.Values{"lang": {"en"}} // direct construction
	s := url.Values{"lang": {"en"}}
	ss :=Valuess{"lang"}
m.Add("item", "1")
m.Add("item", "2")
fmt.Println(m)
fmt.Println(m.Get("lang")) // "en"
fmt.Println(m)
fmt.Println(m.Get("q"))    // ""
fmt.Println(m)
fmt.Println(m.Get("item")) // "1"      (first value)
fmt.Println(m)
fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	s.Add("item", "1")
	ss.Ad("1")
	p.Adds(4)
	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(p)
	fmt.Println(m)
}