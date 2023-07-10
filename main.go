package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/addUser", addUser)
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")

	http.HandleFunc("/testAPI", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("server: %s /\n", r.Method)
		fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
		fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
		fmt.Printf("server: headers:\n")
		for headerName, headerValue := range r.Header {
			fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
		}
		r.ParseForm()
		fmt.Println("request.Form::")
		for key, value := range r.Form {
			fmt.Printf("Key:%s, Value:%s\n", key, value)
		}
		fmt.Printf("server: request body: %s\n", r.Form)

		fmt.Fprintf(w, `{"message": "hello!"}`)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

//func getUser(w http.ResponseWriter, r *http.Request) {

//
//	value, err := client.Get(r.Body).Result()
//	if err != nil {
//		if err == redis.Nil {
//			fmt.Errorf("error")
//		}
//		fmt.Errorf("error")
//	}
//	fmt.Printf(value)
//})
//fmt.Printf("got / request\n")
//io.WriteString(w, "This is my website!\n")
//}

func addUser(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	fmt.Println("request.Form::")
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	fmt.Println("\nrequest.PostForm::")
	for key, value := range r.PostForm {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}

	fmt.Printf("\nName field in Form:%s\n", r.Form["name"])
	fmt.Printf("\nName field in PostForm:%s\n", r.PostForm["name"])
	json, err := json.Marshal(User{Name: r.FormValue("name"), Age: r.FormValue("age")})
	if err != nil {
		fmt.Println(err)
	}
	Id := uuid.New().String()
	err = client.HSet("User", Id, json).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nthe user id :%s\n", Id)
	w.WriteHeader(200)
	return
}
