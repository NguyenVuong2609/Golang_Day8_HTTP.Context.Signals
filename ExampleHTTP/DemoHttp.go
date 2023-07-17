package ExampleHTTP

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
}

func TestHttp() {
	http.HandleFunc("/", GetFacts)
	//http.HandleFunc("/route1", RouteTest)
	//http.HandleFunc("/api/facts", GetFacts)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error creating server %s\n", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s! \n", r.URL.Path[1:])
}
func RouteTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye!")
}
func GetFacts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Printf("Error getting data %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing data %s", err.Error())
	}

	var catFacts []Todo
	err = json.Unmarshal(body, &catFacts)
	fmt.Println(catFacts)

	if err != nil {
		log.Printf("Error parsing json %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(catFacts)
	if err != nil {
		log.Printf("Error encoding json %s", err.Error())
	}
}
