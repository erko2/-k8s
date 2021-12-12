// package main


// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	fmt.Println("Hi Im running in a Golang container, in a docker contanier, in a kubernetes cluster")

// 	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "Hi Im running in a Golang container, in a docker contanier, in a kubernetes cluster")
// 	})
// 	log.
	

// }
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello there! I am running in a GOLANG docker container path based routing(I am a go APP). My load is distributed through Nginx Ingress on Azure kubernetes cluster!!!")
    })


    log.Fatal(http.ListenAndServe(":80", nil))

}