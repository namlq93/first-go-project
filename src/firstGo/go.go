package main
import(
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"strconv"
	"log"
)
type User struct {
	Username string
	Id int
	IsMale bool
	CreatedAt time.Time
}
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/echo", post)
	mux.HandleFunc("/redirect", redirect)
	http.ListenAndServe(":9999", mux)
}
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, Line")
}
func post(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		fmt.Fprintf(w, "Only accept POST request")
	}else{
		isMale, err := strconv.ParseBool(r.FormValue("IsMale"))
		if err != nil{
			fmt.Println("Something wrong with key IsMale")
		}
		id, err := strconv.Atoi(r.FormValue("Id"))
		if err != nil{
			fmt.Println("Something wrong with key Id")
		}
		user := User{
			Username : r.FormValue("Username"),
			Id : id,
			IsMale: isMale,
			CreatedAt: time.Now().Local(),
		} 	
		userJson, err := json.Marshal(user)
		if err != nil{
			log.Fatal(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	}
}
func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://www.tmh-techlab.vn/", 301)
}