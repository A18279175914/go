package main
import "fmt"
import "bytes"
func main(){
	fmt.Println(comma("123456"))
}
func comma(s string) string{
	var buf bytes.Buffer
	n := len(s)
	for i:=0;i<n;i++{
		if (i%3 == 0)&&(i != 0){
			buf.WriteString(",")
		}
		fmt.Printf("%d\n",s[i])
		fmt.Fprintf(&buf,"%c",s[i])
	}
	return buf.String()
}