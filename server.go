package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//https://www.alexedwards.net/blog/serving-static-sites-with-go

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// security measures
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//request used to checker whether requested path is correct
	//if incorrect, server returns a StatusNotFound error
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hiya!")
}

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")

// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }

// func main() {
// 	fileServer := http.FileServer(http.Dir("./static")) //creates file server object
// 	http.Handle("/", fileServer)                        //accepts a path and fileserver
// 	http.HandleFunc("/hello", helloHandler)
// 	http.HandleFunc("/form", formHandler)

// 	// fs := http.FileServer(http.Dir("./css"))
// 	// http.Handle("/css/", http.StripPrefix("/css", fs))

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil { //starts http server with an address and handler
// 		log.Fatal(err)
// 	}
// }
