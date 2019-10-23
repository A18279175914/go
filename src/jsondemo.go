package main
import(
	"encoding/json"
	"fmt"
)
// Phone _
type Phone struct {
    Value string
}
 
// UnmarshalJSON _
func (p *Phone) UnmarshalJSON(data []byte) error {
	// 这里简单演示一下，简单判断即可
	fmt.Println(len(data))
    if len(data) != 11 {
		fmt.Println(len(data))
        return fmt.Errorf("phone format error")
    }
    p.Value = string(data)
    return nil
}
 
// UnmarshalJSON _
// func (p *Phone) MarshalJSON() (data []byte, err error) {
//     if p != nil {
//         data = []byte(p.Value)
// 	}
// 	fmt.Println(2)
// 	fmt.Println(len(data))
//     return
// }
 
// UserRequest _
type UserRequest struct {
    Name  string
    Phone Phone
}
func main() {
	user := UserRequest{}
    user.Name = "ysy"
	user.Phone.Value = "18900001111"
	phone :="18900001111"
	var jsonBlob = []byte(`
		{"Name": "ysy", "Phone": "18900001111"}
	`)
	fmt.Println(len(phone))
	data, _ := json.Marshal(user)
	fmt.Println(jsonBlob)
	fmt.Println(data)
	fmt.Println(AddUser(data))
}
func AddUser(data []byte)(err error)  {
    user := &UserRequest{}
    err = json.Unmarshal(data, user)
    if err != nil {
        return
	}

	return nil
}