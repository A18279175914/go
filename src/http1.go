package main
import (
	"net/http"
	"log"
	"fmt"
)
func main(){
	db := database{"shoes":50,"socks":5}
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	mux := http.NewServeMux()
    mux.HandleFunc("/list", http.HandlerFunc(db.list))
    mux.HandleFunc("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
	log.Fatal(http.ListenAndServe("localhost:8001",mux))
}

type dollars float32
type database map[string]dollars
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
func (db database) list(w http.ResponseWriter , req *http.Request){
	for item,price:=range db{
		fmt.Fprintf(w,"%s: %s\n",item,price)
	}
}
func (db database) price(w http.ResponseWriter , req *http.Request){
	item :=req.URL.Query().Get("item")
	price,ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such item: %q\n", item)
        return
	}
	fmt.Fprintf(w,"%s\n",price)
}