package main

import (
	"fmt"
	"log"
	"net/http"
)

// wil be used often (w http.ResponseWriter, r *http.Request)
func formHandler(w http.ResponseWriter, r *http.Request) { // the handle form to accept info given by the user
	err := r.ParseForm()
	// you want people to submit somthing into their html file parsing the form
	// you catch the error if err is not equal to nil to print it out
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful") // this is printing to a file which is the other connect from
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// any api or any route you will have a response and a resquest
// r = request is something the user sends to the server
//w = response is what the server sends back to the user
// * is pointing to the request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { //if not being sent to the correct path you will catch the error
		http.Error(w, "404 not found", http.StatusNotFound)
		return // return from function if it is the condition
	}
	//when you type hello into your browser that is a get method. you dont want people to post anything to hello just get hello printed
	if r.Method != "GET" { // catch the error
		http.Error(w, "method is not supported", http.StatusNotFound) //http status not found is inside the http package
		return
	}
	fmt.Fprintf(w, "hello!") // Fprint prints to file Fprintf prints to fuile along side any format
}
func main() {
	fileServer := http.FileServer(http.Dir("./static")) // telling golang to check out the static directory also too look at the html file
	// http.Handle is pointing to the index.html in the static folder
	http.Handle("/", fileServer)                        //start handling your root route send it to the file server
	//http handle func "/form" id pointing to the form.html file in the static folder.
	http.HandleFunc("/form", formHandler)               // handle func is inside the http pkg
	// http/hanflefunc "/hello" is pointing to the to a hello page that is not in the static folder just allows to print hello onto the screen
	http.HandleFunc("/hello", helloHandler)             // prints out hello to the screen you will need to create a handle func for that.

	fmt.Printf("Starting server at port 8080\n") // this will print once server connects
	err := http.ListenAndServe(":8080", nil)     // catching any errors if given. // assigning err to ListenAndServe
	if err != nil {
		log.Fatal(err)
	}

}

// to run server you simply go into the directory of the file in terminal and run go build then you run go run main.org or name of file
