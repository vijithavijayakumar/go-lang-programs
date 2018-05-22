package main
 
import "fmt"
import "encoding/json"
 
var JSON = `{
    "name":"TOM",
    "jobtitle":"Software Developer",
    "phone":"1234567890"
    "email":"tom@gmail.com"
}`
 
func main() {
    var info map[string]interface{}
    json.Unmarshal([]byte(JSON),&info)
 
    fmt.Println(info["name"])
    fmt.Println(info["jobtitle"])
    fmt.Println(info["email"])
    fmt.Println(info["phone"])
}
