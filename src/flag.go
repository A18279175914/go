package main
import (
	"flag"
	"fmt"
	"strings"
)
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func main() {
	flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
	}
	var names [5]string
	nameses := names[:3]
	namese := append(nameses,"aaa")
	name :=append(nameses,"sss")
	name[0]="sks"
	fmt.Print(namese)
	fmt.Print(nameses)
	fmt.Print(name)
	ages := make(map[string]int)
	fmt.Println(ages["aaa"])
}