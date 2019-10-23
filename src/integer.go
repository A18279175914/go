package main
import(
	"fmt"
	"crypto/sha256"
)
type Employee struct{
	ID   int
	Name  string
	Salary int
	ManagerID int
}
var dilbert Employee


func main(){

	dilbert.ID=33333
	// EmployeeById(2).ID = 3333
	EmployeeById(&dilbert).ID =55544
fmt.Println(dilbert)
	var x int8 = 8>>2
	var i int8 = 127
	var s string="sajdfj"
	var z float64 = 258.22
	var h=string(i)
	var c =int8(z)
	var k int =0
	var l int = 5
	var sss rune = 22
	var ssss int32 =22
	var gg []int
	var graph=make(map[string]map[string]bool)
	edges := graph["aaa"]
	if edges == nil{
		edges = make(map[string]bool)
	}
	edges["aaa"]=true
	ages := map[string]int{"A":0}
	ages["aa"]=2
	age,ok :=ages["aaa"]
	ageone,okone :=ages["aaa"]
	months := [...]string{1:"January",2:"February",3:"March",4:"April",5:"May",6:"June",7:"August"}
	q1 := months[1:4]
	q1[2]="ssss"
	q2 := months[2:5]
	var q3 = make([]string,5)
	copy(q1,q2)
	q3[0]="kkkk"
	q2[1]="aaaa"
	sub,err := graph["aaa"]["aaa"]
	fmt.Println(err)
	fmt.Println(sub)
	fmt.Println(edges["aaa"])
	fmt.Println(age)
	fmt.Println(ok)
	fmt.Println(ageone)
	fmt.Println(okone)
	fmt.Println(ages["A"])
	fmt.Println(q3)
	fmt.Println(cap(gg))
	fmt.Println(sss == ssss)
	fmt.Println(cap(q2))
	fmt.Println(len(q1))
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(cap(q3))
	fmt.Println(months)
	fmt.Println(cap(months))
	// var zz  =[256]byte{5:3,200:4}

	var jj uint64 =8888
	var g float64=22222.111111
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%16b\n",jj)
	fmt.Printf("%16b\n",byte(jj>>0))
	fmt.Printf("%16b\n",byte(jj>>1*8))
	// fmt.Println(zz)
	fmt.Println(byte(jj>>8))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	for i,v := range h{
		if i>0{
			fmt.Printf("%c",v)
		}
		fmt.Printf("%c",v)
	}
	fmt.Println(i+3,x,c)
	fmt.Println(s[1])
	fmt.Println(k%l)
	fmt.Printf("%[1]x\n",i)
	fmt.Printf("%08b\n",x)
	fmt.Printf("%10.3f\n",g)
	var arr = []string{"sdf","sdfff"}
	fmt.Println(len(arr))
}
func EmployeeById(id *Employee) *Employee{
	// dilbert.ID=222
	id.ID=5555
	return &dilbert
}
