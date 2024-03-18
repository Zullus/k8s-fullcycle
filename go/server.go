package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedat = time.Now()


func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/confimap", configMap)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	hostname, _ := os.Hostname()
	ipserver := os.Getenv("IP_SERVER")
	fmt.Fprintf(w, "<h1>Hello, I'm %s! I'm %s</h1><h2>from %s with ip %s </h2>", name, age, hostname, ipserver)

}

func configMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("myfamily/family.txt")

	if(err != nil){
		log.Fatalf("Error reading file: ", err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))

}

func secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s, Password: %s", user, password)

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedat)
   
	if duration.Seconds() < 10 {	
	  w.WriteHeader(500)
	  w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
	  w.WriteHeader(200)
	  w.Write([]byte("ok"))
	}
  }
